package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strings"
)

func init() {
	rootCmd.AddCommand(Open)
}

var Open = &cobra.Command{
	Use:   "open ",
	Short: "open your file in term",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := strings.Join(args, " ")
		dosyaicerik, _ := ioutil.ReadFile(string(name))
		fmt.Println(string(dosyaicerik))
	},
}

func init() {

}
