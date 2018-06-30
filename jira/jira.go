package jira

import (
	"log"

	jira "github.com/andygrunwald/go-jira"

	"github.com/jb-hirad/go-tag/config"
	"github.com/jb-hirad/go-tag/util"
)

// UpdateIssues updates a list of tickets' fixed in version to the version supplied
func UpdateIssues(tickets []string, verison string) {
	jiraConfig := config.ReadJira()

	tp := jira.BasicAuthTransport{
		Username: jiraConfig.Username,
		Password: jiraConfig.Password,
	}

	client, err := jira.NewClient(tp.Client(), jiraConfig.Host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ticket := range tickets {
		data := map[string]interface{}{
			"fields": map[string][]string{
				"customfield_10700": []string{verison},
			},
		}
		res, err := client.Issue.UpdateIssue(ticket, data)
		if err != nil {
			log.Fatal(err)
		}
		util.PrintStatusCode(ticket, res.StatusCode)
	}
}
