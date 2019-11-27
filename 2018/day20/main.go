package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type ds struct {
	i     int
	s, rs string
}

//DigitString custom type declaring a list of pointer to the underlying struct.
type DigitString []*ds

//Generic functions so sort.Sort() will work
func (d DigitString) Len() int           { return len(d) }
func (d DigitString) Less(i, j int) bool { return d[j].rs < d[i].rs }
func (d DigitString) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

func li(nums []int) {
	//Slice of pointers
	dds := make(DigitString, len(nums))
	//Original digitString, so that we create all the structs in one go, and can use the pointers in the next for-loop
	ds := make([]ds, len(dds))
	mx := 0
	//Populate ds with position and string, while finding the longest string
	for j, i := range nums {
		p := &ds[j]
		dds[j] = p
		p.i = i
		p.s = strconv.Itoa(i)
		if len(p.s) > mx {
			mx = len(p.s)
		}
	}
	//Create repeating string based on max length
	for _, p := range dds {
		p.rs = strings.Repeat(p.s, (mx+len(p.s)-1)/len(p.s))
	}
	sort.Sort(dds)
	//Output slice
	for _, p := range dds {
		fmt.Printf("%s", p.s)
	}
}

func main() {
	nums := []int{7096, 3, 3924, 2404, 4502, 4800, 74, 91, 9, 7, 9, 6790, 5, 59, 9, 48, 6345,
		88, 73, 88, 956, 94, 665, 7, 797, 3978, 1, 3922, 511, 344, 6, 10, 743, 36,
		9289, 7117, 1446, 10, 7466, 9, 223, 2, 6, 528, 37, 33, 1616, 619, 494, 48, 9,
		5106, 144, 12, 12, 2, 759, 813, 5156, 9779, 969, 3, 257, 3, 4910, 65, 1, 907,
		4464, 15, 8685, 54, 48, 762, 7952, 639, 3, 4, 8239, 4, 21, 306, 667, 1, 2, 90,
		42, 6, 1, 3337, 6, 803, 3912, 85, 31, 30, 502, 876, 8686, 813, 880, 5309, 20,
		27, 2523, 266, 101, 8, 3058, 7, 56, 6961, 46, 199, 866, 4, 184, 4, 9675, 92}
	li(nums)
}
