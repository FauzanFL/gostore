package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (server *Server) initRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", server.Home).Methods("GET")
	server.Router.HandleFunc("/products", server.Products).Methods("GET")

	staticFileDir := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDir))
	server.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods("GET")
}
