package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jb-hirad/go-tag/git"
	"github.com/jb-hirad/go-tag/jira"
)

func main() {
	info := git.GetVersionInfo()
	build := info.Current.Name
	tickets := info.Tickets

	fmt.Println("Tickets included in this build: ")
	for _, ticket := range tickets {
		fmt.Println(ticket)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Update Fix in Version of above tickets to %s? [Y/N]", build)
	confirm, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Unable to confirm")
	} else if strings.ToUpper(confirm) != "Y" {
		fmt.Println("Aborting")
		os.Exit(1)
	}

	jira.UpdateIssues(tickets, build)
}
