package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var ret [][]uint8
	for y := 0; y < dy; y++ {
		var tmp []uint8
		for x := 0; x < dy; x++ {
			tmp = append(tmp, uint8((x+y)/2))
		}
		ret = append(ret, tmp)
	}
	return ret
}

func Pic2(dx, dy int) [][]uint8 {
	var ret [][]uint8 = make([][]uint8, dy)
	for i := range ret {
		ret[i] = make([]uint8, dx)
		for j := range ret[i] {
			ret[i][j] = uint8((i + j) / 2)
		}
	}
	return ret
}

func main() {
	pic.Show(Pic2)
}
