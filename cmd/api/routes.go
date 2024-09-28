package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// Register methods, URL patterns, and handlers using HandlerFunc().
	// http.MethodGet and http.MethodPost are "GET" and "POST" constants.
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/articles", app.createArticleHandler)
	router.HandlerFunc(http.MethodGet, "/v1/articles/:id", app.showArticleHandler)

	return router
}
