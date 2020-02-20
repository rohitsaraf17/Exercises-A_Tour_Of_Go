package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var temp [][]uint8
	for i:=0; i<dx; i++ {
		var res []uint8
		for j:=0; j<dy; j++ {
			res = append(res, uint8(i^j))
		}
		temp = append(temp, res)
	}
	return temp
}

func main() {
	pic.Show(Pic)
}
