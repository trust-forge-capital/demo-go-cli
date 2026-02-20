package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(healthCmd)
}

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Show basic health information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("status=ok app=demo-go-cli version=%s time=%s\n", Version, time.Now().Format(time.RFC3339))
	},
}
