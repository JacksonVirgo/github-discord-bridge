package controllers

import (
	"os"

	"github.com/octokit/go-sdk/pkg"
)

func StartGithub(token string) (*pkg.Client, error) {
	client, err := pkg.NewApiClient(
		pkg.WithTokenAuthentication(os.Getenv("GITHUB_TOKEN")),
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}