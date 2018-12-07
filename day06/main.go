package main

import (
	"fmt"
	"strconv"
)

var n, zcnt, ocnt int
var sum uint64
var s string

func main() {
	n = 18163106
	for i := 0; i <= n; i++ {
		s = strconv.Itoa(i)
		zcnt, ocnt = 0, 0
		for _, c := range s {
			if c == '0' {
				zcnt++
			} else {
				ocnt++
			}
		}
		if zcnt > ocnt {
			sum += uint64(i)
		}
	}
	fmt.Println(sum)
}
