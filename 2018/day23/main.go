package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

//Int helper function to convert byte value for a number to an int
func main() {
	Int := func(r byte) int { return int(r - '0') }
	var i, n, cnt, firstSum, secondSum, dayNumber, genderNumber, firstCheck, secondCheck int
	firstMultipliers := []int{3, 7, 6, 1, 8, 9, 4, 5, 2, 0}
	secondMultipliers := []int{5, 4, 3, 2, 7, 6, 5, 4, 3, 2}
	dat, _ := ioutil.ReadFile("./input-fnr.txt")
	for _, line := range strings.Split(string(dat), "\n") {
		if len(line) == 0 {
			break
		}
		dayNumber, _ = strconv.Atoi(line[0:2])
		genderNumber, _ = strconv.Atoi(line[8:9])
		if line[2:4] != "08" || genderNumber%2 != 0 || dayNumber > 31 {
			continue
		}
		for i, firstSum, secondSum = 0, 0, 0; i < 10; i++ {
			n = Int(line[i])
			firstSum += firstMultipliers[i] * n
			secondSum += secondMultipliers[i] * n
		}
		firstCheck, secondCheck = Int(line[9]), Int(line[10])
		firstSum, secondSum = (11-(firstSum%11))%11, (11-(secondSum%11))%11
		if firstCheck == firstSum && secondCheck == secondSum && firstSum != 10 && secondSum != 10 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
