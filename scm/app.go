package scm

import (
	"context"
	"time"
)

type (
	InstallationToken struct {
		Token     string
		ExpiresAt *time.Time
	}

	// Installation represents a GitHub app install
	Installation struct {
		ID                  int64
		AppID               int64
		TargetID            int64
		TargetType          string
		RepositorySelection string
		Account             Account
		AccessTokensLink    string
		RepositoriesURL     string
		Link                string
		Events              []string
		CreatedAt           *time.Time
		UpdatedAt           *time.Time
	}

	// AppService for GitHub App support
	AppService interface {
		CreateInstallationToken(ctx context.Context, id int64) (*InstallationToken, *Response, error)

		GetRepositoryInstallation(ctx context.Context, fullName string) (*Installation, *Response, error)
	}
)
