package handlers

import (
	"net/http"
)

func (app *application) SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.homeHandler)
	mux.HandleFunc(`/artists/`, app.artistHandler)

	fs := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	return mux
}
