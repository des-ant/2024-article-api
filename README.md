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

curl localhost:8080/v1/tags/health/20160922
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

<!-- Q&A -->
## :grey_question: Q&A

- What languages and libraries did you use and why?

  + Aside from the recommendation, I used Go because it is a statically typed language that is easy to read and write. It also has a rich standard library that makes it easy to build web servers.
  + I also used the `httprouter` library because:
    + It is a lightweight and fast HTTP request router that is easy to use.
    + I wanted the API to consistenly send JSON responses wherever possible -
      the standard library `http.ServeMux` sends plaintext (non-JSON) responses
      `404` and `405` responses when a matching route is not found.
    + `http.ServeMux` does not automatically handle `OPTIONS` requests, which
      are required for CORS preflight requests.
    + `httprouter` supports both of the above features and is stable and well-tested.
    + However, because `httprouter` is minimalistic, I had to implement some
      features myself, such as handling invalid routes and requests. This was
      good for learning purposes but added complexity and prolonged development
      time.
  + I used the `testify` library for testing because it provides useful
    assertion functions that make tests easier to read and write.
  + For the data store, I had initially planned to use PostgreSQL but due to time
    constraints, I used an in-memory store. However, I structured the code so
    that the in-memory store can be easily replaced with any other data store
    e.g. PostgreSQL.

- Why did you structure the project the way you did?

```plaintext
2024-article-api
├─ cmd/
│  ├─ api/
├─ internal/
│  ├─ data/
│  ├─ validator/
├─ vendor/
```

  + The application was structured this way to separate concerns and make the
    codebase easier to navigate and maintain.
  + The `cmd/api` directory contains the main application code, including the
    server setup and configuration.
  + The `internal` directory contains the application's internal packages, such
    as the data store and request validator. We import these packages in
    `cmd/api`.
  + The `vendor` directory contains the project's dependencies. It is
    technically not necessary but I chose to include it for a few reasons:
    + It ensures the project builds without external dependencies.
    + It speeds up CI/CD jobs by avoiding dependency downloads.
    + It did not add significant overhead to the project and the project is not
      intended to be a library.
    + It can improve the code review process by separating dependency changes.

- How did you approach testing?

  + For testing, I primarily wanted to cover the end-to-end flow of the
    endpoints mentioned in the requirements and ensure the API was functioning from a user's perspective. I primarily used integration testing to test the interaction between different parts of the application, such as the HTTP server, routes, handlers, and application logic, in a realistic environment.

- What additional features did you add and why?

  + I added API versioning to the routes to ensure backward compatibility with
    future versions of the API. This is important for maintaining the API and
    ensuring that clients can continue to use the API without breaking changes.
  + I added a healthcheck endpoint to return the status of the server and some
    system information. This is useful for monitoring the server and diagnosing
    issues.
  + I added a `Makefile` to simplify common development tasks, such as building,
    running, and testing the application. This makes it easier to work on the
    project and ensures consistency across different environments.
  + I handled invalid routes and requests to provide a better user experience
    and prevent potential security vulnerabilities. This is important for
    ensuring the API is robust and secure.
    + For example I limited the size of the request body to prevent denial of
      service attacks, optimise performance and prevent memory exhaustion.
  + I used Docker to containerize the application so it could be run in any
    environment without additional setup. This also makes it easier to deploy to
    a cloud provider like AWS or GCP.
  + I added middleware to handle panics during request handling. This ensures
    that instead of just closing the connection, a `500 Internal Server Error`
    response is sent to the client, providing context for the error.
  + I also enveloped the JSON response by nesting data under a key like
    `"article"`. This makes the response more self-documenting, reduces
    client-side errors, and mitigates certain security vulnerabilities in older
    browsers.

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
