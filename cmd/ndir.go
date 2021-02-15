package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func init() {
	rootCmd.AddCommand(Mdir)
}

var Mdir = &cobra.Command{
	Use:   "mdir",
	Short: "create directory in term",
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := strings.Join(args, " ")
		os.Mkdir(name, 0700)
	},
}
