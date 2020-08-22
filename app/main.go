package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/utschenik/PropertiesApi/db"
	routes "github.com/utschenik/PropertiesApi/router"
)

func main() {
	// this sets up the connection to the database
	db.InitClusterGetCollection()
	// this sets up the chi router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	routes.SetRestRoutes(r)
	http.ListenAndServe(":8081", r)
}
