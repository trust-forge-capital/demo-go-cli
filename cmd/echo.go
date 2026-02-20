package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var upper bool

func init() {
	echoCmd.Flags().BoolVar(&upper, "upper", false, "Output in uppercase")
	rootCmd.AddCommand(echoCmd)
}

var echoCmd = &cobra.Command{
	Use:   "echo [text]",
	Short: "Echo input text",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		out := strings.Join(args, " ")
		if upper {
			out = strings.ToUpper(out)
		}
		fmt.Println(out)
	},
}
