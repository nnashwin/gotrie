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
	if t.Char != "" {
		err = errors.New("The TrieNode already has a character, can not be inserted twice!")
		return
	}

	t.Char = s

	return
}

func main() {
	fmt.Println("vim-go")
}
