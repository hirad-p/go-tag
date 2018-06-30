package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GithubConfig is a github configuration struct
type GithubConfig struct {
	Org, Repo, Token string
}

// JiraConfig is a Jira configuration struct
type JiraConfig struct {
	Username, Password, Host string
}

// ReadGithub returns the environment variable for Github
func ReadGithub() GithubConfig {
	load()
	return GithubConfig{
		Org:   os.Getenv("GIT_ORG"),
		Repo:  os.Getenv("GIT_REPO"),
		Token: os.Getenv("GIT_ACCESS_TOKEN"),
	}
}

// ReadJira returns the environment variable for Jira
func ReadJira() JiraConfig {
	load()
	return JiraConfig{
		Username: os.Getenv("JIRA_USER"),
		Password: os.Getenv("JIRA_PASS"),
		Host:     os.Getenv("JIRA_HOST"),
	}
}

func load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
