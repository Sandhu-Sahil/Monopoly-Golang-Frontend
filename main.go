package main

import (
	"fmt"
	"log"
	"monopoly-Sandhu-Sahil-frontend/router"
	"net/http"
)

func main() {
	// add static files in http server
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	router.BuildRoutes()

	fmt.Print("Server started on http://localhost:5000/")

	err := http.ListenAndServe(":5000", nil)
	log.Fatal(err)
}
