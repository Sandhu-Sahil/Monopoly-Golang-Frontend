package router

import (
	"monopoly-Sandhu-Sahil-frontend/handler"
	"net/http"
)

func BuildRoutes() {
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/classic", handler.Classic)
	http.HandleFunc("/play", handler.ClassicPlay)
}
