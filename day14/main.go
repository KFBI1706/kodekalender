package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	var x, y, n, x1, y1, x2, y2, j int16
	var LEN, i uint32
	var b byte
	mp := make(map[string]bool)
	LEN = 2000
	b = 48
	dat, err := ioutil.ReadFile("./input-bounding-crisscross.txt")
	check(err)
	mp[fmt.Sprintf("%d,%d", x, y)] = true
	for i = 1; i < LEN; i += 2 {
		n = int16(dat[i-1] & ^b) //Remove 48
		for j = 0; j < n; j++ {
			switch dat[i] {
			case 72: //H
				if x++; x > x2 {
					x2 = x
				}
			case 86: //V
				if x--; x < x1 {
					x1 = x
				}
			case 70: //F
				if y++; y > y2 {
					y2 = y
				}
			case 66: //B
				if y--; y < y1 {
					y1 = y
				}
			}
			mp[fmt.Sprintf("%d,%d", x, y)] = true
		}
	}
	box := float64(((y2 - y1) + 1) * ((x2 - x1) + 1))
	visitedCnt := float64(len(mp))
	fmt.Printf("%.16f\n", visitedCnt/(box-visitedCnt))
}
