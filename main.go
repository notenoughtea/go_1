package main

import (
	"fmt"
	"slices"
	"strings"
)

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
	result := "можно пройти - "
	for n := 0; n < len((*loc).paths); n++ {
		path := (*loc).paths[n]
		if n < len((*loc).paths)-1 {
			result += path + ", "
		} else {
			result += path
		}
	}
	return result
}

func findObjectItems(loc *location) string {
	result := ""
	if len((*loc).objects) > 0 {
		for i, obj := range (*loc).objects {
			result += "на " + obj.title + ": "
			for i, item := range obj.items {
				if i < len(obj.items)-1 {
					result += item.title + ", "
				} else {
					result += item.title
				}
			}
			if i < len((*loc).objects)-1 {
				result += ". "
			}
		}
	}
	return result
}

func (p *player) look(loc *[]location) string {
	result := ""
	neededLocation := location{}
	for i := 0; i < len(*loc); i++ {
		if p.place.title == (*loc)[i].title {
			neededLocation = (*loc)[i]
		}
	}
	switch p.place.title {
	case "кухня":
		answer := "ты находишься на кухне, "
		answer += findObjectItems(&neededLocation) + ", "
		answer += "надо собрать рюкзак и идти в универ. "
		answer += findPaths(&neededLocation)
		result = answer

	case "коридор":
		answer := "ты находишься на " + p.place.title + ", "
		answer += findObjectItems(&neededLocation) + ", "
		answer += "надо собрать рюкзак и идти в универ. "
		answer += findPaths(&neededLocation)
		result = answer

	case "комната":
		answer := findObjectItems(&neededLocation) + ", "
		answer += findPaths(&neededLocation)
		result = answer

	case "улица":
		answer := "ты находишься на " + p.place.title + ", "
		answer += findObjectItems(&neededLocation) + ", "
		answer += "надо собрать рюкзак и идти в универ. "
		answer += findPaths(&neededLocation)
		result = answer
	}
	return result
}

func (p *player) move(target string, loc *[]location) string {
	neededLocation := &location{}
	for i := 0; i < len(*loc); i++ {
		if target == (*loc)[i].title {
			neededLocation = &(*loc)[i]
		}
	}
	if (neededLocation).locked {
		return "дверь закрыта"
	}
	fmt.Println(p.place.title)
	result := ""
	switch (neededLocation).title {
	case "кухня":
		answer := "кухня, ничего интересного. "
		answer += findPaths(neededLocation)
		result = answer
		p.place = neededLocation

	case "коридор":
		answer := "ничего интересного. "
		answer += findPaths(neededLocation)
		result = answer
		p.place = neededLocation

	case "комната":
		answer := "ты в своей комнате. "
		answer += findPaths(neededLocation)
		result = answer
		p.place = neededLocation

	case "улица":
		result = "на улице весна. можно пройти - домой"
	}
	return result
}

func (p *player) takeItem(item string, loc *[]location) string {

	neededLocation := &location{}
	for i := 0; i < len(*loc); i++ {
		if p.place.title == (*loc)[i].title {
			neededLocation = &(*loc)[i]
		}
	}
	result := ""
	for n := 0; n < len(neededLocation.objects); n++ {
		for i := 0; i < len(neededLocation.objects[n].items); i++ {
			if (neededLocation.objects[n].items[i].title == item) && item == "рюкзак" {
				result := "вы надели: рюкзак"
				p.inventory = append(p.inventory, neededLocation.objects[n].items[i])
				neededLocation.objects[n].items = slices.Delete(neededLocation.objects[n].items, i, i)
				p.hasBackpack = true
				return result
			}
			if (neededLocation.objects[n].items[i].title == item) && p.hasBackpack {
				result := "предмет добавлен в инвентарь: " + item
				p.inventory = append(p.inventory, neededLocation.objects[n].items[i])
				neededLocation.objects[n].items = slices.Delete(neededLocation.objects[n].items, i, i)
				return result
			}
			if (neededLocation.objects[n].items[i].title == item) && !p.hasBackpack {
				result := "нет места"
				return result
			}
		}

	}
	result = "нет такого"
	return result
}

func (p player) useItem(item string, target string, loc *[]location) string {
	result := ""
	neededLocation := location{}
	for i := 0; i < len(*loc); i++ {
		if p.place.title == (*loc)[i].title {
			neededLocation = (*loc)[i]
		}
	}
	if len(target) == 0 {
		for n := 0; n < len(p.inventory); n++ {
			if p.inventory[n].title == item {
				result := "вы использовали: " + item
				return result
			} else {
				result := "нет предмета в инветаре - " + item
				return result
			}
		}
	}
	if len(target) > 0 {
		for n := 0; n < len(p.inventory); n++ {
			if (p.inventory[n].title == item) && (target == "дверь") && (neededLocation.locked == true) {
				result := "дверь открыта"
				neededLocation.locked = false
				return result
			} else {
				result := "не к чему применить"
				return result
			}
		}
	}

	return result
}

func parseCommand(command string) []string {
	splited := strings.Split(command, " ")
	return splited
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
		objects: []obj{},
	},
}

func (p *player) start(command []string, locations *[]location) string {
	result := ""
	switch command[0] {
	case "осмотреться":
		result = p.look(locations)

	case "идти":
		result = p.move(command[1], locations)

	case "взять":
		result = p.takeItem(command[1], locations)

	case "надеть":
		result = p.takeItem(command[1], locations)

	case "применить":
		if len(command) > 1 {
			result = p.useItem(command[1], command[2], locations)
		} else {
			result = p.useItem(command[1], "", locations)
		}
	default:
		result = "неизвестная команда"
	}
	return result
}

func initGame(cases []string) []string {
	var instance *[]location = &locations
	var hero *player = &player{place: &(*instance)[0], hasBackpack: false}
	result := []string{}

	for _, c := range cases {
		result = append(result, hero.start(parseCommand(c), instance))
	}
	return result
}
