package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type pair struct {
	original, square int
}

func main() {
	var i float64
	dat, err := ioutil.ReadFile("./input-gullbursdag.txt")
	check(err)
	squares := make([]pair, 100)
	for i = 1; i <= 100; i++ {
		squares[int(i)-1] = pair{int(i), int(math.Pow(i, 2))}
	}

	var cnt int
	for _, line := range strings.Split(string(dat), "\n") {
		parts := strings.Split(line, ".")
		if len(parts) != 2 {
			break
		}
		i, _ := strconv.Atoi(parts[1])
		for age := 0; age < 100; age++ {
			for _, sq := range squares {
				if sq.square == i+age && age == sq.original {
					fmt.Printf("%s born in %d at the age of %d matches this year: %d\n", parts[0], i, age, i+age)
					cnt++

				}
			}
		}
	}
	fmt.Println(cnt)
}
