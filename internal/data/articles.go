package data

// Article represents a single article in the system.
type Article struct {
	ID    int64       `json:"id"`
	Title string      `json:"title"`
	Date  ArticleDate `json:"date"`
	Body  string      `json:"body"`
	Tags  []string    `json:"tags"`
}
