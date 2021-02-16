package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func init() {
	rootCmd.AddCommand(Cfile)
}

var Cfile = &cobra.Command{
	Use:   "cfile",
	Short: "create file in term",
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := strings.Join(args, " ")
		os.Create(name)
	},
}
