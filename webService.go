package main

import (
	"fmt"

	"github.com/pmillerdev/test-mod/models"
)

// Web service
func main() {
	// use User from the models package
	u := models.User{
		ID:        2,
		FirstName: "Frank",
		LastName:  "Furter",
	}
	fmt.Println(u)
}
