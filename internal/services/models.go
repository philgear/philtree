package services

import "time"

// Website represents a website entry
type Website struct {
	Name        string
	URL         string
	LastUpdate  time.Time
	Description string
}

// SocialMedia represents a social media entry
type SocialMedia struct {
	Platform    string
	URL         string
	LastUpdate  time.Time
	Description string
}

// AffiliateLink represents an affiliate link entry
type AffiliateLink struct {
	Name string
	URL  string
}
