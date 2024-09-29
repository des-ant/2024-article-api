package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/des-ant/2024-article-api/internal/data"
	"github.com/des-ant/2024-article-api/internal/validator"
)

// createArticleHandler creates a new article in the system.
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

// showArticleHandler retrieves an article by ID.
// TODO: Use DB to retrieve articles.
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

// getArticlesByTagAndDateHandler retrieves articles by tag and date.
// TODO: Use DB to retrieve articles.
func (app *application) getArticlesByTagAndDateHandler(w http.ResponseWriter, r *http.Request) {
	tagName, date, err := app.readTagAndDateParams(r)
	if err != nil {
		app.notFoundResponse(w, r)
	}

	articles, err := app.daos.Articles.GetArticlesByTagAndDate(tagName, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Get the last 10 article IDs.
	var articleIDs []int64
	for i, article := range articles {
		if i >= 10 {
			break
		}
		articleIDs = append(articleIDs, article.ID)
	}

	relatedTags := app.daos.Articles.GetRelatedTags(articles)

	// Remove the current tag from related tags.
	relatedTags = filter(relatedTags, func(t string) bool {
		return t != tagName
	})

	tagSummary := data.TagSummary{
		Tag:         tagName,
		Count:       len(articles),
		Articles:    articleIDs,
		RelatedTags: relatedTags,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"tag_summary": tagSummary}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
