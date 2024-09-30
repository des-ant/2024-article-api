package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/des-ant/2024-article-api/internal/data/mocks"
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

func TestCreateArticleValidation(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name           string
		data           map[string]interface{}
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Valid Article",
			data: map[string]interface{}{
				"id":    1,
				"title": "latest science shows that potato chips are better for you than sugar",
				"date":  "2016-09-22",
				"body":  "some text, potentially containing simple markup about how potato chip",
				"tags":  []string{"health", "fitness", "science"},
			},
			expectedStatus: http.StatusCreated,
			expectedBody: `{
							"article": {
									"id": 1,
									"title": "latest science shows that potato chips are better for you than sugar",
									"date": "2016-09-22",
									"body": "some text, potentially containing simple markup about how potato chip",
									"tags": ["health", "fitness", "science"]
							}
					}`,
		},
		{
			name: "Missing Title",
			data: map[string]interface{}{
				"id":   2,
				"date": "2016-09-22",
				"body": "some text, potentially containing simple markup about how potato chip",
				"tags": []string{"health", "fitness", "science"},
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   `{"error": {"title": "must be provided"}}`,
		},
		{
			name: "Title Too Long",
			data: map[string]interface{}{
				"id":    3,
				"title": string(make([]byte, 501)),
				"date":  "2016-09-22",
				"body":  "some text, potentially containing simple markup about how potato chip",
				"tags":  []string{"health", "fitness", "science"},
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   `{"error": {"title": "must not be more than 500 bytes long"}}`,
		},
		{
			name: "Missing Body",
			data: map[string]interface{}{
				"id":    4,
				"title": "latest science shows that potato chips are better for you than sugar",
				"date":  "2016-09-22",
				"tags":  []string{"health", "fitness", "science"},
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   `{"error": {"body": "must be provided"}}`,
		},
		{
			name: "Missing Tags",
			data: map[string]interface{}{
				"id":    5,
				"title": "latest science shows that potato chips are better for you than sugar",
				"date":  "2016-09-22",
				"body":  "some text, potentially containing simple markup about how potato chip",
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   `{"error": {"tags": "must be provided"}}`,
		},
		{
			name: "Not Enough Tags",
			data: map[string]interface{}{
				"id":    6,
				"title": "latest science shows that potato chips are better for you than sugar",
				"date":  "2016-09-22",
				"body":  "some text, potentially containing simple markup about how potato chip",
				"tags":  []string{},
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   `{"error": {"tags": "must contain at least 1 tag"}}`,
		},
		{
			name: "Too Many Tags",
			data: map[string]interface{}{
				"id":    7,
				"title": "latest science shows that potato chips are better for you than sugar",
				"date":  "2016-09-22",
				"body":  "some text, potentially containing simple markup about how potato chip",
				"tags":  []string{"tag1", "tag2", "tag3", "tag4", "tag5", "tag6", "tag7", "tag8", "tag9", "tag10", "tag11"},
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   `{"error": {"tags": "must not contain more than 10 tags"}}`,
		},
		{
			name: "Duplicate Tags",
			data: map[string]interface{}{
				"id":    8,
				"title": "latest science shows that potato chips are better for you than sugar",
				"date":  "2016-09-22",
				"body":  "some text, potentially containing simple markup about how potato chip",
				"tags":  []string{"health", "health"},
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   `{"error": {"tags": "must not contain duplicate values"}}`,
		},
		{
			name: "Invalid ID",
			data: map[string]interface{}{
				"id":    -1,
				"title": "latest science shows that potato chips are better for you than sugar",
				"date":  "2016-09-22",
				"body":  "some text, potentially containing simple markup about how potato chip",
				"tags":  []string{"health", "fitness", "science"},
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   `{"error": {"id": "must be a positive integer"}}`,
		},
		{
			name:           "Missing Data",
			data:           map[string]interface{}{},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   `{"error": {"id": "must be a positive integer", "title": "must be provided", "body": "must be provided", "tags": "must be provided", "date": "must be provided and valid"}}`,
		},
		{
			name: "Invalid Date",
			data: map[string]interface{}{
				"id":    9,
				"title": "latest science shows that potato chips are better for you than sugar",
				"date":  "invalid",
				"body":  "some text, potentially containing simple markup about how potato chip",
				"tags":  []string{"health", "fitness", "science"},
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"invalid date format"}`,
		},
		{
			name: "Empty Data",
			data: map[string]interface{}{
				"id":    0,
				"title": "",
				"date":  "",
				"body":  "",
				"tags":  []string{},
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"invalid date format"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postStatusCode, _, body := ts.postJSON(t, "/v1/articles", tt.data)
			assert.Equal(t, tt.expectedStatus, postStatusCode)
			require.JSONEq(t, tt.expectedBody, body)
		})
	}
}

func TestShowArticleHandler(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	mockArticle := map[string]interface{}{
		"id":    1,
		"title": "latest science shows that potato chips are better for you than sugar",
		"date":  "2016-09-22",
		"body":  "some text, potentially containing simple markup about how potato chip",
		"tags":  []string{"health", "fitness", "science"},
	}

	statusCode, _, _ := ts.postJSON(t, "/v1/articles", mockArticle)
	if statusCode != http.StatusCreated {
		t.Fatalf("expected status %d; got %d", http.StatusCreated, statusCode)
	}

	tests := []struct {
		name           string
		id             string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid ID",
			id:             "1",
			expectedStatus: http.StatusOK,
			expectedBody: `{
							"article": {
									"id": 1,
									"title": "latest science shows that potato chips are better for you than sugar",
									"date": "2016-09-22",
									"body": "some text, potentially containing simple markup about how potato chip",
									"tags": ["health", "fitness", "science"]
							}
					}`,
		},
		{
			name:           "Non-existent ID",
			id:             "999",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"the requested resource could not be found"}`,
		},
		{
			name:           "Invalid ID",
			id:             "abc",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"the requested resource could not be found"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("/v1/articles/%s", tt.id)
			statusCode, _, body := ts.get(t, url)
			assert.Equal(t, tt.expectedStatus, statusCode)
			require.JSONEq(t, tt.expectedBody, body)
		})
	}
}

func TestGetArticlesByTagAndDateHandler(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	mockArticles := mocks.InitMockArticles()

	for _, article := range mockArticles {
		articleMap := map[string]interface{}{
			"id":    article.ID,
			"title": article.Title,
			"date":  article.Date.String(),
			"body":  article.Body,
			"tags":  article.Tags,
		}
		statusCode, _, _ := ts.postJSON(t, "/v1/articles", articleMap)
		if statusCode != http.StatusCreated {
			t.Fatalf("expected status %d; got %d", http.StatusCreated, statusCode)
		}
	}

	tests := []struct {
		name           string
		tagName        string
		date           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid Tag and Date",
			tagName:        "health",
			date:           "20160922",
			expectedStatus: http.StatusOK,
			expectedBody: `{
					"tag_summary": {
							"tag": "health",
							"count": 17,
							"articles": [17, 19, 20, 21, 22, 23, 24, 25, 26, 27],
							"related_tags": ["diet", "exercise", "fitness", "hydration", "lifestyle", "meditation", "mental health", "mindfulness", "mobility", "nutrition", "science", "self-care", "sleep", "stress management", "wellness", "yoga"]
					}
			}`,
		},
		{
			name:           "Non-existent Tag",
			tagName:        "nonexistent",
			date:           "20160922",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"the requested resource could not be found"}`,
		},
		{
			name:           "Invalid Date",
			tagName:        "health",
			date:           "invalid-date",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"the requested resource could not be found"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("/v1/tags/%s/%s", tt.tagName, tt.date)
			statusCode, _, body := ts.get(t, url)
			assert.Equal(t, tt.expectedStatus, statusCode)

			compareJSONBodies(t, tt.expectedBody, body)
		})
	}
}
