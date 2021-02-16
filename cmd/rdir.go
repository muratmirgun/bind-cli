package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func init() {
	rootCmd.AddCommand(Rdir)
}

var Rdir = &cobra.Command{
	Use:   "rdir",
	Short: "remove directory in term",
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := strings.Join(args, " ")
		os.Remove(name)
	},
}
