package main

import (
	"fmt"
	"math"
)

// the operator ':=' can be used to create variables only INSIDE FUNCTIONS(including the main() function)
//while the use of the var keyword can be used to create variables anywhere in the code base(even outside the main() function)

func main() {

	//declaring a variable automatically initialized with 0 value.
	var age int // variable declaration
	fmt.Println("My age is: ", age)
	age = 29 //assignment
	fmt.Println("My age is: ", age)
	age = 54 // assignment
	fmt.Println("My new age is: ", age)

	//declaring a variable with initial value
	var name string = "Tom"
	fmt.Println("My name is: ", name)

	//declaring a variable with Type inference
	var income = 8000.00
	fmt.Println("My income is: ", income)
	fmt.Printf("%T\n", income) // checking the type of a variable

	//multiple variable declaration
	var width, height int = 100, 50 //the type can be removed if the variables have initial values
	fmt.Println("width is: ", width, "height is: ", height)

	//declaring multiple variables with different types
	var (
		name2   = "naveen"
		age2    = 29
		height2 int
	)
	fmt.Println("My name is", name2, ", age is", age2, "and height is", height2)

	//short hand declaration
	count := 10
	fmt.Println("Count =", count)

	//multiple variables with short hand declaration
	//can only be used when at least one of the variables on the left of := is a new variable
	name3, age3 := "naveen", 29
	fmt.Println("my name is", name3, ", age is", age3)

	//assigning values which are computed during run time
	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Println("Mininum value is: ", c)
}
