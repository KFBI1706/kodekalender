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
	var x, y, n int
	dat, err := ioutil.ReadFile("./input-crisscross.txt")
	check(err)
	order := make([]rune, 0)
	nums := strings.FieldsFunc(strings.Split(string(dat), "\n")[0], func(r rune) bool {
		switch r {
		case 'H', 'V', 'F', 'B':
			order = append(order, r)
			return true
		}
		return false
	})

	for i := 0; i < len(nums); i++ {
		n, _ = strconv.Atoi(nums[i])
		switch order[i] {
		case 'H':
			x += n
		case 'V':
			x -= n
		case 'F':
			y += n
		case 'B':
			y -= n
		}
	}
	fmt.Printf("[%d,%d]\n", x, y)
}
