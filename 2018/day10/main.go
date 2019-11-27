package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

type stack []int

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Push - Essential to all other operations, adds a value onto the stack and returns the stack to the calling function
func (s stack) Push(v int) stack {
	return append(s, v)
}

//Pop = / Fjern en verdi fra stacken.
func (s stack) Pop() (stack, int) {
	l := len(s)
	if l == 0 {
		panic("stack empty")
	}
	return s[:l-1], s[l-1]
}

//SumAll = : Fjern alle verdier på stacken, summer disse, og legg resultatet på stacken.
func (s stack) SumAll() stack {
	var sum int
	for _, v := range s {
		sum += v
	}
	return stack{sum}
}

//Add3Val = | Legg verdien 3 på stacken.
func (s stack) Add3Val() stack {
	return append(s, 3)
}

// Sum2 = ' Fjern to verdier fra stacken, summer dem, legg resultatet på stacken.
func (s stack) Sum2() stack {
	var a, b int
	s, a = s.Pop()
	s, b = s.Pop()
	return append(s, a+b)
}

//Dot  = . Fjern en verdi A, fra stacken. Fjern en verdi B, fra stacken. Legg resultatet av A-B på stacken. Legg B-A på stacken.
func (s stack) Dot() stack {
	var a, b int
	s, a = s.Pop()
	s, b = s.Pop()
	return s.Push(a - b).Push(b - a)
}

//Underscore = _ Fjern en verdi A, fra stacken. Fjern en verdi B, fra stacken. Legg resultatet av A*B på stacken. Legg A på stacken.
func (s stack) Underscore() stack {
	var a, b int
	s, a = s.Pop()
	s, b = s.Pop()
	return s.Push(a * b).Push(a)
}

//Copy = i Legg til en kopi av den siste verdien på stacken.
func (s stack) Copy() stack {
	var v int
	s, v = s.Pop()
	return s.Push(v).Push(v)
}

//Inc = \ Inkrementer den siste verdien på stacken med 1.
func (s stack) Inc() stack {
	var v int
	s, v = s.Pop()
	return s.Push(v + 1)
}

//Mul * Fjern en verdi A, fra stacken. Fjern en verdi B, fra stacken. Beregn A/B, rund av til heltall i retning mot 0, og legg resultatet på stacken.
func (s stack) Mul() stack {
	var a, b int
	s, a = s.Pop()
	s, b = s.Pop()
	return s.Push(int(math.Ceil(float64(a) / float64(b))))
}

//RmEven = ] Fjern en verdi fra stacken. Dersom den fjernede verdien var et partall, legg verdien 1 på stacken.
func (s stack) RmEven() stack {
	var i int
	s, i = s.Pop()
	if i%2 == 0 {
		return s.Push(1)
	}
	return s
}

//RmOdd = [ Fjern en verdi fra stacken. Dersom den fjernede verdien var et oddetall, legg verdien tilbake på stacken.
func (s stack) RmOdd() stack {
	var i int
	s, i = s.Pop()
	if i%2 == 1 {
		return s.Push(i)
	}
	return s
}

//RmThree = ~ Fjern tre verdier fra stacken. Legg tilbake den største av disse.
func (s stack) RmThree() stack {
	var a, b, c int
	s, a = s.Pop()
	s, b = s.Pop()
	s, c = s.Pop()
	l := stack{a, b, c}
	sort.Ints(l)
	return s.Push(l[len(l)-1])
}

func main() {
	dat, err := ioutil.ReadFile("./input.spp")
	check(err)
	lines := strings.Split(string(dat), "\n")
	s := make(stack, 0)
	for i := 0; i < len(lines); i++ {
	Loop:
		for _, c := range lines[i] {
			switch c {
			case ':':
				s = s.SumAll()
			case '|':
				s = s.Add3Val()
			case '\'':
				s = s.Sum2()
			case '.':
				s = s.Dot()
			case '_':
				s = s.Underscore()
			case '/':
				s, _ = s.Pop()
			case 'i':
				s = s.Copy()
			case '\\':
				s = s.Inc()
			case '*':
				s = s.Mul()
			case ']':
				s = s.RmEven()
			case '[':
				s = s.RmOdd()
			case '~':
				s = s.RmThree()
			case ' ':
				s = s.Push(31)
			case 'K':
				break Loop
			}
		}
	}
	sort.Ints(s)
	fmt.Println(s[len(s)-1])
}
