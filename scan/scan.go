package main

import "fmt"

func main() {
	var input string
	for {
		fmt.Print("何か入力してください。'q'で終了:")
		fmt.Scanln(&input)

		if input == "q" {
			break
		}
		fmt.Println(input, "が入力されました")
	}
	fmt.Println("お疲れさまでした")
}
