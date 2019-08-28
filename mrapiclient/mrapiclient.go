package mrapiclient

import (
	"context"
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
)

// MrAPIClient contains the MR API OAUTH Client
type MrAPIClient struct {
	*http.Client
}

// New returns a configured MrAPIClient
func New(apiID, apiSecret, apiURL string) (*MrAPIClient, error) {
	config := &clientcredentials.Config{
		ClientID:     apiID,
		ClientSecret: apiSecret,
		TokenURL:     apiURL,
	}

	ctx := context.Background()
	client := config.Client(ctx)
	return &MrAPIClient{client}, nil
}
