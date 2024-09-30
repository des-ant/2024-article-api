# 2024-article-api

<!-- About the Project -->
2024-article-api is a Go-based API server designed to manage and serve articles efficiently.


<!-- Getting Started -->
## 	:toolbox: Getting Started

<!-- Prerequisites -->
### :bangbang: Prerequisites

#### Running the Application
Only [Docker](https://www.docker.com/) is required to run the application

#### Development
For development, the application requires:
* [Go (version 1.23.1)](https://go.dev/dl/)
* [Make](https://www.gnu.org/software/make/)
* [golangci-lint](https://golangci-lint.run/welcome/install/)

<!-- Run Locally -->
### :running: Run Locally

Clone the project using the following command:

```bash
git clone https://github.com/des-ant/2024-article-api.git
```

Move into the project directory:

```bash
cd 2024-article-api
```

If you have Docker installed, you can build and run the application using the
following commands:

```bash
docker build -t 2024-article-api .
docker run -p 8080:8080 2024-article-api --port=8080 --env=development
```

Otherwise, if you have Go and Make installed, you can run the application using
the following commands:

```bash
make run/api port=8080 env=development
```

> **_NOTE:_**: The `port` and `env` flags are optional. The default port is `4000` and the default environment is `development`.


<!-- Running Tests -->
### :test_tube: Running Tests

To run tests, run the following command

```bash
make test
```

To see other available `make` commands, run

```bash
make help
```

<!-- Usage -->
## :eyes: Usage

While the server is running, visit
[localhost:8080/v1/healthcheck](localhost:8080/v1/healthcheck) in your web
browser to check the status of the server.

Here are some example requests you can make to the server, using `curl`:
```bash
curl -d '{
  "id": 1,
  "title": "latest science shows that potato chips are better for you than sugar",
  "date": "2016-09-22",
  "body": "some text, potentially containing simple markup about how potato chip",
  "tags": ["health", "fitness", "science"]
}' -H "Content-Type: application/json" localhost:8080/v1/articles
{
	"article": {
		"id": 1,
		"title": "latest science shows that potato chips are better for you than sugar",
		"date": "2016-09-22",
		"body": "some text, potentially containing simple markup about how potato chip",
		"tags": [
			"health",
			"fitness",
			"science"
		]
	}
}

curl localhost:8080/v1/articles/1
{
	"article": {
		"id": 1,
		"title": "latest science shows that potato chips are better for you than sugar",
		"date": "2016-09-22",
		"body": "some text, potentially containing simple markup about how potato chip",
		"tags": [
			"health",
			"fitness",
			"science"
		]
	}
}

curl -d '{
  "id": 2,
  "title": "breakthrough in sleep science",
  "date": "2016-09-22",
  "body": "scientists have discovered a new way to help you fall asleep faster",
  "tags": ["health", "lifestyle", "science"]
}' -H "Content-Type: application/json" localhost:8080/v1/articles
{
	"article": {
		"id": 2,
		"title": "breakthrough in sleep science",
		"date": "2016-09-22",
		"body": "scientists have discovered a new way to help you fall asleep faster",
		"tags": [
			"health",
			"lifestyle",
			"science"
		]
	}
}

$ curl localhost:8080/v1/tags/health/20160922
{
	"tag_summary": {
		"tag": "health",
		"count": 4,
		"articles": [
			2,
			1
		],
		"related_tags": [
			"science",
			"fitness",
			"lifestyle"
		]
	}
}

curl -i localhost:8080/v1/healthcheck
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 29 Sep 2024 06:20:28 GMT
Content-Length: 102

{
	"status": "available",
	"system_info": {
		"environment": "development",
		"version": "1.0.0"
	}
}
```



<!-- Roadmap -->
## :compass: Roadmap

* [x] Set up project structure
* [x] Create basic HTTP Server
* [ ] Add API endpoints
  * [x] Create routes
    * [x] POST `/articles`
    * [x] GET `/articles/{id}`
    * [x] GET `/tags/{tagName}/{date}`
  * [x] Implement handler logic
    * [x] POST `/articles`
    * [x] GET `/articles/{id}`
    * [x] GET `/tags/{tagName}/{date}`
  * [ ] Create data models
    * [x] Create `Article` struct
    * [x] Create `TagSummary` struct
  * [ ] Set up data store
    * [x] Use in-memory store
    * [ ] Replace in-memory store with PostgreSQL DB
  * [ ] Implement error handling for:
    * [x] Invalid routes
    * [x] Invalid requests
      * [x] POST `/articles`
      * [x] GET `/articles/{id}`
      * [x] GET `/tags/{tagName}/{date}`
    * [x] Panics
* [ ] Add tests
  * [x] POST `/articles`
  * [x] GET `/articles/{id}`
  * [x] GET `/tags/{tagName}/{date}`
