package handlers

import (
	"fmt"
	"net/http"
)

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.errorHandler(w, r, http.StatusNotFound, fmt.Errorf(http.StatusText(http.StatusNotFound)))
		return
	}

	if r.Method != "GET" {
		app.errorHandler(w, r, http.StatusMethodNotAllowed, fmt.Errorf(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

	app.render(w, r, http.StatusOK, "main.html", app.artists)
}
