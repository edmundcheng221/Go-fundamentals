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

	assert("254")

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
	length float64
	width  float64
}

// (r rect) is a receiver -> "computed properties"
func (r rect) area() float64 {
	return r.length * r.width
}

func (r rect) perimeter() float64 {
	return 2*r.length + 2*r.width
}

// interfaces - collection of method signatures

type shape interface {
	area() float64
}

func assert(num interface{}) {
	var i interface{} = "Hello, World!"

	str, ok := i.(string) // stores the string and ok (boolean) if the assertion is true

	fmt.Println(str, ok)

	// type switch
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}

}
