package main

import (
	"net/http"
	"path"

	"github.com/julienschmidt/httprouter"
)

const (
	basePathV1 = "/v1"
)

// routes sets up and returns the main router for the application.
// It delegates the addition of specific API version routes to separate methods.
func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	app.addV1Routes(router)

	return app.recoverPanic(router)
}

// addV1Routes adds all routes for the v1 version of the API to the provided router.
func (app *application) addV1Routes(router *httprouter.Router) {
	app.addRoute(router, http.MethodGet, "/healthcheck", app.healthcheckHandler)
	app.addRoute(router, http.MethodPost, "/articles", app.createArticleHandler)
	app.addRoute(router, http.MethodGet, "/articles/:id", app.showArticleHandler)
}

// addRoute is a helper method that adds a route to the router with the proper base path.
// It joins the base path with the provided route using path.Join to ensure correct formatting.
func (app *application) addRoute(router *httprouter.Router, method, route string, handler http.HandlerFunc) {
	fullPath := path.Join(basePathV1, route)
	router.HandlerFunc(method, fullPath, handler)
}
