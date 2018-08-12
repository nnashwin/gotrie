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

func NewTrie() *Trie {
	t := Trie{}
	t.Root = &TrieNode{Char: "", Next: nil}

	return &t
}

func (tn *TrieNode) insertChar(s string) (err error) {
	if tn.Char != "" {
		err = errors.New("The TrieNode already has a character, can not be inserted twice!")
		return
	}

	tn.Char = s

	return
}

func main() {
	fmt.Println("vim-go")
	t := Trie{}
	fmt.Println(t.Root == nil)
	fmt.Println(&t == nil)
}
