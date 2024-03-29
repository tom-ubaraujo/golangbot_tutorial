package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// variable types:

	//bool
	a := true
	b := false
	fmt.Println("a: ", a, "b: ", b)

	c := a && b
	fmt.Println("c:", c)

	d := a || b
	fmt.Println("d: ", d)

	//signed integers
	/*
		int8: represents 8 bit signed integers
		size: 8 bits
		range: -128 to 127

		int16: represents 16 bit signed integers
		size: 16 bits
		range: -32768 to 32767

		int32: represents 32 bit signed integers
		size: 32 bits
		range: -2147483648 to 2147483647

		int64: represents 64 bit signed integers
		size: 64 bits
		range: -9223372036854775808 to 9223372036854775807

		int: represents 32 or 64 bit integers depending on the underlying platform. You should generally
		be using int to represent integers unless there is a need to use a specific sized integer.
		size: 32 bits in 32 bit systems and 64 bit in 64 bit systems.
		range: -2147483648 to 2147483647 in 32 bit systems and -9223372036854775808 to 9223372036854775807 in 64 bit systems
	*/
	var a1 int = 89
	b1 := 95
	fmt.Println("value of a is:", a1, "and b is", b1)
	fmt.Printf("type of a1 is %T, size of a1 is %d \n", a1, unsafe.Sizeof(a1)) //type and size of a1
	fmt.Printf("type of b1 is %T, size of b1 is %d \n", b1, unsafe.Sizeof(b1)) //type and size of b1

	//unsigned integers
	/*
		uint8: represents 8 bit unsigned integers
		size: 8 bits
		range: 0 to 255

		uint16: represents 16 bit unsigned integers
		size: 16 bits
		range: 0 to 65535

		uint32: represents 32 bit unsigned integers
		size: 32 bits
		range: 0 to 4294967295

		uint64: represents 64 bit unsigned integers
		size: 64 bits
		range: 0 to 18446744073709551615

		uint : represents 32 or 64 bit unsigned integers depending on the underlying platform.
		size : 32 bits in 32 bit systems and 64 bits in 64 bit systems.
		range : 0 to 4294967295 in 32 bit systems and 0 to 18446744073709551615 in 64 bit systems
	*/

	//floating point types
	//float32: 32 bit floating point numbers
	//float64: 64 bit floating point numbers
	a2, b2 := 5.67, 8.97
	fmt.Printf("type of a2 %T b2 %T \n", a2, b2)

	sum := a2 + b2
	diff := a2 - b2
	fmt.Println("sum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("no1:", no1, "no2:", no2, " --> sum", no1+no2, "diff", no1-no2)

	//complex types
	//complex64: complex numbers which have float32 real and imaginary parts
	//complex128: complex numbers with float64 real and imaginary parts
	c1 := complex(5, 7)
	c2 := 8 + 27i
	fmt.Println("c1:", c1, "c2:", c2)
	cadd := c1 + c2
	fmt.Println("sum: ", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)

	//other numeric types
	//byte is an alias of uint8
	//rune is an alias of int32

	//string type
	first := "Ugleiston"
	last := "Araújo"
	name := first + " " + last
	fmt.Println("My name is:", name)

	//type conversion
	i := 55            //int
	j := 67.8          //float64
	sum2 := i + int(j) //int + float64 not allowed
	fmt.Println(sum2)
}
