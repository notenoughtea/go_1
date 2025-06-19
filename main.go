package main

import (
	"slices"
	"strings"
)

type Command interface {
	Execute(args []string, p *player, locations *[]location) string
}

type LookCommand struct{}

func (c LookCommand) Execute(args []string, p *player, locations *[]location) string {
	return p.look(locations)
}

type GoCommand struct{}

func (c GoCommand) Execute(args []string, p *player, locations *[]location) string {
	if len(args) < 1 {
		return "уточните направление"
	}
	return p.move(args[0], locations)
}

type TakeCommand struct{}

func (c TakeCommand) Execute(args []string, p *player, locations *[]location) string {
	if len(args) < 1 {
		return "уточните предмет"
	}
	return p.takeItem(args[0], locations)
}

type UseCommand struct{}

func (c UseCommand) Execute(args []string, p *player, locations *[]location) string {
	if len(args) < 2 {
		return "уточните, что и к чему применить"
	}
	itemName := args[0]
	target := strings.Join(args[1:], " ")
	return p.useItem(itemName, target, locations)
}

type CommandDispatcher struct {
	commands map[string]Command
}

func NewDispatcher() *CommandDispatcher {
	return &CommandDispatcher{
		commands: map[string]Command{
			"осмотреться": LookCommand{},
			"идти":        GoCommand{},
			"взять":       TakeCommand{},
			"применить":   UseCommand{},
			"надеть":      WearCommand{},
		},
	}
}

type WearCommand struct{}

func (c WearCommand) Execute(args []string, p *player, locations *[]location) string {
	if len(args) < 1 {
		return "уточните, что надеть"
	}
	itemName := args[0]
	return p.wearItem(itemName, locations)
}

func (p *player) wearItem(itemName string, locations *[]location) string {
	if strings.ToLower(itemName) != "рюкзак" {
		return "неизвестная команда"
	}
	return p.takeItem("рюкзак", locations)
}

func (d *CommandDispatcher) Dispatch(line string, p *player, locations *[]location) string {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return "введите команду"
	}

	cmdName := parts[0]
	args := parts[1:]

	if cmd, ok := d.commands[cmdName]; ok {
		return cmd.Execute(args, p, locations)
	}
	return "неизвестная команда"
}

type location struct {
	title   string
	locked  bool
	paths   []string
	objects []obj
}

type obj struct {
	title string
	items []item
}

type player struct {
	place       *location
	hasBackpack bool
	inventory   []item
}

type item struct {
	title string
}

func findPaths(loc *location) string {
	if len(loc.paths) == 0 {
		return ""
	}
	result := "можно пройти - "
	for n, path := range loc.paths {
		if n > 0 {
			result += ", "
		}
		result += path
	}
	return result
}

func findObjectItems(loc *location) string {
	var items []string
	for _, obj := range loc.objects {
		if len(obj.items) > 0 {
			itemList := ""
			for i, item := range obj.items {
				if i > 0 {
					itemList += ", "
				}
				itemList += item.title
			}
			items = append(items, "на "+obj.title+": "+itemList)
		}
	}
	return strings.Join(items, ", ")
}

func (p *player) look(loc *[]location) string {
	itemsDesc := findObjectItems(p.place)
	pathsDesc := findPaths(p.place)

	switch p.place.title {
	case "кухня":
		answer := "ты находишься на кухне"
		if itemsDesc != "" {
			answer += ", " + itemsDesc
		}
		if p.hasBackpack {
			answer += ", надо идти в универ. "
		} else {
			answer += ", надо собрать рюкзак и идти в универ. "
		}
		answer += pathsDesc
		return answer

	case "комната":
		if itemsDesc == "" {
			return "пустая комната. " + pathsDesc
		}
		return itemsDesc + ". " + pathsDesc

	default:
		answer := "ты находишься на " + p.place.title
		if itemsDesc != "" {
			answer += ", " + itemsDesc
		}
		answer += ". " + pathsDesc
		return answer
	}
}

func (p *player) move(target string, loc *[]location) string {
	var neededLocation *location
	for i := range *loc {
		if target == (*loc)[i].title {
			neededLocation = &(*loc)[i]
			break
		}
	}
	if neededLocation == nil {
		return "нет пути в " + target
	}
	if neededLocation.locked {
		return "дверь закрыта"
	}
	if !slices.Contains(p.place.paths, target) {
		return "нет пути в " + target
	}

	p.place = neededLocation
	switch neededLocation.title {
	case "улица":
		return "на улице весна. можно пройти - домой"
	case "кухня":
		return "кухня, ничего интересного. " + findPaths(neededLocation)
	case "коридор":
		return "ничего интересного. " + findPaths(neededLocation)
	case "комната":
		return "ты в своей комнате. " + findPaths(neededLocation)
	default:
		return findPaths(neededLocation)
	}
}

