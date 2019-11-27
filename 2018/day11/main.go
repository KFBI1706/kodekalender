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
	var x, y, n int16
	var LEN, i uint32
	var b byte
	LEN = 20000000
	b = 48
	dat, err := ioutil.ReadFile("/tmp/input-crisscross.txt")
	check(err)
	for i = 1; i < LEN; i += 2 {
		n = int16(dat[i-1] & ^b)
		switch dat[i] {
		case 72: //H 64, 8
			x += n
		case 86: //V 64, 16, 4, 2
			x -= n
		case 70: //F 64, 4, 2
			y += n
		case 66: //B 64, 2
			y -= n
		}
	}
	fmt.Printf("[%d,%d]\n", x, y)
}
