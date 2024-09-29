package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/des-ant/2024-article-api/internal/data"
	"github.com/des-ant/2024-article-api/internal/validator"
)

// TODO: Replace this placeholder handler with a function that creates a new
// article.
func (app *application) createArticleHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string           `json:"title"`
		Date  data.ArticleDate `json:"date"`
		Body  string           `json:"body"`
		Tags  []string         `json:"tags"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	article := &data.Article{
		Title: input.Title,
		Date:  input.Date,
		Body:  input.Body,
		Tags:  input.Tags,
	}

	v := validator.New()

	if data.ValidateArticle(v, article); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
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
		Date:  data.ArticleDate(time.Now()),
		Body:  "This is the body of the article.",
		Tags:  []string{"tag1", "tag2", "tag3"},
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"article": article}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
