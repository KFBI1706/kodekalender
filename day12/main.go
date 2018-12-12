package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"unicode/utf8"
)

func main() {
	dat, _ := ioutil.ReadFile("./emoji")
	emojis := strings.Split(string(dat), "\n")[0]
	runes := make([]int, len(emojis)/4)
	for i := 0; i < len(emojis); i += 4 {
		runeValue, _ := utf8.DecodeRuneInString(emojis[i:])
		runes[i/4] = int(runeValue)
	}
	sort.Slice(runes, func(i, j int) bool { return runes[i] > runes[j] }) //Diff min,max = 55
	min := runes[len(runes)-1]
	for i := 32; i < 33; i++ { //Loop 1-128
		for _, r := range emojis {
			fmt.Printf("%c", int(r)-min+i)
		}
		fmt.Printf("\n")
	}
}
