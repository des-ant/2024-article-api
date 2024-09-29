package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/des-ant/2024-article-api/internal/data"
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
		app.notFoundResponse(w, r)
	}

	article := data.Article{
		ID:    id,
		Title: "Article Title",
		Date:  time.Now(),
		Body:  "This is the body of the article.",
		Tags:  []string{"tag1", "tag2", "tag3"},
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"article": article}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
