package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "v0.2.0"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("demo-go-cli version %s\n", Version)
	},
}
