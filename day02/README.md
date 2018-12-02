# Day 2 - 47x Performance improvement, while killing readability and reuseability

This was a really interesting task, after not being able to authenticate and view the task there were still some problems to come;

The first solution I created is in main.go and it shows a balance between readable and short code, but not really that performant
* side note, I debugged this code for like 30 minutes wondering what was wrong, when I found out that I've read the input in the wrong order.

After finally solving the challenge I was really happy until I saw the runtime.. Turns out FscanF is really slow, even rob pike hates it. Well I didn't know so I have to optimize
* https://github.com/golang/go/issues/12275#issuecomment-133796990

Usually I optimize my code a little bit after a solve and committ the changes, but this time I went a little crazy, chasing as low exec time as possible. The result of that is what you see below.
* This isn't good code at all, but it's interesting that the functionality is preserved and the exec time went from 0.9 to 0.19 all because of these optimization

```go
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	cnt := make(map[float32]int)
	file, err := ioutil.ReadFile("./input-rain.txt")
	check(err)
	for _, line := range strings.Split(string(file), "\n") {
		nums := strings.FieldsFunc(line, func(r rune) bool {
			switch r {
			case '(', ')', ',', ';':
				return true
			}
			return false
		})
		if len(nums) != 4 {
			break
		}
		x1, _ := strconv.Atoi(nums[0])
		y1, _ := strconv.Atoi(nums[1])
		x2, _ := strconv.Atoi(nums[2])
		y2, _ := strconv.Atoi(nums[3])
		cnt[float32(y2-y1)/float32(x2-x1)]++
	}
	var big int
	for _, i := range cnt {
		if i > big {
			big = i
		}
	}
	fmt.Println(big)
}
```

pprof cpu graphs compared:
* Unoptimized and readable: https://loot.kfbi.xyz/9517a0f2-60df-4fdb-87f5-5142526a7deb.jpg

* Optimized: https://loot.kfbi.xyz/4e43568e-5d6c-4bfe-8268-19578ed3168d.jpg
