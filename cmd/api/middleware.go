package main

import (
	"fmt"
	"net/http"
)

// recoverPanic recovers from any panics that occur during the request lifecycle.
func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ensure deferred function runs on panic to handle errors gracefully.
		defer func() {
			// Recover from panic to prevent server crash.
			if err := recover(); err != nil {
				// Close connection to signal an error occurred.
				w.Header().Set("Connection", "close")
				// Log the error and send a 500 response to the client.
				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
