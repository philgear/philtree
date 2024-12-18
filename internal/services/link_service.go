// Package services provides business logic for the application
package services

import "github.com/philgear/philtree/internal/config"

// LinkService holds data for websites, social media, and affiliate links
type LinkService struct {
	Websites       []Website
	SocialMedia    []SocialMedia
	AffiliateLinks []AffiliateLink
}

// NewLinkService creates a new LinkService instance from the given configuration
func NewLinkService(cfg *config.Config) *LinkService {
	return &LinkService{
		Websites:       convertWebsites(cfg.Websites),
		SocialMedia:    convertSocialMedia(cfg.SocialMedia),
		AffiliateLinks: convertAffiliateLinks(cfg.AffiliateLinks),
	}
}

func convertWebsites(cfgWebsites []config.Website) []Website {
	websites := make([]Website, len(cfgWebsites))
	for i, w := range cfgWebsites {
		websites[i] = Website{
			Name:        w.Name,
			URL:         w.URL,
			LastUpdate:  w.LastUpdate.Time,
			Description: w.Description,
		}
	}
	return websites
}

func convertSocialMedia(cfgSocialMedia []config.SocialMedia) []SocialMedia {
	socialMedia := make([]SocialMedia, len(cfgSocialMedia))
	for i, s := range cfgSocialMedia {
		socialMedia[i] = SocialMedia{
			Platform:    s.Platform,
			URL:         s.URL,
			LastUpdate:  s.LastUpdate.Time,
			Description: s.Description,
		}
	}
	return socialMedia
}

func convertAffiliateLinks(cfgAffiliateLinks []config.AffiliateLink) []AffiliateLink {
	affiliateLinks := make([]AffiliateLink, len(cfgAffiliateLinks))
	for i, a := range cfgAffiliateLinks {
		affiliateLinks[i] = AffiliateLink(a)
	}
	return affiliateLinks
}

// FormattedLastUpdate returns the LastUpdate formatted as a string
func (w Website) FormattedLastUpdate() string {
	return w.LastUpdate.Format("2006-01-02")
}

// FormattedLastUpdate returns the LastUpdate formatted as a string
func (s SocialMedia) FormattedLastUpdate() string {
	return s.LastUpdate.Format("2006-01-02")
}
