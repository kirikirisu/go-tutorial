package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"io/ioutil"
	"os"
)

func main() {
	reader, ferr := os.Open("grayscale.png")
	if ferr != nil {
		fmt.Println("ファイル読み込みのエラーです。", ferr)
		os.Exit(1)
	}

	img, name, derr := image.Decode(reader)
	if derr != nil {
		fmt.Println("画像変換のエラーです。", derr)
		os.Exit(1)
	} else {
		fmt.Println(name, "形式のデータを得ました。")
	}

	defer reader.Close()

	marks := []string{"*", "+", "-"}
	var marksStr string

	// 画像ファイルの上から下、右から左までfor文を回す
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			grayness := c.Y / (255 / 3)
			if c.Y == 255 {
				grayness = 2
			}
			marksStr += marks[grayness]
		}
		marksStr += "\n"
	}

	wdata := []byte(marksStr)
	werr := ioutil.WriteFile("ascii_art.txt", wdata, 0777)
	if werr != nil {
		fmt.Println("ファイルの書き込みエラーです。", werr)
		os.Exit(1)
	} else {
		fmt.Println("ファイルを保存しました。")
	}
}
