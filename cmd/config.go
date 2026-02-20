package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	configCmd.Flags().StringVar(&configPath, "path", ".demo-go-cli/config.yaml", "Config file path")
	rootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show current config path/status",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(configPath); err == nil {
			fmt.Printf("config exists: %s\n", configPath)
			return
		}
		fmt.Printf("config missing: %s (run: demo-go-cli init)\n", configPath)
	},
}
