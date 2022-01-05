package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
)

type PullRequestItem struct {
	Repository string
	Title      string
	URL        string
}

type PullRequestsData struct {
	PullRequests []PullRequestItem
}

func getClient(token string, ctx context.Context) *github.Client {
	// Create and init a github client from background ctx with a static token
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}

func main() {
	// Unmarshall settings & setup ctx
	var settings Settings
	json.Unmarshal([]byte(settingsJson), &settings)

	ctx := context.Background()
	client := getClient(settings.APIKey, ctx)
	pullRequestData := PullRequestsData{}

	// List all pull requests for repos in settings
	var bVal = false
	fmt.Println("Checking repositories for open pull requests")
	s := spinner.New(spinner.CharSets[36], 300*time.Millisecond)
	s.Start()
	for _, repo := range settings.Repositories {
		pullrequests, resp, err := client.PullRequests.List(ctx, settings.Organization, repo.Name, nil)
		if err != nil {
			fmt.Printf("\nerror: %v\n", err)
			return
		}
		if len(pullrequests) > 0 {
			for _, pulls := range pullrequests {
				pr := PullRequestItem{Repository: strings.ToUpper(repo.Name), Title: *pulls.Title, URL: *pulls.HTMLURL}
				pullRequestData.PullRequests = append(pullRequestData.PullRequests, pr)
				fmt.Println(strings.ToUpper(repo.Name) + " " + *pulls.Title + " " + *pulls.HTMLURL)
			}
		} else {
			bVal = true
		}
		_ = resp
	}
	s.Stop()
	if bVal {
		fmt.Println("No open pull requests at this time")
	}
}
