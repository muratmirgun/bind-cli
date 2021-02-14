package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"io/ioutil"
	"os"
)

var (
	dosyabilgi *os.FileInfo
)

func kontrol(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	prompt := promptui.Select{
		Label: "Chose",
		Items: []string{"See file", "lorem ipsum"},
	}

	_, result, err := prompt.Run()
	var username string
	prompt_1 := promptui.Prompt{
		Label:   "File Name",
		Default: username,
	}

	result_1, err := prompt_1.Run()

	kontrol(err)

	if result == "See file" {
		See(result_1)
	}
}

func See(fileName string) {
	//dosyabilgi,err := os.Stat(fileName)
	dosyaicerik, err := ioutil.ReadFile(fileName)
	fmt.Println(string(dosyaicerik))
	kontrol(err)

}
