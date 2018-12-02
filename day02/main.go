package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Line struct {
	x1, x2, y1, y2 int
}

func (l Line) slope() float32 {
	return float32(l.y2-l.y1) / float32(l.x2-l.x1)
}

func main() {
	lines := make([]Line, 0)
	cnt := make(map[float32]int)
	file, err := os.Open("./input-rain.txt")
	check(err)
	for {
		l := Line{}
		fmt.Fscanf(file, "(%d,%d);(%d,%d)\n", &l.x1, &l.y1, &l.x2, &l.y2)
		if l == (Line{}) {
			break
		}
		lines = append(lines, l)
	}
	for _, line := range lines {
		if i, ok := cnt[line.slope()]; ok {
			cnt[line.slope()] = i + 1
		} else {
			cnt[line.slope()] = 1
		}
	}
	var big int
	for _, i := range cnt {
		if i > big {
			big = i
		}
	}
	fmt.Println(big)
}
