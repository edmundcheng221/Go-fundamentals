package main

import (
	"edmundcheng221/go-essentials/helper"
	errs "errors" // errs is alias
	"fmt"
)

// 8, 16, 32, 64, 128 -> bit size
// Types: int, uint, float64, complex128, bool

// One way to declare
var empty string = "hello world"

func contains(s string, e rune) bool {
	for _, a := range s {
		if string(a) == string(e) {
			return true
		}
	}
	return false
}

func removeLetterFromString(str string, letter rune) string {
	result := ""
	for _, char := range str {
		if char != letter {
			result += string(char)
		}
	}
	return result
}

func findHiddenWord(str string) string {
	dictionary := [2]string{"cat", "baby"}
	for _, word := range dictionary {
		copy := word
		for _, char := range copy {
			if contains(str, char) {
				copy = removeLetterFromString(copy, char)
			}
		}
		if len(copy) == 0 {
			return word
		}
	}
	return "-"
}

func main() {
	str := "jnsfobcsdnvtdsva"
	fmt.Println(findHiddenWord(str))
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
	operators()

	fmt.Println(formatStrings(28.365876))

	fmt.Println(usingErrorsPackage())

	forLoops(20)

	fmt.Println(helper.Helloworld())

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

	str = "hello"

	fmt.Println(str)

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

func operators() {
	number := 20
	fmt.Println(number)
	number %= 3
	fmt.Println(number)
}

func someFunctionThatCanHaveErrors() (int, error) {
	return 20, fmt.Errorf("error")
}

// errors
func someerrors() {
	// one way
	response, err := someFunctionThatCanHaveErrors()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
}

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	cost4cust, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, userError{name: "ahsfbsdh"}
	}
	cost4spouse, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0.0, err
	}
	return cost4cust + cost4spouse, nil

}

func sendSMS(message string) (float64, error) {
	return 20, nil
}

func formatStrings(age float64) string {
	return fmt.Sprintf("Edmund is %.2f years old", age)
}

// errors are interfaces
// You can build custom types that implement error

type userError struct {
	name string
}

func (err userError) Error() string {
	return fmt.Sprintf("Error with user's account, %v", err.name)
}

func usingErrorsPackage() error {
	return errs.New("Some error message")
}

// loops

func forLoops(threshold float64) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// "while" loop -> no explicit while loop, we use for
	// for condition {}
	num := 2
	for num < 5 {
		num++
		fmt.Println(num)
	}
	totalCost := 0.0
	for i := 0; ; i++ { // initial; condition; after -> you can omit parts
		// omitting the condition will make the loop run forever
		totalCost += 1.0 + (0.01 * float64(i))
		if totalCost > threshold {
			return
		}
		fmt.Printf("%.2f", totalCost)
		fmt.Println()
	}

}

// continue -> continues to next iteraction; bail out current loop early
// break -> breaks out of loop
