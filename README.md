# 2024-article-api

<!-- Run Locally -->
### :running: Run Locally

Clone the project

```bash
git clone https://github.com/des-ant/2024-article-api.git
```

Go to the project directory

```bash
cd 2024-article-api
```

Start the server

```bash
$ go run ./cmd/api
time=2024-09-28T15:13:02.746+10:00 level=INFO msg="starting server" addr=:4000 env=development
```


<!-- Usage -->
## :eyes: Usage

While the server is running, visit [localhost:4000/v1/healthcheck](localhost:4000/v1/healthcheck) in your web browser.

Alternatively, use `curl` to make the request from a terminal:
```bash
$ curl -i localhost:4000/v1/healthcheck
HTTP/1.1 200 OK
Date: Sat, 28 Sep 2024 05:25:21 GMT
Content-Length: 58
Content-Type: text/plain; charset=utf-8

status: available
environment: development
version: 1.0.0

$ curl -X POST localhost:4000/v1/articles
create a new article

$ curl localhost:4000/v1/articles/123
{
	"article": {
		"id": 123,
		"title": "Article Title",
		"date": "2024-09-29T11:39:12.858097+10:00",
		"body": "This is the body of the article.",
		"tags": [
			"tag1",
			"tag2",
			"tag3"
		]
	}
}

$ curl localhost:4000/v1/healthcheck
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
  * [ ] Create routes
    * [x] POST `/articles`
    * [x] GET `/articles/{id}`
    * [ ] GET `/tags/{tagName}/{date}`
  * [ ] Implement handler logic
    * [ ] POST `/articles`
    * [ ] GET `/articles/{id}`
    * [ ] GET `/tags/{tagName}/{date}`
  * [ ] Create data models
    * [ ] Create `Article` struct
    * [ ] Create `TagSummary` struct
  * [ ] Set up data store
  * [ ] Implement error handling for:
    * [x] Invalid routes
    * [ ] Invalid requests
      * [ ] POST `/articles`
      * [ ] GET `/articles/{id}`
      * [ ] GET `/tags/{tagName}/{date}`
    * [x] Panics
