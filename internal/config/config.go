// Package config provides configuration structures and loading functionality
package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Config holds the application configuration
type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Websites       []Website       `yaml:"websites"`
	SocialMedia    []SocialMedia   `yaml:"social_media"`
	AffiliateLinks []AffiliateLink `yaml:"affiliate_links"`
}

// Website represents a website entry in the configuration
type Website struct {
	Name        string     `yaml:"name"`
	URL         string     `yaml:"url"`
	LastUpdate  CustomTime `yaml:"last_update"`
	Description string     `yaml:"description"`
}

// SocialMedia represents a social media entry in the configuration
type SocialMedia struct {
	Platform    string     `yaml:"platform"`
	URL         string     `yaml:"url"`
	LastUpdate  CustomTime `yaml:"last_update"`
	Description string     `yaml:"description"`
}

// AffiliateLink represents an affiliate link entry in the configuration
type AffiliateLink struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

// CustomTime is a wrapper for time.Time with custom YAML unmarshaling
type CustomTime struct {
	time.Time
}

// UnmarshalYAML implements custom YAML unmarshaling for CustomTime
func (ct *CustomTime) UnmarshalYAML(value *yaml.Node) error {
	var timeStr string
	if err := value.Decode(&timeStr); err != nil {
		return err
	}

	parsedTime, err := time.Parse("2006-01-02", timeStr)
	if err != nil {
		return err
	}

	ct.Time = parsedTime
	return nil
}

// Load reads the configuration file and returns a Config struct
func Load() (*Config, error) {
	var cfg Config
	file, err := os.ReadFile("configs/config.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
