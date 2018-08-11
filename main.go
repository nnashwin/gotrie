package main

import "fmt"

type Trie struct {
	Root *TrieNode
}

type TrieNode struct {
	Char string
	Next *TrieNode
}

func main() {
	fmt.Println("vim-go")
}
