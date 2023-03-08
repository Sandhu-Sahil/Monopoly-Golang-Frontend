package handler

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"
)

type DATA struct {
	Num     int
	Players []Player
	Message string
	Color   [4]string
}

type Player struct {
	Name    string `json:"name"`
	Piece   string `json:"piece"`
	Balance int    `json:"balance"`
	Color   string
}

var color = [4]string{"red", "green", "blue", "yellow"}
var colors = map[string]string{
	"purple":  "#67449e",
	"aqua":    "#5ac9e4",
	"magenta": "#e843ac",
	"orange":  "#ff992a",
	"red":     "#e01f32",
	"yellow":  "#fece00",
	"green":   "#579e50",
	"blue":    "#3160ae",
	"white":   "#dcdcdc",
}

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
		num, err := strconv.Atoi(req.FormValue("aplayers_num"))
		if err != nil {
			log.Fatal(err)
		}

		Color := map[int]string{
			0: "red",
			1: "green",
			2: "blue",
			3: "yellow",
		}
		Image := map[int]string{
			0: "hat",
			1: "dog",
			2: "car",
			3: "boat",
		}

		var Players []Player
		for i := 0; i < num; i++ {
			var playerTemp Player

			s := fmt.Sprint("aplayer_", i+1, "_name")
			p := fmt.Sprint("apiece", i+1)

			playerTemp.Name = req.FormValue(s)
			piece, err := strconv.Atoi(req.FormValue(p))
			if err != nil {
				log.Fatal(err)
			}
			playerTemp.Piece = Image[piece-1]
			playerTemp.Balance = 1500
			playerTemp.Color = Color[i]

			Players = append(Players, playerTemp)

			// SetItem(s, playerTemp.name)
			// SetItem(p, playerTemp.piece)
		}
		// fmt.Print(players)
		// jsonInfo, err := json.Marshal(players)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Printf("%s\n", jsonInfo)

		lp := filepath.Join("templates", "layout.html")
		fp := filepath.Join("templates", "classicPlay.html")

		tmpl, err := template.ParseFiles(lp, fp)
		if err != nil {
			log.Fatal(err)
		}

		Data := DATA{
			Num:     num,
			Players: Players,
			Message: "",
			Color:   color,
		}

		// tmpl.ExecuteTemplate(res, "layout", Data)
		// Use this when you are adding define also to the lp file, alse it takes lp file as base file
		// by this syntax you just specify that start from layout define

		err = tmpl.Execute(res, Data)
		if err != nil {
			log.Fatal(err)
		}
	}
}
