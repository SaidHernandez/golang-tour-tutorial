package images

import (
	"fmt"
	"image"
)

func CallImages() {
	fmt.Println("Images......")
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
