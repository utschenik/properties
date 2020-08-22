package routes

import (
	"github.com/go-chi/chi"
	"github.com/utschenik/PropertiesApi/router/handlers"
)

// SetRestRoutes sets the routes for the router for the ressource 'properties'
func SetRestRoutes(r chi.Router) {
	r.Route("/properties", func(r chi.Router) {
		r.Get("/", handlers.GetAllProperties)

		r.Route("/{section}", func(r chi.Router) {
			r.Route("/{name}", func(r chi.Router) {
				r.Get("/", handlers.GetPropertiesBySectionAndName)
				r.Post("/", handlers.CreatePropertyForSection)
			})
		})
	})
}
