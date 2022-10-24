package main

import "fmt"

//动物接口

type Animals interface {
	Shout()
}

//有哪些动物

type Dog struct{}
type Cat struct{}
type Xuebao struct{}

//动物叫的方法

func (d Dog) Shout() {
	fmt.Println("汪汪汪~")
}
func (c Cat) Shout() {
	fmt.Println("喵喵喵~")
}
func (fifteen Xuebao) Shout() {
	fmt.Println("雪豹闭嘴!")
}

//只要是动物类,就叫(接口类型做参数)

func AnimalShout(animal Animals) {
	animal.Shout()
}
func main() {
	Dog := Dog{}
	Cat := Cat{}
	Xuebao := Xuebao{}
	AnimalShout(Dog)
	AnimalShout(Cat)
	AnimalShout(Xuebao)
}
