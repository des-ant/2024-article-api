package data

import (
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

// DAOs represents a collection of data access objects.
type DAOs struct {
	Articles *ArticleDAO
}

// NewDAOs creates a new instance of DAOs.
func NewDAOs() *DAOs {
	return &DAOs{
		Articles: NewArticleDAO(),
	}
}
