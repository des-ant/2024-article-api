package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/des-ant/2024-article-api/internal/data"
	"github.com/julienschmidt/httprouter"
)

// readIDParam retrieves the "id" URL parameter from the request context,
// converts it to an integer, and returns it. Returns 0 and an error if unsuccessful.
func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

// readTagAndDateParams retrieves the "tagName" and "date" URL parameters from the request context,
// converts the date to a time.Time object, and returns them. Returns an error if unsuccessful.
func (app *application) readTagAndDateParams(r *http.Request) (string, data.ArticleDate, error) {
	params := httprouter.ParamsFromContext(r.Context())

	tagName := params.ByName("tagName")
	dateStr := params.ByName("date")

	// Parse the date
	date, err := time.Parse("20060102", dateStr)
	if err != nil {
		return "", data.ArticleDate{}, errors.New("invalid date format")
	}

	return tagName, data.ArticleDate(date), nil
}

// envelope is a generic type that we can use to hold the response envelope.
type envelope map[string]any

// writeJSON writes the provided data to the http.ResponseWriter as JSON.
func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// Append a newline to make it easier to view in terminal applications.
	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

// readJSON decodes the request body into the provided destination.
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	// Limit the size of the request body to prevent potential denial-of-service attacks.
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	// Require that JSON keys in the request body must match the destination struct fields.
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		var maxBytesError *http.MaxBytesError

		switch {
		// Handle JSON syntax errors to provide clear error location.
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		// Handle unexpected EOF to provide a generic syntax error message.
		// Decode() will return io.ErrUnexpectedEOF if the JSON ends abruptly.
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		// Handle type errors to help clients debug incorrect JSON fields.
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		// Handle empty body to inform clients that body must not be empty.
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		// Handle unknown fields to inform clients that the request body contains unknown keys.
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)

		// Handle body size limit exceeded to inform clients that the body must not exceed 1MB.
		case errors.As(err, &maxBytesError):
			return fmt.Errorf("body must not be larger than %d bytes", maxBytesError.Limit)

		// Panic on invalid unmarshal to catch non-nil pointer issues.
		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		default:
			return err
		}
	}

	// Ensure the request body only contains a single JSON value.
	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}

// filter returns a new slice containing only the elements of slice that satisfy the predicate.
func filter(slice []string, predicate func(string) bool) []string {
	var result []string
	for _, s := range slice {
		if predicate(s) {
			result = append(result, s)
		}
	}
	return result
}
