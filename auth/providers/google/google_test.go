package google_test

import (
	"context"
	"github.com/Goldziher/go-monorepo/auth/providers/google"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConfig(t *testing.T) {
	t.Setenv("BASE_URL", "http://localhost:3000")
	t.Setenv("ENVIRONMENT", "development")
	t.Setenv("GOOGLE_CLIENT_ID", "googleClientId")
	t.Setenv("GOOGLE_CLIENT_SECRET", "googleClientSecret")
	t.Setenv("PORT", "3000")

	config, err := google.GetConfig(context.TODO())

	assert.Nil(t, err)
	assert.Equal(t, config.ClientID, "googleClientId")
	assert.Equal(t, config.ClientSecret, "googleClientSecret")
	assert.Equal(t, config.RedirectURL, "http://localhost:3000/oauth/google/callback")
}