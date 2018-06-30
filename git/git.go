package git

import (
	"context"
	"log"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"

	"github.com/jb-hirad/go-tag/config"
	"github.com/jb-hirad/go-tag/util"
)

// VersionInfo holds information about the current releast and the previous one
type VersionInfo struct {
	Past, Current ReleaseInfo
	Tickets       []string
}

// ReleaseInfo holds the information of last relese
type ReleaseInfo struct {
	Name string
	Time *time.Time
}

// GetVersionInfo returns the current version info. Version info consist of current release and past release
func GetVersionInfo() VersionInfo {
	githubConfig := config.ReadGithub()

	// authorization
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubConfig.Token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// get a list of tags which resemble builds
	tags, _, err := client.Repositories.ListTags(ctx, githubConfig.Org, githubConfig.Repo, nil)
	if err != nil {
		log.Fatal(err)
	}

	// get most recent build and extract the release info
	current := tags[0]
	currentCommit, _, err := client.Repositories.GetCommit(ctx, githubConfig.Org, githubConfig.Repo, *current.Commit.SHA)
	if err != nil {
		log.Fatal(err)
	}
	currentRelease := ReleaseInfo{Name: util.ExtractBuildNumber(*current.Name), Time: currentCommit.Commit.Author.Date}

	// get previous build and extract the release info
	past := tags[1]
	pastCommit, _, err := client.Repositories.GetCommit(ctx, githubConfig.Org, githubConfig.Repo, *past.Commit.SHA)
	if err != nil {
		log.Fatal(err)
	}
	pastRelease := ReleaseInfo{Name: util.ExtractBuildNumber(*past.Name), Time: pastCommit.Commit.Author.Date}

	// Get a list of issues
	opt := github.IssueListByRepoOptions{Since: *pastRelease.Time}
	issues, _, err := client.Issues.ListByRepo(ctx, githubConfig.Org, githubConfig.Repo, &opt)
	if err != nil {
		log.Fatal(err)
	}

	return VersionInfo{Past: pastRelease, Current: currentRelease, Tickets: util.FilterIssues(issues)}
}
