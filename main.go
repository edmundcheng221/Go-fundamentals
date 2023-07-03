package main

import (
	"fmt"
)

// 8, 16, 32, 64, 128 -> bit size
// Types: int, uint, float64, complex128, bool

// One way to declare
var empty string = "hello world"

func main() {
	value := fmt.Sprintf("%d", variables())
	fmt.Println(value)

	pointers()

	x, _ := getCoordinates(22.3, 23.3) // the _ means we want to ignore the return value y
	fmt.Println(x)

	getCarDetails()

	r := rect{length: 5, width: 3}
	result := r.area()
	fmt.Println(result) // Output: 15

}

func variables() int {
	happy := "happy" // shorthand
	fmt.Printf(happy + "\n")
	milage, company := 200000, "Toyota"
	fmt.Println(milage, company)
	age := 2.6
	fmt.Println(int(age))
	return 20
}

func pointers() {
	var num int = 42
	var ptr *int = &num

	fmt.Println("Value of num:", num)
	fmt.Println("Value of ptr:", ptr)
	fmt.Println("Value at memory address pointed by ptr:", *ptr)

	*ptr = 99
	fmt.Println("Updated value of num:", num)
}

func getCoordinates(x, y float64) (float64, float64) {
	return x, y
}

// Structs are collections of key value pairs

type car struct {
	Make   string
	Model  string
	weight float64
}

func getCarDetails() {
	car := car{
		Make:   "Tesla",
		Model:  "3",
		weight: 1255.6,
	}
	fmt.Printf("The %s model %s weighs %.1f pounds\n", car.Make, car.Model, car.weight)
	car.Make = "Toyota"
	car.Model = "Prius"
	car.weight = 600.29
	fmt.Printf("New: The %s model %s weighs %.1f pounds\n", car.Make, car.Model, car.weight)

	// anon struct

	food := struct {
		category string
		rating   int
	}{
		category: "pizza",
		rating:   9,
	}
	fmt.Print(food)
}

type rect struct {
	length int
	width  int
}

// (r rect) is a receiver -> "computed properties"
func (r rect) area() int {
	return r.length * r.width
}
