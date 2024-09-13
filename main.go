package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"flag"

	"github.com/briandowns/spinner"
	"github.com/google/go-github/v64/github"
	"golang.org/x/oauth2"
)

type User struct {
    Login   string
}

type PullRequestItem struct {
	Title   string
	URL     string
	ItemUser    User
}

type RepositoryData struct {
	Name         string
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

    // Preparing flag for user search
    userStr := flag.String("u", "NO_USER", "Only search for a specific github username")
    flag.Parse()
    var isUserSearch = (*userStr != "NO_USER")

	// Unmarshall settings & setup ctx
	var settings Settings
	json.Unmarshal([]byte(settingsJson), &settings)

	ctx := context.Background()
	client := getClient(settings.APIKey, ctx)
	repositoryData := []RepositoryData{}

	// List all pull requests for repos in settings
	var bVal = false
	s := spinner.New(spinner.CharSets[36], 300*time.Millisecond)
	if(isUserSearch) {
        s.Prefix = "Finding open pull requests for user " + *userStr + " "
    } else {
        s.Prefix = "Checking for open pull requests "
    }

	s.Start()
	for _, repo := range settings.Repositories {
		pullrequests, resp, err := client.PullRequests.List(ctx, settings.Organization, repo.Name, nil)
		s.Suffix = " " + repo.Name
		if err != nil {
			fmt.Printf("\nerror: %v\n", err)
			return
		}
		if len(pullrequests) > 0 {
			data := RepositoryData{Name: strings.ToUpper(repo.Name)}
			var tmpUser = ""
			for _, pulls := range pullrequests {
				pr := PullRequestItem{Title: *pulls.Title, URL: *pulls.HTMLURL, ItemUser: User{*pulls.User.Login}}
				data.PullRequests = append(data.PullRequests, pr)
				tmpUser = pr.ItemUser.Login
			}
            if(!isUserSearch) {
                repositoryData = append(repositoryData, data)
                bVal = true
            }
            if(isUserSearch && strings.ToUpper(tmpUser) == strings.ToUpper(*userStr)) {
                repositoryData = append(repositoryData, data)
                bVal = true
            }
		}
		_ = resp
	}
	s.Stop()
	if !bVal {
		fmt.Println("No open pull requests at this time")
	} else {
		fmt.Println("These pull requests might need your attention")
		fmt.Println("")
            for _, repo := range repositoryData {
                fmt.Println("[ " + repo.Name + " ]")
                for _, pr := range repo.PullRequests {
                    if(!isUserSearch) {
                        fmt.Println(pr.Title + " " + pr.URL)
                    } else if (strings.ToUpper(*userStr) == strings.ToUpper(pr.ItemUser.Login)){
                        fmt.Println(pr.Title + " " + pr.URL)
                    } else {
                        fmt.Println("No open pull requests for user ", *userStr)
                    }
                }
                fmt.Println("")
            }
	}
}
