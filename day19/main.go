package main

import (
	"fmt"
	"sort"
)

func main() {
	s1, s2 := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. God jul.", "The quick brown fox jumps over the lazy dog. Godt nytt ar."
	fmt.Println(ld(s1, s2))
}

func ld(s1, s2 string) int {
	d := make([]int, 3)
	if len(s2) > len(s1) {
		s1, s2 = s2, s1
	}
	distances := make([]int, len(s1))
	for i := range s1 {
		distances[i] = i
	}
	for ix2, c2 := range s2 {
		newDistances := []int{ix2 + 1}
		for ix1, c1 := range s1 {
			if c1 == c2 {
				newDistances = append(newDistances, distances[ix1])
			} else {
				d = []int{distances[ix1], distances[ix1] + 1, newDistances[len(newDistances)-1]}
				sort.Slice(d, func(i, j int) bool { return d[i] > d[j] })
				newDistances = append(newDistances, 1+d[len(d)-1])
			}
		}
		distances = newDistances
	}
	return distances[len(distances)-1]
}
