package main

import "fmt"

// constant expressions generated at compile time
const (
	// using iota
	first = iota
	second
)

const (
	third  = iota + 4
	fourth = 2 << iota
)

func dataTypes() {
	// int and reassignment
	var i int
	i = 2
	fmt.Println(i)

	// floats / decimals
	var f float32 = 3.141
	fmt.Println(f)

	// implicit declarations
	name := "Bob"
	fmt.Println(name)

	b := true
	fmt.Println(b)

	// complex numbers
	c := complex(3, 4)
	fmt.Println(c)
	r, im := real(c), imag(c)
	fmt.Println(r)
	fmt.Println(im)

	// pointers
	var firstName *string = new(string)
	// addressof
	*firstName = "Arthur"
	fmt.Println(*firstName)

	// constants
	const d int = 3
	fmt.Println(d + 14)

	// type coercion
	fmt.Println(float32(d) + 1.5)

	// print iotas
	fmt.Println(first, second, third, fourth)
}
