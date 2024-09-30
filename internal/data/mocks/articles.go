package mocks

import (
	"log"

	"github.com/des-ant/2024-article-api/internal/data"
)

// initMockArticles initializes a set of mock articles for testing purposes.
func InitMockArticles() []*data.Article {
	// Duplicate values to represent separate articleDates created with different requests.
	dateStrings := []string{
		"2016-09-22", "2016-09-22", "2016-09-22", "2016-09-23",
		"2021-01-01", "2021-01-02", "2021-01-03",
		"2022-05-01", "2022-06-15", "2022-07-20", "2022-08-30",
	}

	parsedDates, err := data.ParseArticleDates(dateStrings)
	if err != nil {
		log.Fatalf("Error parsing dates: %v", err)
	}

	return []*data.Article{
		{
			ID:    1,
			Title: "latest science shows that potato chips are better for you than sugar",
			Date:  parsedDates[0],
			Body:  "some text, potentially containing simple markup about how potato chip",
			Tags:  []string{"health", "fitness", "science"},
		},
		{
			ID:    2,
			Title: "breakthrough in sleep science",
			Date:  parsedDates[1],
			Body:  "scientists have discovered a new way to help you fall asleep faster",
			Tags:  []string{"health", "lifestyle", "science"},
		},
		{
			ID:    3,
			Title: "new species of bird found",
			Date:  parsedDates[2],
			Body:  "a new species of bird has been found in the pacific",
			Tags:  []string{"biology", "animals", "science"},
		},
		{
			ID:    4,
			Title: "olympics are coming",
			Date:  parsedDates[3],
			Body:  "the olympics are a time of great excitement for many people",
			Tags:  []string{"sports", "entertainment", "world"},
		},
		{
			ID:    5,
			Title: "Hello, World!",
			Date:  parsedDates[4],
			Body:  "This is the first article in the system.",
			Tags:  []string{"welcome", "first"},
		},
		{
			ID:    6,
			Title: "A New Beginning",
			Date:  parsedDates[5],
			Body:  "This is the second article in the system.",
			Tags:  []string{"welcome", "second"},
		},
		{
			ID:    7,
			Title: "The Final Article",
			Date:  parsedDates[6],
			Body:  "This is the third article in the system.",
			Tags:  []string{"welcome", "third"},
		},
		{
			ID:    8,
			Title: "Advancements in AI Technology",
			Date:  parsedDates[7],
			Body:  "Recent advancements in AI technology have shown promising results in various fields.",
			Tags:  []string{"technology", "AI", "innovation"},
		},
		{
			ID:    9,
			Title: "Climate Change and Its Impact",
			Date:  parsedDates[8],
			Body:  "Climate change continues to have a significant impact on the environment and human life.",
			Tags:  []string{"environment", "climate change", "science"},
		},
		{
			ID:    10,
			Title: "Exploring the Depths of the Ocean",
			Date:  parsedDates[9],
			Body:  "Scientists have made new discoveries while exploring the depths of the ocean.",
			Tags:  []string{"oceanography", "science", "exploration"},
		},
		{
			ID:    11,
			Title: "The Future of Space Travel",
			Date:  parsedDates[10],
			Body:  "Space agencies are planning new missions to explore the outer reaches of our solar system.",
			Tags:  []string{"space", "technology", "exploration"},
		},
	}
}
