package main

import (
	"fmt"
	"net/http"

	"github.com/luchanos/go_web_project/pkg/handlers"
)

var portNumer = ":8081"

// Main function
func main() {
	// fmt.Println("Hello, World!")

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/divide", handlers.Divide)

	_, err := fmt.Println(fmt.Sprintf("Starting application at %s", portNumer))
	_ = http.ListenAndServe(portNumer, nil)
	if err != nil {
		fmt.Printf("Error handler: %s", err)
	}
}
