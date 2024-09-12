package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
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


type Issue struct {
	Title string `json:"title"`
}

func GetIssues() ([]string, error) {
    url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", githubContext.author, githubContext.repo)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Printf("Error creating request: %s", err)
        return nil, err
    }

    req.Header.Set("Authorization", "token "+githubContext.token)
    req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Printf("Error sending request: %s", err)
        return nil, err
    }

    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Printf("Error reading response body: %s", err)
        return nil, err
    }

    var issues []Issue
    err = json.Unmarshal(body, &issues)
    if err != nil {
        log.Printf("Error parsing JSON: %s", err)
        return nil, err
    }

    var titles []string
    for _, issue := range issues {
        titles = append(titles, issue.Title)
    }

    return titles, nil
}