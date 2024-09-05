package routes

import (
	chi "github.com/go-chi/chi/v5"
	"github.com/tyagnii/todo/internal/options"
	"github.com/tyagnii/todo/internal/transport/handlers"
)

// List http server routes
func InitRouter(cfg options.Options, h handlers.Handler) (chi.Router, error) {
	// Routers
	router := chi.NewRouter()
	router.Get("/tasks", h.GetTasks)
	router.Post("/tasks/", h.PostTasks)
	router.Get("/tasks/{id}", h.GetTaskID)
	router.Put("/tasks/{id}", h.PutTask)
	router.Delete("/tasks/{id}", h.DeleteTask)

	return router, nil
}
