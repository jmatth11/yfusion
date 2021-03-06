package yfusion

import "time"

// ReviewsData - Represents the return data from a business review request
type ReviewsData struct {
	Total             int
	PossibleLanguages []string `json:"possible_languages"`
	Reviews           []*ReviewInfo
	Error             *BusinessMigratedError
}

// ReviewInfo - Represents review info data inside of the return data from review request
type ReviewInfo struct {
	ID          string
	Rating      int
	User        UserInfo
	Text        string
	TimeCreated *time.Time `json:"time_created"`
	URL         string
}

// UserInfo - Represents user info data inside of a review
type UserInfo struct {
	ID         string
	ProfileURL string `json:"profile_url"`
	ImageURL   string `json:"image_url"`
	Name       string
}
