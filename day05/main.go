package main

import (
	"fmt"
	"strconv"
	"strings"
)

var MAX int
var sarr, operators []string

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func eval(s string) (ans int) {
	var order []rune
	snums := strings.FieldsFunc(s, func(r rune) bool {
		switch r {
		case '+', '-':
			order = append(order, r)
			return true
		}
		return false
	})
	var nums []int
	for _, s := range snums {
		n, _ := strconv.Atoi(s)
		nums = append(nums, n)
	}
	ans = nums[0]
	for i := 1; i < len(nums); i++ {
		switch order[i-1] {
		case '+':
			ans += nums[i]
		case '-':
			ans -= nums[i]
		}
	}
	return ans
}

func TheAnswerToTheUniverseAndEverything(s string, operations int) (cnt int) {
	if operations == MAX {
		if eval(s) == 42 {
			return 1
		} else {
			return 0
		}
	}
	for _, op := range operators {
		cnt += TheAnswerToTheUniverseAndEverything(s+op+sarr[operations], operations+1)
	}
	return
}

func main() {
	operators = []string{"+", "-", ""}
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 7, 6, 5, 4, 3, 2, 1}
	MAX = len(nums)
	sarr = make([]string, 0)
	for _, n := range nums {
		sarr = append(sarr, strconv.Itoa(n))
	}
	fmt.Println(TheAnswerToTheUniverseAndEverything(sarr[0], 1))
}
