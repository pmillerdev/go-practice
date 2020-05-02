package controllers

import "net/http"

func RegiserControllers() {
	uc := newUserController()

	// pass user controller to http.Handle
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
