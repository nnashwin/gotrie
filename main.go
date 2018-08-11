package main

import (
	"errors"
	"fmt"
)

type Trie struct {
	Root *TrieNode
}

type TrieNode struct {
	Char string
	Next *TrieNode
}

func (t *TrieNode) insertChar(s string) (err error) {
	if t.Char == "" {
		t.Char = s
	} else {
		err = errors.New("The TrieNode already has a character, can not be inserted twice!")
		return
	}

	return
}

func main() {
	fmt.Println("vim-go")
}
