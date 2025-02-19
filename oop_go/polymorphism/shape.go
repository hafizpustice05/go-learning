package polymorphism

import "fmt"

type Shape interface {
	Render()
}

type Circle struct{}

func (c Circle) Render() {
	fmt.Println("Circle Rendering method")
}

type Squre struct{}

func (c Squre) Render() {
	fmt.Println("Square rendering method")
}
