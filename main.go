package main

import (
	"fmt"
	"net/http"
)

var portNumer = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is About page with %d", sum))
}

// Add Functioner
func addValues(x, y int) int {
	return x + y
}

// Main function
func main() {
	// fmt.Println("Hello, World!")

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(portNumer, nil)
	_, err := fmt.Println(fmt.Sprintf("Starting application at %s", portNumer))
	if err != nil {
		fmt.Printf("Error handler: %s", err)
	}
}
