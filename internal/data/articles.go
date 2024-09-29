package data

import (
	"github.com/des-ant/2024-article-api/internal/validator"
)

// Article represents a single article in the system.
type Article struct {
	ID    int64       `json:"id"`
	Title string      `json:"title"`
	Date  ArticleDate `json:"date"`
	Body  string      `json:"body"`
	Tags  []string    `json:"tags"`
}

// ValidateArticle validates the provided Article struct and adds an error message
// to the validator instance if any of the validation rules fail.
func ValidateArticle(v *validator.Validator, article *Article) {
	v.Check(article.Title != "", "title", "must be provided")
	v.Check(len(article.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(article.Body != "", "body", "must be provided")

	v.Check(article.Tags != nil, "tags", "must be provided")
	v.Check(len(article.Tags) >= 1, "tags", "must contain at least 1 tag")
	v.Check(len(article.Tags) <= 10, "tags", "must not contain more than 10 tags")
	v.Check(validator.Unique(article.Tags), "tags", "must not contain duplicate values")
}
