package controllers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)


type GithubContext struct {
	token string
	repo string
	author string
}

var githubContext = &GithubContext{}

func LoadGithubContext() error {
	githubToken := os.Getenv("GITHUB_TOKEN")
	repo := os.Getenv("GITHUB_REPO")
	author := os.Getenv("GITHUB_AUTHOR")

	fmt.Printf("Loaded Github Context: %v, %v, %v\n", githubToken, repo, author)
	if githubToken == "" || repo == "" || author == "" {
		return errors.New("missing environment variables")
	}

	*githubContext = GithubContext{
		token: githubToken,
		repo: repo,
		author: author,
	}

	return nil
}

func GetIssues() (string, error) {
    url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", githubContext.author, githubContext.repo)
	req, err := http.NewRequest("GET", url, nil)
    if err != nil {
		return "", err
    }

	req.Header.Set("Authorization", "token "+githubContext.token)
    req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
		return "", err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
		return "", err
    }

	return string(body), nil
}