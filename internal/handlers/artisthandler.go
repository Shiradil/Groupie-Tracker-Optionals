package handlers

import (
	"fmt"
	models2 "groupie-tracker/internal/models"
	"net/http"
	"strconv"
	"strings"
)

func (app *application) artistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		app.errorHandler(w, r, http.StatusMethodNotAllowed, fmt.Errorf(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

	urlParts := strings.Split(r.URL.Path, "/")
	if len(urlParts) < 3 {
		app.errorHandler(w, r, http.StatusNotFound, fmt.Errorf(http.StatusText(http.StatusNotFound)))
		return
	}

	artistIDStr := urlParts[2]
	artistID, err := strconv.Atoi(artistIDStr)
	if err != nil {
		app.errorHandler(w, r, http.StatusBadRequest, fmt.Errorf(http.StatusText(http.StatusBadRequest)))
		return
	}

	artist, err := app.GetArtistsByID(artistID)
	if err != nil {
		app.errorHandler(w, r, http.StatusNotFound, fmt.Errorf(http.StatusText(http.StatusNotFound)))
		return
	}

	app.render(w, r, http.StatusOK, "artist.html", artist)
}

func (app *application) GetArtistsByID(id int) (models2.Artist, error) {
	return app.artists[id-1], nil
}

func (app *application) GetLocationByID(id int) (models2.Locations, error) {
	return app.locations[id-1], nil
}

func (app *application) GetRelationByID(id int) (models2.Relation, error) {
	return app.relations[id-1], nil
}
