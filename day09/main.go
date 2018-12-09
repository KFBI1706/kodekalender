package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type hash struct {
	Char               string `json:"ch"`
	Hash               string `json:"hash"`
	PrevHash, NextHash *hash
}

var chain []hash

func main() {
	chain := make([]hash, 0)
	dat, err := ioutil.ReadFile("./input-hashchain.json")
	check(err)
	err = json.Unmarshal(dat, &chain)
	check(err)
	for i := 0; i < len(chain); i++ {
		for j := 0; j < len(chain); j++ {
			if fmt.Sprintf("%x", md5.Sum([]byte(chain[j].Hash+chain[i].Char))) == chain[i].Hash {
				chain[j].NextHash = &chain[i]
				chain[i].PrevHash = &chain[j]
			}
		}
	}
	var solution string
	var h *hash
	for i := 0; i < len(chain); i++ {
		if chain[i].PrevHash == nil {
			h = &chain[i]
			break
		}
	}
	for {
		solution += h.Char
		if h.NextHash == nil {
			break
		}
		h = h.NextHash
	}
	fmt.Println(solution)
}
