package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegiserControllers() {
	uc := newUserController()

	// pass user controller to http.Handle
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
