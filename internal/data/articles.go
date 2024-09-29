package data

import (
	"errors"
	"sync"

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

// TagSummary represents a summary of tags for a given article.
type TagSummary struct {
	Tag         string   `json:"tag"`
	Count       int      `json:"count"`
	Articles    []int64  `json:"articles"`
	RelatedTags []string `json:"related_tags"`
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

// ArticleDAO represents the data access object for articles.
type ArticleDAO struct {
	articles map[int64]Article
	mutex    sync.RWMutex
}

// NewArticleDAO creates a new instance of ArticleDAO.
func NewArticleDAO() *ArticleDAO {
	return &ArticleDAO{
		articles: make(map[int64]Article),
	}
}

// Insert adds a new article to the store.
func (dao *ArticleDAO) Insert(article *Article) error {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()

	if _, exists := dao.articles[article.ID]; exists {
		return errors.New("duplicate key, article with ID already exists")
	}

	dao.articles[article.ID] = *article

	return nil
}

// Get retrieves an article by ID.
func (dao *ArticleDAO) Get(id int64) (*Article, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	dao.mutex.RLock()
	defer dao.mutex.RUnlock()

	article, exists := dao.articles[id]
	if !exists {
		return nil, ErrRecordNotFound
	}

	return &article, nil
}

// GetArticlesByTagAndDate retrieves articles by tag and date.
func (dao *ArticleDAO) GetArticlesByTagAndDate(tag string, date ArticleDate) ([]Article, error) {
	dao.mutex.RLock()
	defer dao.mutex.RUnlock()

	var result []Article
	for _, article := range dao.articles {
		if article.Date == date && contains(article.Tags, tag) {
			result = append(result, article)
		}
	}

	if len(result) == 0 {
		return nil, errors.New("no articles found for the given tag and date")
	}

	return result, nil
}

// GetRelatedTags retrieves related tags from a list of articles.
func (dao *ArticleDAO) GetRelatedTags(articles []Article) []string {
	tagSet := make(map[string]struct{})
	for _, article := range articles {
		for _, tag := range article.Tags {
			tagSet[tag] = struct{}{}
		}
	}

	var relatedTags []string
	for tag := range tagSet {
		relatedTags = append(relatedTags, tag)
	}

	return relatedTags
}

// Helper function to check if a slice contains a string.
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
