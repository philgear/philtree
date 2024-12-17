package services

import (
	"github.com/philgear/philtree/internal/config"
)

type Link struct {
	Title string
	URL   string
}

type LinkService struct {
	cfg *config.Config
}

func NewLinkService(cfg *config.Config) *LinkService {
	return &LinkService{cfg: cfg}
}

func (ls *LinkService) GetWebsites() []Link {
	websites := make([]Link, len(ls.cfg.Websites))
	for i, w := range ls.cfg.Websites {
		websites[i] = Link{Title: w.Name, URL: w.URL}
	}
	return websites
}

func (ls *LinkService) GetSocialMediaLinks() []Link {
	socialLinks := make([]Link, len(ls.cfg.SocialMedia))
	for i, sm := range ls.cfg.SocialMedia {
		socialLinks[i] = Link{Title: sm.Platform, URL: sm.URL}
	}
	return socialLinks
}

func (ls *LinkService) GetAffiliateLinks() []Link {
	affiliateLinks := make([]Link, len(ls.cfg.AffiliateLinks))
	for i, al := range ls.cfg.AffiliateLinks {
		affiliateLinks[i] = Link{Title: al.Name, URL: al.URL}
	}
	return affiliateLinks
}
