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
	fmt.Println(a, b, c)
	inc = 3
	for i := 0; i < len(s)-2; i += inc {
		cs := s[i : i+inc]
		fmt.Println(i, cs)
		if (inc == 3 && (s[i] != s[i+1] && s[i] != s[i+2] && s[i+1] != s[i+2]) || (s[i] == s[i+1] && s[i] == s[i+2])) || (inc == 2 && s[i] == s[i+1]) {
			fmt.Println("tied")
			//when tied, inc doesn't change
		} else {
			//s[i] == s[i+1] || s[i] == s[i+2]
			var mx, ix1, ix2, d int
			for j, t := range rps {
				rpsc[j] = strings.Count(cs, string(t))
				if rpsc[j] >= mx {
					mx = rpsc[j]
					ix1 = j
				}
			}
			if ix2 = ix1 + 2%3; rpsc[(ix1+1)%3] > rpsc[(ix1+2)%3] {
				ix2 = (ix1 + 1) % 3
			}
			d = (3 + ix1 - ix2) % 3
			//fmt.Println(d, mx, ix1)
			switch d {
			case 1:
				if inc == 2 { //ix1 won
					for j, c := range cs {
						if rps[ix1] == c {
							(*standings[j])++
							break
						}
					}
					fmt.Printf("1:%s won \n", string(rps[ix1]))
					standings = []*int{&a, &b, &c}
					for _, s := range standings {
						fmt.Printf("%d", *s)
					}
					fmt.Println()
					if inc == 2 {
						i--
					}
					inc = 3
				} else { //2 of 3 won, so inc is set to 2
					i++
					tmpStandings := make([]*int, 0)
					for j, c := range cs {

						if rps[ix1] == c {
							fmt.Printf("appending %d %d\n", j, *standings[j])
							tmpStandings = append(tmpStandings, standings[j])
						}
					}
					standings = tmpStandings
					fmt.Println(len(standings))
					inc = 2
				}
			case 2: //ix2 vant
				for j, c := range cs {
					if rps[ix2] == c {
						(*standings[j])++
						break
					}
				}
				fmt.Printf("2:%s won \n", string(rps[ix2]))
				standings = []*int{&a, &b, &c}
				for _, s := range standings {
					fmt.Printf("%d,", *s)
				}
				fmt.Println()
				if inc == 2 {
					i--
				}
				inc = 3
			default:
				panic("0000")
			}

		}
	}
	for _, s := range standings {
		fmt.Printf("%d,", *s)
	}
	fmt.Println()
}
