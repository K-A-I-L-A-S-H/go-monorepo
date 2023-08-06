package providers_test

import (
	"context"
	"github.com/Goldziher/go-monorepo/auth/constants"
	"github.com/Goldziher/go-monorepo/auth/providers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProvider(t *testing.T) {
	for _, testCase := range []struct {
		Provider    string
		ExpectError bool
	}{
		{
			constants.ProviderGithub,
			false,
		},
		{
			constants.ProviderGitlab,
			true,
		},
		{
			constants.ProviderBitBucket,
			true,
		},
		{
			constants.ProviderGoogle,
			true,
		},
		{
			"facebook",
			true,
		},
	} {
		config, err := providers.GetProvider(context.TODO(), testCase.Provider)
		if testCase.ExpectError {
			assert.Nil(t, config)
			assert.NotNil(t, err)
		} else {
			assert.NotNil(t, config)
			assert.Nil(t, err)
		}
	}
}
