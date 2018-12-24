package main

import (
	"fmt"
	"math/big"
	"unicode/utf8"
)

func main() {
	var r rune
	s := "GODJULOGGODTNYTTÅR"
	ans := big.NewInt(0)
	for i, w, val := 0, 0, int64(0); i < len(s); i += w {
		r, w = utf8.DecodeRuneInString(s[i:])
		if val = int64(r - 'A' + 1); w == 2 {
			val -= 104
		}
		prod := big.NewInt(val)
		uct, e := big.NewInt(29), big.NewInt(int64(len(s)-i-2))
		uct.Exp(uct, e, nil)
		ans.Add(ans, prod.Mul(prod, uct))
	}
	fmt.Println(ans)
}
