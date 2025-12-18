package main

type some1 struct {
	a bool
	b int32
	c string
} //это больше из за выравнивания

type some2 struct {
	b int32
	c string
	a bool
}

//проверить через вывод значения unsafe.Sizeof
//some1 - 24; some2 - 32
