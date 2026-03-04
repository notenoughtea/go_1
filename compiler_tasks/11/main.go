package main

import "fmt"

type impl struct{}

type I interface {
	C()
}

func (*impl) C() {}

func A() I {
	return nil
}

func B() I {
	var ret *impl
	return ret
}

func main() {
	a := A()            //nil и значение и тип
	b := B()            //nil только значение
	fmt.Println(a == b) //false
}
