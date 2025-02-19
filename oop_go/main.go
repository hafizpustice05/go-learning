package main

import (
	"fmt"
	"go-crush-course/learngo/oop_go/composition"
	"go-crush-course/learngo/oop_go/polymorphism"
)

func main() {

	c := composition.NewCar("Mitsubishi", 1300, 25)
	fmt.Println(c.Hp())
	fmt.Println(c.EngineName())
	fmt.Println(c.WheelDimension())
	fmt.Println(c.WheelName())

	fmt.Println("---------------Interface--------------")

	var cir polymorphism.Shape = polymorphism.Circle{}
	var cir1 polymorphism.Shape = polymorphism.Circle{}
	var r polymorphism.Shape = polymorphism.Squre{}

	cir.Render()
	cir1.Render()
	r.Render()
}
