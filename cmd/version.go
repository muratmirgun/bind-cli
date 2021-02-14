package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Bixi Cli Version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bixi Cli Version v1.0")
	},
}
