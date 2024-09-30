package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthcheck(t *testing.T) {
	expectedBody := "{\n\t\"status\": \"available\",\n\t\"system_info\": {\n\t\t\"environment\": \"test\",\n\t\t\"version\": \"1.0.0\"\n\t}\n}"

	// Create a new instance of our application struct.
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/v1/healthcheck")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, expectedBody)
}
