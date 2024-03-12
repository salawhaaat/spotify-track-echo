package client

import (
	"context"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

var Client *spotify.Client

func New(cfg *clientcredentials.Config) {
	accessToken, err := cfg.Token(context.Background())
	if err != nil {
		return
	}
	c := spotify.Authenticator{}.NewClient(accessToken)
	Client = &c
}
