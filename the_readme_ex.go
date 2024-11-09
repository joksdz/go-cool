package main

import (
	"math"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy, dy)
	i := 10
	for y := range pic {
		pic[y] = make([]uint8, dx)

		for x := range pic[y] {

			i++
			pic[y][x] = uint8(float64((x+y)*i) / math.Pi)
		}
	}

	return pic
}

func main() {
	pic.Show(Pic)
}
