package main

import (
	"encoding/json"
	"log"
	"os"
	"slices"
	"strings"
)

type GameConfig struct {
	Errors   map[string]string `json:"errors"`
	Messages map[string]string `json:"messages"`
	Commands map[string]string `json:"commands"`
	Items    map[string]string `json:"items"`
	Targets  map[string]string `json:"targets"`
	Places   map[string]string `json:"places"`
}

var config GameConfig

func loadConfig(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("ошибка загрузки config.json: %v", err)
	}
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatalf("ошибка разбора config.json: %v", err)
	}
}

func loadLocations(filename string) []location {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Ошибка при чтении %s: %v", filename, err)
	}

	var locs []location
	if err := json.Unmarshal(data, &locs); err != nil {
		log.Fatalf("Ошибка при парсинге JSON: %v", err)
	}

	return locs
}

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
		return config.Messages["what_direction"]
	}
	return p.move(args[0], locations)
}

type TakeCommand struct{}

func (c TakeCommand) Execute(args []string, p *player, locations *[]location) string {
	if len(args) < 1 {
		return config.Messages["what_item"]
	}
	return p.takeItem(args[0], locations)
}

type UseCommand struct{}

func (c UseCommand) Execute(args []string, p *player, locations *[]location) string {
	if len(args) < 2 {
		return config.Messages["what_to_use"]
	}
	itemName := args[0]
	target := strings.Join(args[1:], " ")
	return p.useItem(itemName, target, locations)
}

type WearCommand struct{}

func (c WearCommand) Execute(args []string, p *player, locations *[]location) string {
	if len(args) < 1 {
		return config.Messages["what_to_wear"]
	}
	itemName := args[0]
	return p.wearItem(itemName, locations)
}

type CommandDispatcher struct {
	commands map[string]Command
}

func NewDispatcher() *CommandDispatcher {
	return &CommandDispatcher{
		commands: map[string]Command{
			config.Commands["look"]: LookCommand{},
			config.Commands["go"]:   GoCommand{},
			config.Commands["take"]: TakeCommand{},
			config.Commands["use"]:  UseCommand{},
			config.Commands["wear"]: WearCommand{},
		},
	}
}

func (d *CommandDispatcher) Dispatch(line string, p *player, locations *[]location) string {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return config.Messages["enter_command"]
	}

	cmdName := parts[0]
	args := parts[1:]

	if cmd, ok := d.commands[cmdName]; ok {
		return cmd.Execute(args, p, locations)
	}
	return config.Errors["unknown_command"]
}

type location struct {
	title         string
	locked        bool
	description   string
	welcome       string
	noItemMessage string
	paths         []string
	objects       []obj
}

type obj struct {
	title string
	items []item
}

type item struct {
	title string
}

type player struct {
	place         *location
	hasBackpack   bool
	inventory     []item
	seenLocations map[string]bool
}

func (p *player) move(target string, locs *[]location) string {
	var neededLocation *location
	for i := range *locs {
		if (*locs)[i].title == target {
			neededLocation = &(*locs)[i]
			break
		}
	}
	if neededLocation == nil || !slices.Contains(p.place.paths, target) {
		return config.Errors["no_path"] + target
	}
	if neededLocation.locked {
		return config.Errors["door_closed"]
	}

	p.place = neededLocation
	return neededLocation.description + ". " + findPaths(neededLocation)
}

func (p *player) takeItem(itemName string, loc *[]location) string {
	if itemName == config.Items["backpack"] {
		for objIdx := range p.place.objects {
			for itemIdx, it := range p.place.objects[objIdx].items {
				if it.title == itemName {
					p.hasBackpack = true
					p.inventory = append(p.inventory, it)
					p.place.objects[objIdx].items = slices.Delete(p.place.objects[objIdx].items, itemIdx, itemIdx+1)
					return config.Messages["you_wear"] + itemName
				}
			}
		}
		return config.Errors["item_not_found"]
	}
	if !p.hasBackpack {
		return config.Errors["no_inventory"]
	}
	for objIdx := range p.place.objects {
		for itemIdx, it := range p.place.objects[objIdx].items {
			if it.title == itemName {
				p.inventory = append(p.inventory, it)
				p.place.objects[objIdx].items = slices.Delete(p.place.objects[objIdx].items, itemIdx, itemIdx+1)
				return config.Messages["item_taken"] + itemName
			}
		}
	}
	return config.Errors["item_not_found"]
}

func (p *player) useItem(itemName string, target string, loc *[]location) string {
	hasItem := false
	for _, it := range p.inventory {
		if it.title == itemName {
			hasItem = true
			break
		}
	}
	if !hasItem {
		return config.Errors["no_item_in_inventory"] + itemName
	}
	if itemName == config.Items["keys"] && target == config.Targets["door"] {
		for i := range *loc {
			if (*loc)[i].title == config.Places["street"] {
				(*loc)[i].locked = false
				return config.Messages["door_open"]
			}
		}
	}
	return config.Errors["no_object_in_location"]
}

func (p *player) wearItem(itemName string, locations *[]location) string {
	if strings.ToLower(itemName) != config.Items["backpack"] {
		return config.Errors["unknown_command"]
	}
	return p.takeItem(config.Items["backpack"], locations)
}

func (p *player) look(loc *[]location) string {
	itemsDesc := findObjectItems(p.place)
	pathsDesc := findPaths(p.place)

	locTitle := p.place.title
	var result string
	if !p.seenLocations[locTitle] {
		if p.place.welcome != "" {
			result += p.place.welcome
		} else {
			result += p.place.description
		}
		p.seenLocations[locTitle] = true
	} else {
		if itemsDesc == "" {
			result += p.place.noItemMessage
		} else {
			result += itemsDesc
		}
	}
	result += ". " + pathsDesc
	return result
}

func findPaths(loc *location) string {
	if len(loc.paths) == 0 {
		return ""
	}
	result := config.Messages["can_go"]
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
			items = append(items, obj.title+": "+itemList)
		}
	}
	return strings.Join(items, ", ")
}

func (p *player) start(cmd []string, locations *[]location) string {
	dispatcher := NewDispatcher()
	if len(cmd) == 0 {
		return config.Messages["enter_command"]
	}
	return dispatcher.Dispatch(strings.Join(cmd, " "), p, locations)
}

func initGame(cases []string) []string {
	loadConfig("config.json")
	originalLocations := loadLocations("locations.json")
	locationsCopy := make([]location, len(originalLocations))
	for i := range originalLocations {
		objectsCopy := make([]obj, len(originalLocations[i].objects))
		for j := range originalLocations[i].objects {
			itemsCopy := make([]item, len(originalLocations[i].objects[j].items))
			copy(itemsCopy, originalLocations[i].objects[j].items)
			objectsCopy[j] = obj{
				title: originalLocations[i].objects[j].title,
				items: itemsCopy,
			}
		}
		pathsCopy := make([]string, len(originalLocations[i].paths))
		copy(pathsCopy, originalLocations[i].paths)
		locationsCopy[i] = location{
			title:         originalLocations[i].title,
			locked:        originalLocations[i].locked,
			description:   originalLocations[i].description,
			welcome:       originalLocations[i].welcome,
			noItemMessage: originalLocations[i].noItemMessage,
			paths:         pathsCopy,
			objects:       objectsCopy,
		}
	}

	hero := &player{
		place:         &locationsCopy[0],
		hasBackpack:   false,
		seenLocations: make(map[string]bool),
	}

	var result []string
	for _, c := range cases {
		result = append(result, hero.start(strings.Fields(c), &locationsCopy))
	}
	return result
}
