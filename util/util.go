package util

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/google/go-github/github"
)

// PrintStatusCode prints a prettified response depeneding on
// the http status code
func PrintStatusCode(ticket string, code int) {
	switch code {
	case 204:
		fmt.Printf("Ticket %s was not updated.\n", ticket)
	default:
		fmt.Printf("Unable to identify code %d for ticket %s\n", code, ticket)
	}
}

// FilterIssues return new slice of ticket numbers that are pull requests and have jira
// identifiable tickets
func FilterIssues(issues []*github.Issue) []string {
	result := make([]string, len(issues))
	counter := 0
	for _, issue := range issues {
		if issue.PullRequestLinks != nil {
			if ticket := extractTicketNumber(*issue.Title); ticket != "" {
				result[counter] = ticket
				counter++
			}
		}
	}
	return result[:counter]
}

// ExtractBuildNumber extracts the build number from a release tag
// expecting <release-x.x.x.x> returning <x.x.x.x>
func ExtractBuildNumber(name string) string {
	return strings.Split(name, "-")[1]
}

func extractTicketNumber(title string) string {
	r, err := regexp.Compile(`(?i)((CS|JIT|CN)-\d+)`)
	if err != nil {
		log.Fatal(err)
	}
	if matches := r.FindStringSubmatch(title); len(matches) > 0 {
		return matches[0]
	}
	return ""
}
