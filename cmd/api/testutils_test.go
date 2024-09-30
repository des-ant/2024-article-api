package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/des-ant/2024-article-api/internal/data"
)

// newTestApplication creates a new instance of the application struct with mocked dependencies.
func newTestApplication(t *testing.T) *application {
	cfg := config{
		port: 4000,
		env:  "test",
	}

	return &application{
		config: cfg,
		logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
		daos:   data.NewDAOs(),
	}
}

// testServer wraps httptest.Server to provide helper methods for testing.
type testServer struct {
	*httptest.Server
}

// newTestServer creates a new instance of testServer with the provided handler.
func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewTLSServer(h)

	return &testServer{ts}
}

// get performs a GET request to the server and returns the response status code, headers, and body.
func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, string) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)

	return rs.StatusCode, rs.Header, string(body)
}

// postJSON performs a POST request to the server with JSON data and returns the response status code, headers, and body.
func (ts *testServer) postJSON(t *testing.T, urlPath string, data interface{}) (int, http.Header, string) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	rs, err := ts.Client().Post(ts.URL+urlPath, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)

	return rs.StatusCode, rs.Header, string(body)
}

// checkResponse checks the response status code and body against the expected values.
func checkResponse(t *testing.T, ts *testServer, urlPath string, expectedStatusCode int, expectedBody string) {
	getStatusCode, _, body := ts.get(t, urlPath)
	assert.Equal(t, expectedStatusCode, getStatusCode)
	require.JSONEq(t, expectedBody, body)
}
