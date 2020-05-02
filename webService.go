package main

import (
	"net/http"

	"github.com/pmillerdev/go-practice/controllers"
)

// Web service
func main() {
	controllers.RegiserControllers()
	http.ListenAndServe(":3000", nil)
}
