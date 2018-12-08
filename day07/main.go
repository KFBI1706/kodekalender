package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var arr []int

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var curr, lo, hi, N, L int
	var M, P []int
	nums := make([]int, 0)
	dat, err := ioutil.ReadFile("./input-vekksort.txt")
	check(err)
	for _, line := range strings.Split(string(dat), "\n") {
		if curr, err = strconv.Atoi(line); err != nil {
			continue
		}
		nums = append(nums, curr)
	}
	N = len(nums)
	P = make([]int, N)
	M = make([]int, N+1)
	L = 0
	for i := 0; i < N; i++ {
		lo = 1
		hi = L
		for lo <= hi {
			mid := (lo + hi) / 2
			if nums[M[mid]] <= nums[i] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
		newL := lo
		P[i] = M[newL-1]
		M[newL] = i

		if newL > L {
			L = newL
		}
	}

	fmt.Println(L)
}

//smallest := make([]int, len(nums))
//for i := 0; i < len(nums); i++ {
//	if i%100000 == 0 {
//		fmt.Println(i)
//	}
//	cnt = 0
//	//num position and how many lower numbers there are after it
//	cause := make([]int, 0)
//	for j := i; j < len(nums); j++ {
//		if nums[i] > nums[j] {
//			cause = append(cause, nums[j])
//			cnt++
//		}
//		if cnt > 1 {
//			fmt.Printf("Eliminating %d, at pos %d because of %v\n", nums[i], i, cause)
//			nums = append(nums[:i], nums[i+1:]...)
//			break
//		}
//	}
//}
