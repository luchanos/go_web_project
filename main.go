package main

import (
	"errors"
	"fmt"
	"net/http"
	"text/template"
)

var portNumer = ":8081"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is About page with %d", sum))
}

// Add Functioner
func addValues(x, y int) int {
	return x + y
}

// For parsing template
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error during parsing template:", err)
		return
	}
}

// Func to divide
func Divide(w http.ResponseWriter, r *http.Request) {
	x, y := 100.0, 10.0
	f, err := divideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", x, y, f))
}

// Func for division
func divideValues(x, y float64) (float64, error) {
	if y == 0 {
		err := errors.New("Cannot divide by zero!")
		return 0, err
	}
	return x / y, nil
}

// Main function
func main() {
	// fmt.Println("Hello, World!")

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	_ = http.ListenAndServe(portNumer, nil)
	_, err := fmt.Println(fmt.Sprintf("Starting application at %s", portNumer))
	if err != nil {
		fmt.Printf("Error handler: %s", err)
	}
}
