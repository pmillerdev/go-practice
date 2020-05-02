package main

import (
	"fmt"
)

func main() {
	// arrays
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	// implicit initialization
	arr1 := [3]int{1, 2, 3}
	fmt.Println(arr1)

	// slice from array declaration
	slice := arr1[:]
	arr1[1] = 42
	slice[2] = 27
	fmt.Println(arr1, slice)
	// slices
	slice1 := []int{1, 2, 3}
	// append operation resizes memory for array
	slice1 = append(slice1, 4, 8, 16)
	fmt.Println(slice1)
	// subslices
	s2 := slice1[1:]
	s3 := slice1[:2]
	s4 := slice1[1:2]
	fmt.Println(s2, s3, s4)

	// maps
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])
	// assignment via key
	m["foo"] = 27
	fmt.Println(m)
	// deletion via key
	delete(m, "foo")
	fmt.Println(m)

	// structs
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}
	// initilization
	var u user
	// assigning values
	u.ID = 1
	u.FirstName = "Bob"
	u.LastName = "Builder"
	fmt.Println(u)
	// shorthand initilization
	u1 := user{
		ID:        2,
		FirstName: "Pete",
		LastName:  "Miller",
	}
	fmt.Println(u1)

}
