package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	configPath string
	force      bool
)

func init() {
	initCmd.Flags().StringVar(&configPath, "path", ".demo-go-cli/config.yaml", "Config file path")
	initCmd.Flags().BoolVar(&force, "force", false, "Overwrite existing config")
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize demo config file",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := os.MkdirAll(filepath.Dir(configPath), 0o755); err != nil {
			return err
		}

		if _, err := os.Stat(configPath); err == nil && !force {
			return fmt.Errorf("config already exists: %s (use --force to overwrite)", configPath)
		}

		content := "app: demo-go-cli\nversion: " + Version + "\n"
		if err := os.WriteFile(configPath, []byte(content), 0o644); err != nil {
			return err
		}

		fmt.Printf("initialized config: %s\n", configPath)
		return nil
	},
}
