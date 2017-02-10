// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to create and use your own mux.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	m := createHandlers()
	// Start the http server to handle the request.
	log.Panic(http.ListenAndServe(":3000", m))
}

func createHandlers() *http.ServeMux {

	// Create a new mux for handling routes.
	m := http.NewServeMux()

	// Bind a handler to the root route.
	f := func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello World!")
	}
	m.HandleFunc("/", f)

	// Bind a handler to the /prt route.
	f = func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Olá Mundo")
	}
	m.HandleFunc("/prt", f)

	// Bind a handler to the /spa route.
	f = func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hola Mundo")
	}
	m.HandleFunc("/spa", f)

	// Bind a handler to the /fre route.
	f = func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Bonjour le monde")
	}
	m.HandleFunc("/fre", f)

	return m
}
