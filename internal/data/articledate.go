package data

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

var ErrInvalidArticleDateFormat = errors.New("invalid date format")

// ArticleDate is a custom type for handling dates in the format "YYYY-MM-DD".
type ArticleDate time.Time

// MarshalJSON implements the json.Marshaler interface.
// It will encode the time as a string in the format "2006-01-02".
func (ad ArticleDate) MarshalJSON() ([]byte, error) {
	formattedTime := time.Time(ad).Format("2006-01-02")
	return json.Marshal(formattedTime)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// It expects the time to be a string in the format "2006-01-02".
// We use a pointer receiver because UnmarshalJSON modifies the receiver.
func (ad *ArticleDate) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidArticleDateFormat
	}

	parsedTime, err := time.Parse("2006-01-02", unquotedJSONValue)
	if err != nil {
		return ErrInvalidArticleDateFormat
	}

	*ad = ArticleDate(parsedTime)

	return nil
}

// ToTime converts ArticleDate back to time.Time.
func (ad ArticleDate) ToTime() time.Time {
	return time.Time(ad)
}

// String implements the fmt.Stringer interface.
func (ad ArticleDate) String() string {
	return time.Time(ad).Format("2006-01-02")
}
