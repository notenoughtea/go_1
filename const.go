package main

const (
	Look = "осмотреться"
	Go   = "идти"
	Take = "взять"
	Use  = "применить"
	Wear = "надеть"

	ItemBackpack = "рюкзак"
	ItemKeys     = "ключи"
	ItemNotes    = "конспекты"
	ItemTea      = "чай"
	TargetDoor   = "дверь"

	ObjTable = "столе"
	ObjChair = "стуле"

	LocationStreet   = "улица"
	LocationKitchen  = "кухня"
	LocationCorridor = "коридор"
	LocationRoom     = "комната"
	LocationHome     = "домой"

	WhatToUse     = "уточните, что и к чему применить"
	WhatDirection = "уточните направление"
	WhatItem      = "уточните предмет"
	WhatToWear    = "уточните, что надеть"
	EnterCommand  = "введите команду"
	DoorOpen      = "дверь открыта"
	YouWear       = "вы надели: "
	ItemTaken     = "предмет добавлен в инвентарь: "

	ErrUnknownCommand     = "неизвестная команда"
	ErrNoInventorySpace   = "некуда класть"
	ErrDoorClosed         = "дверь закрыта"
	ErrItemNotFound       = "нет такого"
	ErrNoItemInInventory  = "нет предмета в инвентаре - "
	ErrNoObjectInLocation = "не к чему применить"
)
