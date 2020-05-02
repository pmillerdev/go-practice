package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	// specify regex type
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller!"))
}

// contructor func (always begins with new)
func newUserController() *userController {
	return &userController{
		// look for paths that include /users/a number
		userIDPattern: regexp.MustCompile(`&/users/(\d+)/?`),
	}
}
