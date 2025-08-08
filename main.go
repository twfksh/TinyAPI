package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/twfksh/TinyAPI/internal/app"
	"github.com/twfksh/TinyAPI/internal/routes"
)

func main() {
	var port int

	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	app, err := app.InitApplication()
	if err != nil {
		panic(err)
	}
	defer app.Database.Close()

	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("Listening on port: %d\n", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
