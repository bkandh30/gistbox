package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"gistbox.bhavya.net/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	gists, err := app.gists.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Gists = gists

	app.render(w, r, http.StatusOK, "home.tmpl", data)
}

func (app *application) gistView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	gist, err := app.gists.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Gist = gist

	app.render(w, r, http.StatusOK, "view.tmpl", data)
}

func (app *application) gistCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, r, http.StatusOK, "create.tmpl", data)
}

func (app *application) gistCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "Create a new Gist"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.gists.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/gist/view/%d", id), http.StatusSeeOther)
}
