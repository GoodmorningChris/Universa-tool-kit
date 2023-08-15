package server

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

func GetRandomAvtar(name string) {

	width, height := 64, 64
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	var red = [...]int{142, 255, 190, 130}
	var green = [...]int{207, 190, 184, 176}
	var blue = [...]int{201, 122, 220, 210}
	for x := 0; x < 64; x++ {
		for y := 0; y < 64; y++ {
			if x%8 == 0 && y%8 == 0 {
				t := rand.Uint32() % 4
				img.Set(x, y, color.RGBA{uint8(red[t]), uint8(green[t]), uint8(blue[t]), 0xff})
			} else {
				c := img.RGBAAt(x-x%8, y-y%8)
				img.Set(x, y, c)
			}
		}

	}

	f, err1 := os.Create("./" + name)
	if err1 != nil {
		fmt.Println(err1)
	}
	png.Encode(f, img)
	srcFile, _ := os.Open(name)

	UploadAliyunOss2(srcFile, name)
}
