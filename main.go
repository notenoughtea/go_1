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
	place      location
	insventory []item
}

type item struct {
	title string
}

func (p player) look(loc location) string {
	answer := ""
	if p.place.title == "кухня" {
		answer += "ты находишься на кухне"
	}
	if len(loc.objects) > 0 {
		for _, obj := range loc.objects {
			if obj.title == "стол" {
				answer += ", на столе: "
				for _, item := range obj.items {
					answer += item.title + ", "
				}
			}
		}
	}
	answer += ", надо собрать рюкзак и идти в универ. можно пройти - "
	for _, path := range loc.paths {
		answer += path + ", "
	}
	return answer
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

func initGame(locs []location) {
	locsInstance := locs
	var hero player = player{place: locs[0]}
	fmt.Println(hero.look(locsInstance[0]))
}

func main() {
	initGame(locations)
}
