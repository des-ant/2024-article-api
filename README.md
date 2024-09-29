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
$ curl -d '{
  "title": "latest science shows that potato chips are better for you than sugar",
  "date": "2016-09-22",
  "body": "some text, potentially containing simple markup about how potato chip",
  "tags": ["health", "fitness", "science"]
}' -H "Content-Type: application/json" localhost:4000/v1/articles
{Title:latest science shows that potato chips are better for you than sugar Date:2016-09-22 Body:some text, potentially containing simple markup about how potato chip Tags:[health fitness science]}

$ curl localhost:4000/v1/articles/123
{
	"article": {
		"id": 123,
		"title": "Article Title",
		"date": "2024-09-29",
		"body": "This is the body of the article.",
		"tags": [
			"tag1",
			"tag2",
			"tag3"
		]
	}
}

$ curl -i localhost:4000/v1/healthcheck
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
