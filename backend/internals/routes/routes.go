package routes

import (
	"github.com/Rituraj-13/userReg/backend/internals/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux{
	r := chi.NewRouter()
	r.Get("/health", app.HealthCheck)
	r.Post("/create", app.UserHandler.HandleCreateUser)
	return r
}