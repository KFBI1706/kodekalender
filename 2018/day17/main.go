package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

//heap's algorithm permutations
func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

type cord struct {
	x, y float64
}

func main() {
	var x, y, distance, minDistance float64
	var c []string
	dat, _ := ioutil.ReadFile("./input-reiserute.txt")
	cords := make([]cord, 0)
	arr := make([]int, 0)
	for i, line := range strings.Split(string(dat), "\n") {
		if c = strings.Split(line, ","); len(c) != 2 {
			break
		}
		x, _ = strconv.ParseFloat(c[0], 64)
		y, _ = strconv.ParseFloat(c[1], 64)
		cords = append(cords, cord{x, y})
		arr = append(arr, i)
	}

	minDistance = math.MaxFloat64
	for _, perm := range permutations(arr) {
		distance = 0.0
		for i := 0; i < len(perm)-1; i++ {
			c1 := cords[perm[i]]
			c2 := cords[perm[i+1]]
			distance += math.Sqrt(math.Pow(c1.x-c2.x, 2) + math.Pow(c1.y-c2.y, 2))
		}
		if distance < minDistance {
			minDistance = distance
		}
	}
	fmt.Println(minDistance)
}
