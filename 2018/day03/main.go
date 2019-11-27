package main

import "fmt"

//MAX 2^32
const MAX = 4294967296

var primes []int
var primesToFind int

func juletall(x1, n, index int) int {
	var cnt int
	if n == 13 {
		return 1
	}
	for i, x2 := range primes[index:] {
		if prod := x1 * x2; prod < MAX {
			cnt += juletall(prod, n+1, index+i)
		} else {
			break
		}
	}
	return cnt
}

func main() {
	//2^32/2^24=2^8=512 primes
	primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509}
	//2^11*3^13 < 4294967296 < 2^10*3^14
	fmt.Println(juletall(2048, 0, 0))
}
