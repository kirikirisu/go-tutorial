package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var input string
	var content string
	for {
		fmt.Print("何か入力してください。's'で保存: ")
		fmt.Scanln(&input)

		if input == "s" {
			break
		}
		content += input
	}

	wdata := []byte(content)
	err := ioutil.WriteFile("datafile.txt", wdata, 0777)
	if err != nil {
		fmt.Println("書き込みエラーです。", err)
		os.Exit(1)
	} else {
		fmt.Println(content)
		fmt.Println("入力された情報をファイルに保存しました。")
	}
}
