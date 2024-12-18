package services

import (
	"github.com/philgear/philtree/internal/config"
	"github.com/philgear/philtree/internal/models"
)

// LinkService handles operations related to links
type LinkService struct {
	cfg *config.Config
}

// NewLinkService creates a new instance of LinkService
func NewLinkService(cfg *config.Config) *LinkService {
	return &LinkService{cfg: cfg}
}

// GetWebsites returns all website links
func (s *LinkService) GetWebsites() []models.Website {
	websites := make([]models.Website, len(s.cfg.Websites))
	for i, w := range s.cfg.Websites {
		websites[i] = models.NewWebsite(w.Name, w.URL, "", "Main") // Add a description if available in your config
	}
	return websites
}

// GetSocialMediaLinks returns all social media links
func (s *LinkService) GetSocialMediaLinks() []models.SocialMediaLink {
	socialLinks := make([]models.SocialMediaLink, len(s.cfg.SocialMedia))
	for i, sm := range s.cfg.SocialMedia {
		socialLinks[i] = models.NewSocialMediaLink(sm.Platform, sm.URL, sm.Platform)
	}
	return socialLinks
}

// GetAffiliateLinks returns all affiliate links
func (s *LinkService) GetAffiliateLinks() []models.AffiliateLink {
	affiliateLinks := make([]models.AffiliateLink, len(s.cfg.AffiliateLinks))
	for i, al := range s.cfg.AffiliateLinks {
		affiliateLinks[i] = models.NewAffiliateLink(al.Name, al.URL, "", al.Name) // Add a description if available in your config
	}
	return affiliateLinks
}

// GetAllLinks returns all links categorized
func (s *LinkService) GetAllLinks() map[string]interface{} {
	return map[string]interface{}{
		"websites":       s.GetWebsites(),
		"socialMedia":    s.GetSocialMediaLinks(),
		"affiliateLinks": s.GetAffiliateLinks(),
	}
}
