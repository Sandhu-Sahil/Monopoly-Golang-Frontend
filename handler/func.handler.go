package handler

import (
	"log"
	"net/http"
	"text/template"
)

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.Error(res, "404 not found", http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		http.Error(res, "method is not supported", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(res, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Classic(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/classic" {
		http.Error(res, "404 not found", http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		http.Error(res, "method is not supported", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/classicMonopoly.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(res, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ClassicPlay(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/play" {
		http.Error(res, "404 not found", http.StatusNotFound)
		return
	}
	if req.Method == "GET" {
		http.Error(res, "method is not supported", http.StatusNotFound)
		return
	}
	if req.Method == "POST" {
		// lp := filepath.Join("templates", "templates/layout.html")
		// fp := filepath.Join("templates", "templates/classicPlay.html")

		tmpl, err := template.ParseFiles("templates/layout.html", "templates/classicPlay.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(res, "layout", nil)
	}
}
