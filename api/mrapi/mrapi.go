package mrapi

import (
	"context"
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
)

// Client contains the MR API Client
type Client struct {
	*http.Client
}

// New returns a configured API Client for the
func New(ID, secret, URL string) (*Client, error) {
	config := &clientcredentials.Config{
		ClientID:     ID,
		ClientSecret: secret,
		TokenURL:     URL + "/oauth/v2/token",
	}

	ctx := context.Background()
	client := config.Client(ctx)
	return &Client{client}, nil
}
