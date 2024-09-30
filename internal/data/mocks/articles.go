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
		"2021-01-01", "2021-01-02", "2021-01-03", "2022-05-01",
		"2022-06-15", "2022-07-20", "2022-08-30", "2016-09-22",
		"2016-09-22", "2016-09-22", "2016-09-22", "2016-09-22",
		"2016-09-22", "2016-09-22", "2016-09-22", "2016-09-22",
		"2016-09-22", "2016-09-22", "2016-09-22", "2016-09-22",
		"2016-09-22", "2016-09-22", "2016-09-22",
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
		{
			ID:    12,
			Title: "health benefits of a balanced diet",
			Date:  parsedDates[11],
			Body:  "a balanced diet is crucial for good health",
			Tags:  []string{"health"},
		},
		{
			ID:    13,
			Title: "how to manage stress for better health",
			Date:  parsedDates[12],
			Body:  "managing stress is important for maintaining good health",
			Tags:  []string{"health"},
		},
		{
			ID:    14,
			Title: "health benefits of regular check-ups",
			Date:  parsedDates[13],
			Body:  "regular check-ups can help you stay healthy",
			Tags:  []string{"health"},
		},
		{
			ID:    15,
			Title: "how to boost your immune system",
			Date:  parsedDates[14],
			Body:  "boosting your immune system is important for good health",
			Tags:  []string{"health"},
		},
		{
			ID:    16,
			Title: "health benefits of meditation",
			Date:  parsedDates[15],
			Body:  "meditation can improve your overall health",
			Tags:  []string{"health"},
		},
		{
			ID:    17,
			Title: "how to stay healthy while traveling",
			Date:  parsedDates[16],
			Body:  "tips on staying healthy while traveling",
			Tags:  []string{"health"},
		},
		{
			ID:    18,
			Title: "new species of bird found",
			Date:  parsedDates[17],
			Body:  "a new species of bird has been found in the pacific",
			Tags:  []string{"science"},
		},
		{
			ID:    19,
			Title: "new year's resolutions for better health",
			Date:  parsedDates[18],
			Body:  "setting new year's resolutions can help improve your health",
			Tags:  []string{"health"},
		},
		{
			ID:    20,
			Title: "how to stay healthy during winter",
			Date:  parsedDates[19],
			Body:  "tips on staying healthy during the cold months",
			Tags:  []string{"health", "fitness", "nutrition", "wellness", "exercise", "diet"},
		},
		{
			ID:    21,
			Title: "healthy eating habits",
			Date:  parsedDates[20],
			Body:  "developing healthy eating habits is crucial",
			Tags:  []string{"health", "self-care", "hydration", "sleep", "stress management"},
		},
		{
			ID:    22,
			Title: "importance of mental health",
			Date:  parsedDates[21],
			Body:  "mental health is just as important as physical health",
			Tags:  []string{"health", "mental health", "yoga", "meditation", "lifestyle"},
		},
		{
			ID:    23,
			Title: "health benefits of yoga",
			Date:  parsedDates[22],
			Body:  "yoga can improve your overall health",
			Tags:  []string{"health", "mindfulness", "mobility"},
		},
		{
			ID:    24,
			Title: "how to maintain a healthy lifestyle",
			Date:  parsedDates[23],
			Body:  "tips on maintaining a healthy lifestyle",
			Tags:  []string{"health"},
		},
		{
			ID:    25,
			Title: "health benefits of drinking water",
			Date:  parsedDates[24],
			Body:  "drinking water is essential for good health",
			Tags:  []string{"health"},
		},
		{
			ID:    26,
			Title: "how to improve your health with exercise",
			Date:  parsedDates[25],
			Body:  "exercise is key to improving your health",
			Tags:  []string{"health"},
		},
		{
			ID:    27,
			Title: "health benefits of a balanced diet",
			Date:  parsedDates[26],
			Body:  "a balanced diet is crucial for good health",
			Tags:  []string{"health"},
		},
	}
}
