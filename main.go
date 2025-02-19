package main

import "fmt"

type Student struct {
	name string // name of the object
	id   int    // its value
	age  int
}

func main() {

	// var fulName, country string
	// var age int
	// var gpa float64

	fmt.Println("Hello, Go!")

	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Printf("\ni=%d ", i)
		sum += i
	}
	fmt.Println("ths sum is= ", sum)

	array := []float64{7.0, 8.5, 9.1, 5.4}
	x := Sum(&array) // Note the explicit address-of operator
	fmt.Println("THe sum is ", x)

	hafizul := Student{"Hafizul Islam", 101, 29}
	displayInfo(hafizul)
	hafizul.increaseAge(2)
	displayInfo(hafizul)

}

func Sum(a *[]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

func displayInfo(s Student) {
	fmt.Println("Id: ", s.id)
	fmt.Println("Name: ", s.name)
	fmt.Println("Age: ", s.age)
}

func (s *Student) increaseAge(val int) {
	s.age = s.age + val
}
