package handlers

import (
	"fmt"
	"net/http"

	"github.com/luchanos/go_web_project/pkg/math"
	"github.com/luchanos/go_web_project/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
	sum := math.AddValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is About page with %d", sum))
}

// Func to divide
func Divide(w http.ResponseWriter, r *http.Request) {
	x, y := 100.0, 10.0
	f, err := math.DivideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", x, y, f))
}
