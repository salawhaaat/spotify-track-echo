package config

import (
	"flag"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

type Config struct {
	ClientID     string
	ClientSecret string
}

func New() *clientcredentials.Config {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()
	cfg := &Config{}
	if err := cleanenv.ReadConfig(*configPath, cfg); err != nil {
		log.Fatalf("cannot read config file: %v", err)
	}
	config := &clientcredentials.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		TokenURL:     spotify.TokenURL,
	}

	if config.TokenURL == "" {
		log.Fatal("TokenURL is empty in client credentials config")
	}

	return config
}
