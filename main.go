package main

import "fmt"

func main() {
	var i int
	i = 2
	fmt.Println(i)

	var f float32 = 3.141
	fmt.Println(f)

	name := "Bob"
	fmt.Println(name)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r)
	fmt.Println(im)

	var firstName *string = new(string)

	*firstName = "Arthur"
	fmt.Println(*firstName)
}
