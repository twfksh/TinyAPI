package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/twfksh/TinyAPI/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/healthcheck", app.HealthCheck)
	r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutById)

	r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)

	return r
}
