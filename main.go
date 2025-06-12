package main

import "fmt"

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
	place       location
	hasBackpack bool
	inventory   []item
}

type item struct {
	title string
}

func findPaths(loc location) string {
	result := "можно пройти - "
	for n := 0; n < len(loc.paths); n++ {
		path := loc.paths[n]
		if n < len(loc.paths)-1 {
			fmt.Println(n)
			result += path + ", "
		} else {
			result += path
		}
	}
	return result
}

func findObjectItems(loc location) string {
	result := ""
	if len(loc.objects) > 0 {
		for i, obj := range loc.objects {
			result += "на " + obj.title + ": "
			for i, item := range obj.items {
				if i < len(obj.items)-1 {
					result += item.title + ", "
				} else {
					result += item.title
				}
			}
			if i < len(loc.objects)-1 {
				result += ", "
			}
		}
	}
	return result
}

func (p player) look(loc location) string {
	result := ""
	switch p.place.title {
	case "кухня":
		answer := "ты находишься на кухне, "
		answer += findObjectItems(loc) + ", "
		answer += "надо собрать рюкзак и идти в универ. "
		answer += findPaths(loc)
		result = answer

	case "корридор":
		answer := "ты находишься на " + p.place.title + ", "
		answer += findObjectItems(loc) + ", "
		answer += "надо собрать рюкзак и идти в универ. "
		answer += findPaths(loc)
		result = answer

	case "комната":
		answer := findObjectItems(loc) + ", "
		answer += findPaths(loc)
		result = answer

	case "улица":
		answer := "ты находишься на " + p.place.title + ", "
		answer += findObjectItems(loc) + ", "
		answer += "надо собрать рюкзак и идти в универ. "
		answer += findPaths(loc)
		result = answer
	}
	return result
}

func (p player) move(loc location) string {
	if loc.locked {
		return "дверь закрыта"
	}
	result := ""
	switch loc.title {
	case "кухня":
		answer := "кухня, ничего интересного. "
		answer += findPaths(loc)
		result = answer
		p.place = loc

	case "корридор":
		answer := "ничего интересного, "
		answer += findPaths(loc)
		result = answer
		p.place = loc

	case "комната":
		answer := "ты в своей комнате, "
		answer += findPaths(loc)
		result = answer
		p.place = loc

	case "улица":
		result = "на улице весна. можно пройти - домой"
	}
	return result
}

func (p player) take(item string, loc location) string {
	result := ""
	for n := 0; n < len(loc.objects); n++ {
		for i := 0; i < len(loc.objects[n].items); i++ {
			if (loc.objects[n].items[i].title == item) && (loc.objects[n].items[i].title == "рюкзак") {
				result := "вы надели: рюкзак"
				p.inventory = append(p.inventory, loc.objects[n].items[i])
				p.inventory = append(p.inventory, loc.objects[n].items[i])
				p.hasBackpack = true
				return result
			}
			if (loc.objects[n].items[i].title == item) && (p.hasBackpack == true) {
				result := "предмет добавлен в инвентарь: " + item
				p.inventory = append(p.inventory, loc.objects[n].items[i])
				loc.objects[n] = Delete(i, i)
				return result
			}
			if (loc.objects[n].items[i].title == item) && (p.hasBackpack == false) {
				result := "нет места"
				return result
			} else {
				result := "нет такого"
				return result
			}
		}
	}
	return result
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
		paths:  []string{"корридор"},
		objects: []obj{
			{
				title: "стол",
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
		paths:  []string{"корридор"},
		objects: []obj{
			{
				title: "стол",
				items: []item{items[3], items[2]},
			},
			{
				title: "стул",
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

func initGame() {
	instance := locations
	var hero player = player{place: instance[2]}
	//fmt.Println(hero.look(locs[2]))
	fmt.Println(hero.move(instance[0]))
}
