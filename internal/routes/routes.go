package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/twfksh/TinyAPI/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/healthcheck", app.HealthCheck)

	r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)
	r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutById)
	r.Put("/workouts/{id}", app.WorkoutHandler.HandleUpdateWorkoutById)
	r.Delete("/workouts/{id}", app.WorkoutHandler.HandleDeleteWorkoutById)

	return r
}
