package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var a, b, c, inc int
	rps := []rune{'R', 'P', 'S'}
	rpsc := make([]int, len(rps))
	standings := []*int{&a, &b, &c}
	dat, _ := ioutil.ReadFile("./input-rpslog.txt")
	s := string(dat)
	inc = 3
	for i := 0; i < len(s)-2; i += inc {
		cs := s[i : i+inc]
		if !((inc == 3 && (s[i] != s[i+1] && s[i] != s[i+2] && s[i+1] != s[i+2]) || (s[i] == s[i+1] && s[i] == s[i+2])) || (inc == 2 && s[i] == s[i+1])) {
			var mx, ix1, ix2, cx, d int //Counts for rps, set max and the index of that max. And ix2 to the other tool in RSP that's not ix1
			for j, t := range rps {
				rpsc[j] = strings.Count(cs, string(t))
				if rpsc[j] >= mx {
					mx = rpsc[j]
					ix1 = j
				}
			}
			if ix2, cx = ix1+2%3, ix1; rpsc[(ix1+1)%3] > rpsc[(ix1+2)%3] {
				ix2 = (ix1 + 1) % 3
			}
			d = (3 + ix1 - ix2) % 3 //If d == 1 ix1 won, if d == 2 ix2 won
			switch d {
			case 1:
				if inc == 2 { //ix1 won
					cx = ix1
				} else { //2 of 3 won, so inc is set to 2
					tmpStandings := make([]*int, 0)
					for j, c := range cs {
						if rps[ix1] == c {
							tmpStandings = append(tmpStandings, standings[j])
						}
					}
					standings = tmpStandings
					i++
					inc = 2
					continue
				}
			case 2: //ix2 vant
				cx = ix2
			}
			for j, c := range cs {
				if rps[cx] == c {
					(*standings[j])++
					break
				}
			}
			standings = []*int{&a, &b, &c}
			if inc == 2 {
				i--
			}
			inc = 3
		}
	}
	fmt.Printf("%d,%d,%d\n", *standings[0], *standings[1], *standings[2])
}
