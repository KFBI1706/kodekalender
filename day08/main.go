package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type doll struct {
	color  string
	height int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./input-dolls.txt")
	check(err)
	dolls := make([]doll, 0)
	for _, line := range strings.Split(string(dat), "\n") {
		props := strings.Split(line, ",")
		if len(props) != 2 {
			break
		}
		i, _ := strconv.Atoi(props[1])
		dolls = append(dolls, doll{props[0], i})
	}
	sort.Slice(dolls, func(i, j int) bool {
		return dolls[i].height > dolls[j].height
	})
	var maxSize int
	for i, d := range dolls {
		arr := []doll{d}
		for j := i; j < len(dolls); j++ {
			if arr[len(arr)-1].height > dolls[j].height && arr[len(arr)-1].color != dolls[j].color {
				arr = append(arr, dolls[j])
			}
		}
		if len(arr) > maxSize {
			maxSize = len(arr)
		}
	}
	fmt.Println(maxSize)
}
