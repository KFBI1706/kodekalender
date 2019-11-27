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
		cs := s[i : i+inc] //Current string
		if !((inc == 3 && (s[i] != s[i+1] && s[i] != s[i+2] && s[i+1] != s[i+2]) || (s[i] == s[i+1] && s[i] == s[i+2])) || (inc == 2 && s[i] == s[i+1])) {
			var mx, ix1, ix2, cx, d int //mx, ix1, ix2 d - max value rpsc, index1, index2, index for both tools, diff - diff for determining which side won
			for j, t := range rps {     //Count rock paper scissors and order ix1,ix2
				if rpsc[j] = strings.Count(cs, string(t)); rpsc[j] >= mx {
					mx = rpsc[j]
					ix1 = j
				}
			}
			if ix2, cx = ix1+2%3, ix1; rpsc[(ix1+1)%3] > rpsc[(ix1+2)%3] { //Assign ix2
				ix2 = (ix1 + 1) % 3
			}
			d = (3 + ix1 - ix2) % 3 //If d == 1 ix1 won, if d == 2 ix2 won
			switch d {
			case 1:
				if inc == 2 { //ix1 won and there's only two people playing
					cx = ix1
				} else { //ix1 won, but 2 of 3 won, so inc is set to 2 and the standings array is adjusted
					tmpStandings := make([]*int, 0)
					for j, c := range cs {
						if rps[ix1] == c {
							tmpStandings = append(tmpStandings, standings[j])
						}
					}
					standings = tmpStandings
					i++
					inc = 2  //only check the next two chars, as only two players are playing against each other and none won
					continue //continue so standings are not calculated or reset and preventing inc from being reset
				}
			case 2: //ix2 won, and ix2 is always 1 so ix2 has won regardless of 2 or 3 players
				cx = ix2
			}
			for j, c := range cs { //Loop the current standings
				if rps[cx] == c { //winning index equals the position in the current string.and therefore the player
					(*standings[j])++ //increment the player score
					break             //only one winner, break
				}
			}
			standings = []*int{&a, &b, &c} //The standings may have been modified to a length of two, but restore it regardless
			if inc == 2 {                  //We add inc for each iteration, but if we also drecrement it we need to manually increment i
				i--
			}
			inc = 3
		}
	}
	fmt.Printf("%d,%d,%d\n", *standings[0], *standings[1], *standings[2])
}