// store other items. If the backpack is not equipped, other items
// cannot be taken. If the item is successfully taken, a success message
// is returned. If the item is not found or cannot be taken, an appropriate
// message is returned.

func (p *player) takeItem(itemName string, loc *[]location) string {
	// Специальная обработка для рюкзака
	if itemName == "рюкзак" {
		for objIdx := range p.place.objects {
			for itemIdx, it := range p.place.objects[objIdx].items {
				if it.title == "рюкзак" {
					p.hasBackpack = true
					p.inventory = append(p.inventory, it)
					p.place.objects[objIdx].items = slices.Delete(
						p.place.objects[objIdx].items,
						itemIdx,
						itemIdx+1,
					)
					return "вы надели: рюкзак"
				}
			}
		}
		return "нет такого"
	}

	// Для других предметов
	if !p.hasBackpack {
		return "некуда класть"
	}

	for objIdx := range p.place.objects {
		for itemIdx, it := range p.place.objects[objIdx].items {
			if it.title == itemName {
				p.inventory = append(p.inventory, it)
				p.place.objects[objIdx].items = slices.Delete(
					p.place.objects[objIdx].items,
					itemIdx,
					itemIdx+1,
				)
				return "предмет добавлен в инвентарь: " + itemName
			}
		}
	}
	return "нет такого"
}

func (p *player) useItem(itemName string, target string, loc *[]location) string {
	// Проверяем есть ли предмет в инвентаре
	hasItem := false
	for _, it := range p.inventory {
		if it.title == itemName {
			hasItem = true
			break
		}
	}
	if !hasItem {
		return "нет предмета в инвентаре - " + itemName
	}

	// Обрабатываем применение ключей к двери
	if itemName == "ключи" && target == "дверь" {
		for i := range *loc {
			if (*loc)[i].title == "улица" {
				(*loc)[i].locked = false
				return "дверь открыта"
			}
		}
	}

	return "не к чему применить"
}

var items = []item{
	{title: "чай"},
	{title: "рюкзак"},
	{title: "конспекты"},
	{title: "ключи"},
}

var locations = []location{
	{
		title:  "кухня",
		locked: false,
		paths:  []string{"коридор"},
		objects: []obj{
			{
				title: "столе",
				items: []item{items[0]},
			},
		},
	},
	{
		title:   "коридор",
		locked:  false,
		paths:   []string{"кухня", "комната", "улица"},
		objects: []obj{},
	},
	{
		title:  "комната",
		locked: false,
		paths:  []string{"коридор"},
		objects: []obj{
			{
				title: "столе",
				items: []item{items[3], items[2]},
			},
			{
				title: "стуле",
				items: []item{items[1]},
			},
		},
	},
	{
		title:   "улица",
		locked:  true,
		paths:   []string{"домой"},
		objects: []obj{},
	},
}

func (p *player) start(cmd []string, locations *[]location) string {
	dispatcher := NewDispatcher()
	if len(cmd) == 0 {
		return "введите команду"
	}
	return dispatcher.Dispatch(strings.Join(cmd, " "), p, locations)
}

func initGame(cases []string) []string {
	// Создаем копию локаций для каждого теста
	locationsCopy := make([]location, len(locations))
	for i := range locations {
		objectsCopy := make([]obj, len(locations[i].objects))
		for j := range locations[i].objects {
			itemsCopy := make([]item, len(locations[i].objects[j].items))
			copy(itemsCopy, locations[i].objects[j].items)
			objectsCopy[j] = obj{
				title: locations[i].objects[j].title,
				items: itemsCopy,
			}
		}
		pathsCopy := make([]string, len(locations[i].paths))
		copy(pathsCopy, locations[i].paths)
		locationsCopy[i] = location{
			title:   locations[i].title,
			locked:  locations[i].locked,
			paths:   pathsCopy,
			objects: objectsCopy,
		}
	}

	var hero = &player{
		place:       &locationsCopy[0], // Начинаем на кухне
		hasBackpack: false,
	}
	var result []string

	for _, c := range cases {
		result = append(result, hero.start(strings.Fields(c), &locationsCopy))
	}

	return result
}
