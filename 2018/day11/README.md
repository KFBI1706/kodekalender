# Day 11 - Tired presumptions

This day started off with a message from a friend at 06:07 saying that he'd allready solved it..

Now that all my competitive nature was out of the window, I still had to solve the task, but as i turns out my brain wasn't set on the same thing.

For some reason, I'm guessing tiredness I assumed there were multiple packages, therefore multiple lines in the text file and that the instructions were ordered. All of these were wrong.

The input was just a single line of distances between 1-9 accompabined by a direction. Santa didn't even travel >9 one way before changing direction. I assumed he'd atleast do it sometimes and decided to account for that with a much slower solution.

All in all it was a fun task and after my brain had booted up it took me a total of 12 minutes to whip up this solution in go. So I'm quite happy.

What I've learned today is bascially, don't rush a task when you're tired, that'll often lead to more bad presumptions, compared to if you just wake up, then do the task.

```go
package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	var x, y, n int16
	var LEN, i uint32
	var b byte
	LEN = 20000000
	b = 48
	dat, err := ioutil.ReadFile("/tmp/input-crisscross.txt")
	check(err)
	for i = 1; i < LEN; i += 2 {
		n = int16(dat[i-1] & ^b)
		switch dat[i] {
		case 72: //H 64, 8
			x += n
		case 86: //V 64, 16, 4, 2
			x -= n
		case 70: //F 64, 4, 2
			y += n
		case 66: //B 64, 2
			y -= n
		}
	}
	fmt.Printf("[%d,%d]\n", x, y)
}
```

And a really fast version, with lots of optimizations, now with a 250ms runtime. Still 100ms behind C so if anyone finds an optimization, ping me!
