package cmd

import (
	"github.com/spf13/cobra"
	"github.com/helaili/octocli/api"
)

var listReposCmdLogin string

// listCmd represents the list command
var listReposCmd = &cobra.Command{
	Use:   "list",
	Short: "List repos within an organization",
	Run: func(cmd *cobra.Command, args []string) {
		api.PrintRepos(listReposCmdLogin)
	},
}

func init() {
	repoCmd.AddCommand(listReposCmd)
	listReposCmd.Flags().StringVarP(&listReposCmdLogin, "login", "l", "", "The name of the parent user or organization")
	listReposCmd.MarkFlagRequired("login")
}
