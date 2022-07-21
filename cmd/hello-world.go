package cmd

import (
	"github.com/spf13/cobra"
)

var helloWorld = &cobra.Command{
	Use:   "say-hello",
	Short: "Displays Hello World",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Println(getWelcomeMessage())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(helloWorld)
}

func getWelcomeMessage() string {
	return "Hello world!"
}
