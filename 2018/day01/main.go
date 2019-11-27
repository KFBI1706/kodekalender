package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./input-vekksort.txt")
	check(err)
	var sum, big, curr int
	for _, line := range strings.Split(string(dat), "\n") {
		if curr, err = strconv.Atoi(line); err != nil {
			continue
		}
		if curr >= big {
			sum += curr
			big = curr
		}
	}
	fmt.Println(sum)
}
