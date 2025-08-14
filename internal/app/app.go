package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/twfksh/TinyAPI/internal/api"
	"github.com/twfksh/TinyAPI/internal/store"
	"github.com/twfksh/TinyAPI/migrations"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	Database       *sql.DB
}

func InitApplication() (*Application, error) {
	db, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// INFO: stores go here
	workoutStore := store.NewPostgresWorkoutStore(db)

	// INFO: handlers go here
	workoutHandler := api.NewWorkoutHandler(workoutStore)

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		Database:       db,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is available")
}
