package main

type place struct {
	title string
	locked bool
	objects []obj
}

type obj struct{
	title string
	items []item
}

type player struct {
	name  string
	place string
	lookedAround bool
	insventory []item
}

type item struct{
	title string
}


levelTitles := []string{"кухня", "комната", "улица", "корридор"}
lockedPlaces := []string{"улица"}
loot := map[string]string{"комната":}


func createLevels (titles []string, lockedList []string) []string {
		lev := []string{}
		for i := 0; i < len(titles); i++ {
					for n := 0; n < len(lockedList); n++ {
									if (titles[i] == lockedList(n)) {
									lev = append(lev, place{title: titles[i], locked: true})
									break
					}
}
									lev = append(lev, place{title: titles[i]})
	return lev
	}
}
func initGame() {
	levels := []place{}
	levels = append(levels, createLevels(levelTitles, lockedList))
	var BobGross player = player{}
}
