package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	slice := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		slice[i] = make([]uint8, dx)

	}

	for y := range slice {
		for x := range slice[y] {
			slice[y][x] = uint8(x * y)

		}
	}
	return slice
}

func main() {
	pic.Show(Pic)
}
