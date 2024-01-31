package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)

	for i := range img {
		img[i] = make([]uint8, dx)
		for y := range img {
			img[i][y] = uint8(i) ^ uint8(y)
		}

	}
	return img
}

func main() {
	pic.Show(Pic)
}
