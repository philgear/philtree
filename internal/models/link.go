package models

// Link represents a generic link structure
type Link struct {
	Title       string
	URL         string
	Description string
	Icon        string
}

// Website represents one of your main websites
type Website struct {
	Link
	Category string
}

// SocialMediaLink represents a social media profile link
type SocialMediaLink struct {
	Link
	Platform string
}

// NewWebsite creates a new Website instance
func NewWebsite(title, url, description, category string) Website {
	return Website{
		Link: Link{
			Title:       title,
			URL:         url,
			Description: description,
			Icon:        "website-icon", // You can customize this
		},
		Category: category,
	}
}

// NewSocialMediaLink creates a new SocialMediaLink instance
func NewSocialMediaLink(title, url, platform string) SocialMediaLink {
	return SocialMediaLink{
		Link: Link{
			Title: title,
			URL:   url,
			Icon:  platform + "-icon", // e.g., "twitter-icon"
		},
		Platform: platform,
	}
}