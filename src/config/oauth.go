package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	Oauth2Config *oauth2.Config
	authUrl      string
)

func Oauth() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	Oauth2Config = &oauth2.Config{
		ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("OAUTH_REDIRECT_URI"),
		Scopes:       []string{"profile", "email"},
		Endpoint:     google.Endpoint,
	}

	authUrl = Oauth2Config.AuthCodeURL("state", oauth2.AccessTypeOffline)
}

func GetAuthUrl() string {
	return authUrl
}

func Exchange(code string) (*oauth2.Token, error) {
	ctx := context.Background()
	token, err := Oauth2Config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	return token, nil
}
