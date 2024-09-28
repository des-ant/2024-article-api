package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
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
	// When httprouter is parsing a request, any interpolated URL parameters will
	// be stored in the request context. Get URL parameters from the request
	// context as a Param slice.
	params := httprouter.ParamsFromContext(r.Context())

	// Convert the "id" parameter to an integer to ensure it's a valid positive ID.
	// Return 404 if the conversion fails or the ID is less than 1.
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Otherwise, interpolate the movie ID in a placeholder response.
	fmt.Fprintf(w, "show the details of article %d\n", id)
}
