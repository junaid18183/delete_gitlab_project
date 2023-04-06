package main

import (
	"fmt"
	"log"
	"os"

	gitlab "github.com/xanzy/go-gitlab"
)

func main() {
	// Get the GitLab access token from the environment variable
	accessToken := os.Getenv("GITLAB_ACCESS_TOKEN")
	if accessToken == "" {
		log.Fatal("GITLAB_ACCESS_TOKEN environment variable not set")
	}

	// Create a new GitLab API client
	git, err := gitlab.NewClient(accessToken)
	if err != nil {
		log.Fatalf("Failed to create GitLab client: %v", err)
	}

	// Get the project name from the command line arguments
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <project name>")
	}
	projectName := "enbuild-staging/" + os.Args[1]

	// Find the project ID for the given project name

	project, _, err := git.Projects.GetProject(projectName, nil)
	if err != nil {
		log.Fatalf("Failed to get project ID for %s: %v", projectName, err)
	}

	// Print the project ID
	fmt.Printf("Project ID for %s: %d\n", projectName, project.ID)

	// Delete the project
	_, err = git.Projects.DeleteProject(project.ID)
	if err != nil {
		log.Fatalf("Failed to delete project: %v", err)
	}

	// Print a success message
	fmt.Printf("Successfully deleted project %s\n", projectName)
}
