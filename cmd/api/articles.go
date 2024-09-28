package main

import (
	"fmt"
	"net/http"
)

// TODO: Replace this placeholder handler with a function that creates a new
// article.
func (app *application) createArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new article")
}

// showArticleHandler retrieve the interpolated "id" parameter from the current
// URL and includes it in a placeholder response.
// TODO: Replace this placeholder handler with a function that returns the JSON
// representation of a specific article.
func (app *application) showArticleHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Otherwise, interpolate the movie ID in a placeholder response.
	fmt.Fprintf(w, "show the details of article %d\n", id)
}
