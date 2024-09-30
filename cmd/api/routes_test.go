package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHealthcheck(t *testing.T) {
	// Create a new instance of our application struct.
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/v1/healthcheck")

	assert.Equal(t, code, http.StatusOK)
	assert.Contains(t, body, "available")
}

func TestPostAndGetArticle(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	data := map[string]interface{}{
		"id":    1,
		"title": "latest science shows that potato chips are better for you than sugar",
		"date":  "2016-09-22",
		"body":  "some text, potentially containing simple markup about how potato chip",
		"tags":  []string{"health", "fitness", "science"},
	}

	postStatusCode, _, _ := ts.postJSON(t, "/v1/articles", data)
	assert.Equal(t, postStatusCode, http.StatusCreated)

	getStatusCode, _, body := ts.get(t, "/v1/articles/1")
	assert.Equal(t, getStatusCode, http.StatusOK)
	expectedJSON := `{
			"article": {
					"id": 1,
					"title": "latest science shows that potato chips are better for you than sugar",
					"date": "2016-09-22",
					"body": "some text, potentially containing simple markup about how potato chip",
					"tags": ["health", "fitness", "science"]
			}
	}`
	require.JSONEq(t, expectedJSON, body)
}
