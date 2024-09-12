package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateIssueComment(issueNumber int, comment string) error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%d/comments", GithubContext.author, GithubContext.repo, issueNumber)
	reqBody, err := json.Marshal(map[string]string{
		"body": comment,
	})
	if err != nil {
		log.Printf("Error marshalling request body: %s", err)
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Printf("Error creating request: %s", err)
		return err
	}

	req.Header.Set("Authorization", "token "+GithubContext.token)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %s", err)
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err)
		return err
	}

	var response CreateIssueResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("Error parsing JSON: %s", err)
		return err
	}

	return err
}
