package main

type Person struct {
	ID        int
	FirstName string
	LastName  string
}

type HTTPRequest struct {
	Method string
}

func main() {
	// standard for loop syntax
	for i := 0; i < 5; i++ {
		println(i)
		// if statement
		if i == 3 {
			// continue jumps to the next interation of loop
			continue
		}
		println("continuing..")
	}

	// infinite loops
	var j int
	for {
		if j == 5 {
			// removing break will make it infinite
			break
		}
		j++
	}

	// looping over slices using range
	slice := []int{1, 2, 3}
	for i, v := range slice {
		println(i, v)
	}

	// looping over maps
	Map := map[string]int{"http": 80, "https": 443}
	for k, v := range Map {
		println(k, v)
	}

	// panic functions
	// panic("Something bad happened")

	// delcare structs
	p1 := Person{
		ID:        1,
		FirstName: "Bob",
		LastName:  "The Builder",
	}

	p2 := Person{
		ID:        2,
		FirstName: "Frank",
		LastName:  "Butcher",
	}

	// conditionals for structs
	if p1.ID == p2.ID {
		println("Same user!")
	} else if p1.FirstName == p2.FirstName {
		println("Similar user!")
	} else {
		println("Different user!")
	}

	// switch statements
	r := HTTPRequest{Method: "PUT"}

	switch r.Method {
	case "GET":
		println("Get request")
		// jump to next case
		fallthrough
	case "DELETE":
		println("Delete request")
	case "POST":
		println("Post request")
	case "PUT":
		println("Put request")
	default:
		println("Unhandled method")
	}
}
