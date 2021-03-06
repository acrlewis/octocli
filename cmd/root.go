package cmd

import (
	"fmt"
	"os"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile, token, server string
var markdown bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "octocli",
	Short: "A Command Line Interface for GitHub",
	Long: `octocli is a command line interface which goal is to make life easier
for GitHub Organizations and GitHub Enterprise administrators

Developement takes place at https://github.com/helaili/octocli.
Pull Requests are welcomed!`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.octocli.yaml)")

	rootCmd.PersistentFlags().StringVarP(&token, "token", "k", "", "Personal authentication token to use. Required when environement variable GITHUB_AUTH_TOKEN is not set")
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))

	rootCmd.PersistentFlags().StringVarP(&server, "server", "s", "github.com", "Hostname of the GitHub Enterprise server. Using github.com if omitted")
	viper.BindPFlag("server", rootCmd.PersistentFlags().Lookup("server"))

	rootCmd.PersistentFlags().BoolVarP(&markdown, "markdown", "m", false, "Use markdown for output")
	viper.BindPFlag("markdown", rootCmd.PersistentFlags().Lookup("markdown"))

	if os.Getenv("GITHUB_AUTH_TOKEN") == "" {
		rootCmd.MarkPersistentFlagRequired("token")
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".octocli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".octocli")
	}

	viper.AutomaticEnv() // read in environment variables that match
	viper.BindEnv("token", "GITHUB_AUTH_TOKEN")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
