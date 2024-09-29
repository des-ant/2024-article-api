package data

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
