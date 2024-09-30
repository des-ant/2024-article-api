package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"sort"
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

// sortArticlesAndTags sorts the "articles" and "related_tags" slices in the tag_summary map.
// JSON marshalling does not guarantee the order of slices, so we need to sort them to compare them in tests.
func sortArticlesAndTags(bodyMap map[string]interface{}) {
	// Check if the bodyMap contains a "tag_summary" key.
	// This ensures that we only attempt to sort if the key exists and is of the correct type.
	tagSummary, ok := bodyMap["tag_summary"].(map[string]interface{})
	if !ok {
		return
	}

	// Sort the "articles" slice if it exists.
	// We use type assertion to check if "articles" is a slice of empty interfaces.
	if articles, ok := tagSummary["articles"].([]interface{}); ok {
		sort.Slice(articles, func(i, j int) bool {
			// Convert the articles to float64 to compare them.
			// JSON unmarshalling represents all numbers as float64.
			return articles[i].(float64) < articles[j].(float64)
		})
	}

	// Sort the "related_tags" slice if it exists.
	// We use type assertion to check if "related_tags" is a slice of empty interfaces.
	if relatedTags, ok := tagSummary["related_tags"].([]interface{}); ok {
		sort.Slice(relatedTags, func(i, j int) bool {
			// Convert the related tags to strings to compare them.
			// JSON unmarshalling represents all strings as string.
			return relatedTags[i].(string) < relatedTags[j].(string)
		})
	}
}

// compareJSONBodies compares the expected and actual JSON bodies, ignoring the order of slices.
func compareJSONBodies(t *testing.T, expectedBody, actualBody string) {
	var expectedBodyMap, actualBodyMap map[string]interface{}

	// Unmarshal the expected JSON body into a map.
	// This allows us to work with the JSON data in a structured way.
	err := json.Unmarshal([]byte(expectedBody), &expectedBodyMap)
	require.NoError(t, err)

	// Unmarshal the actual JSON body into a map.
	// This allows us to work with the JSON data in a structured way.
	err = json.Unmarshal([]byte(actualBody), &actualBodyMap)
	require.NoError(t, err)

	// Sort the articles and related_tags arrays before comparison.
	// This ensures that the order of elements does not affect the comparison result.
	sortArticlesAndTags(expectedBodyMap)
	sortArticlesAndTags(actualBodyMap)

	// Compare the expected and actual JSON bodies.
	// This checks if the two JSON bodies are equivalent after sorting.
	assert.Equal(t, expectedBodyMap, actualBodyMap)
}
