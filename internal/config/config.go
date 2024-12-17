package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Websites []struct {
		Name string `yaml:"name"`
		URL  string `yaml:"url"`
	} `yaml:"websites"`
	SocialMedia []struct {
		Platform string `yaml:"platform"`
		URL      string `yaml:"url"`
	} `yaml:"social_media"`
	AffiliateLinks []struct {
		Name string `yaml:"name"`
		URL  string `yaml:"url"`
	} `yaml:"affiliate_links"`
	Design struct {
		PrimaryColor   string `yaml:"primary_color"`
		SecondaryColor string `yaml:"secondary_color"`
		TertiaryColor  string `yaml:"tertiary_color"`
		FontFamily     string `yaml:"font_family"`
	} `yaml:"design"`
	Features struct {
		EnableDarkMode  bool `yaml:"enable_dark_mode"`
		EnableAnalytics bool `yaml:"enable_analytics"`
	} `yaml:"features"`
}

func Load() (*Config, error) {
	config := &Config{}

	file, err := os.ReadFile("configs/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return config, nil
}
