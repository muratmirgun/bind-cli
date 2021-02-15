package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func init() {
	rootCmd.AddCommand(info)
}

var info = &cobra.Command{
	Use:   "info",
	Short: "file info in term",
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := strings.Join(args, " ")
		fi, _ := os.Stat(name)
		fmt.Println("File Name: " + fi.Name())
		fmt.Println("File Size: ", fi.Size())
		fmt.Println("File Mod Time: ", fi.ModTime())
	},
}
