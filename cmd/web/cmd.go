package main

import (
	"flag"
	"groupie-tracker/internal/handlers"
	"groupie-tracker/internal/models"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	groups := &models.Data{}

	templateCache, err := handlers.NewTemplateCache()

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	artists, locations, relations, err := handlers.LoadGroupDataCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := handlers.NewApplication(logger, groups, templateCache, artists, locations, relations)

	logger.Info("starting server", "addr", *addr)
	err = http.ListenAndServe(*addr, app.SetupRoutes())
	logger.Error(err.Error())
	os.Exit(1)
}
