package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/helaili/octocli/api"
)

// listCmd represents the list command
var listReposCmd = &cobra.Command{
	Use:   "list",
	Short: "List repos within an organization",
	Run: func(cmd *cobra.Command, args []string) {
		api.PrintRepos(viper.GetString("org"))
	},
}

func init() {
	repoCmd.AddCommand(listReposCmd)
	listReposCmd.Flags().StringP("org", "o", "", "The name of the parent user or organization")
	viper.BindPFlag("org", listReposCmd.Flags().Lookup("org"))
	listReposCmd.MarkFlagRequired("org")
}