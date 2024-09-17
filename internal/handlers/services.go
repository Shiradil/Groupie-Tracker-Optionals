package handlers

import (
	"bytes"
	"fmt"
	"groupie-tracker/internal/models"
	"html/template"
	"log/slog"
	"net/http"
)

type application struct {
	logger        *slog.Logger
	groups        *models.Data
	templateCache map[string]*template.Template
	artists       []models.Artist
	locations     []models.Locations
	relations     []models.Relation
}

func NewApplication(logger *slog.Logger, groups *models.Data, templateCache map[string]*template.Template, artists []models.Artist, locations []models.Locations, relations []models.Relation) *application {
	return &application{
		logger:        logger,
		groups:        groups,
		templateCache: templateCache,
		artists:       artists,
		locations:     locations,
		relations:     relations,
	}
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data any) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}

func (app *application) errorHandler(w http.ResponseWriter, r *http.Request, code int, err error) {
	errStr := err.Error()

	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	errorData := models.Error{
		ErrorCode: code,
		ErrorMsg:  errStr,
	}

	app.logger.Error(err.Error(), method, "uri", uri)
	app.render(w, r, code, "error.html", errorData)
}
