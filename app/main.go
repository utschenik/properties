package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	routes "github.com/utschenik/PropertiesApi/router"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	routes.SetRestRoutes(r)
	// collection := db.InitClusterGetCollection()
	http.ListenAndServe(":8081", r)
}
