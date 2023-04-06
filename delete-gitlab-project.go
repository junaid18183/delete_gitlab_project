package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	gitlab "github.com/xanzy/go-gitlab"
)

var (
	projectName string
	namespace   string
	token       string
)

func main() {
	// Create a new root command for our CLI
	var rootCmd = &cobra.Command{
		Use:   "delete-gitlab-project",
		Short: "Deletes a GitLab project",
		Run:   deleteProject,
	}

	// Add a flag for the GitLab access token
	rootCmd.PersistentFlags().StringVar(&token, "token", "", "GitLab access token")

	// Add a required argument for the GitLab project name
	rootCmd.Flags().StringVarP(&projectName, "name", "n", "", "GitLab project name (required)")
	err := rootCmd.MarkFlagRequired("name")
	if err != nil {
		log.Fatalf("Failed to mark 'name' flag as required: %v", err)
	}

	// Add a flag for the GitLab project namespace
	rootCmd.Flags().StringVarP(&namespace, "namespace", "s", "enbuild-staging", "GitLab project namespace")

	// Execute the root command
	err = rootCmd.Execute()
	if err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}
}

func deleteProject(cmd *cobra.Command, args []string) {
	// Get the GitLab access token from the command line flags or environment variable
	if token == "" {
		token = os.Getenv("GITLAB_ACCESS_TOKEN")
		if token == "" {
			log.Fatal("GitLab access token not provided")
		}
	}

	// Create a new GitLab API client
	git, err := gitlab.NewClient(token)
	if err != nil {
		log.Fatalf("Failed to create GitLab client: %v", err)
	}

	// Find the project ID for the given project name and namespace
	project, _, err := git.Projects.GetProject(namespace+"/"+projectName, nil)
	if err != nil {
		log.Fatalf("Failed to get project ID for %s/%s: %v", namespace, projectName, err)
	}

	// Print the project ID
	fmt.Printf("Project ID for %s/%s: %d\n", namespace, projectName, project.ID)

	// Prompt the user for confirmation before deleting the project
	var confirm string
	fmt.Printf("Are you sure you want to delete project '%s/%s'? (yes/no): ", namespace, projectName)
	_, err = fmt.Scanln(&confirm)
	if err != nil {
		log.Fatalf("Failed to read user input: %v", err)
	}
	if confirm != "yes" {
		fmt.Println("Project deletion aborted.")
		return
	}

	// Delete the project
	_, err = git.Projects.DeleteProject(project.ID)
	if err != nil {
		log.Fatalf("Failed to delete project: %v", err)
	}

	// Print a success message
	fmt.Printf("Successfully deleted project %s/%s\n", namespace, projectName)
}
