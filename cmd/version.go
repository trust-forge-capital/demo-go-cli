package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "v0.1.0"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("demo-go-cli version %s\n", version)
	},
}
