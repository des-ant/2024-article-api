package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/des-ant/2024-article-api/internal/data"
	"github.com/des-ant/2024-article-api/internal/validator"
)

// TODO: Use DB to store articles and remove ID from input struct.
func (app *application) createArticleHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID    int64            `json:"id"`
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
		ID:    input.ID,
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

	err = app.daos.Articles.Insert(article)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/articles/%d", article.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"article": article}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
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

	article, err := app.daos.Articles.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"article": article}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
