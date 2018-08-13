package main

import (
	"errors"
	"fmt"
	"strings"
)

type Trie struct {
	Root *TrieNode
}

type TrieNode struct {
	Char     string
	Children map[string]*TrieNode
}

func NewTrie() *Trie {
	t := Trie{}
	t.Root = &TrieNode{Char: "", Children: make(map[string]*TrieNode)}

	return &t
}

func (t *Trie) AddString(s string) (err error) {
	if t.Root == nil {
		err = errors.New("The Trie was not initialized with the proper params, please either use the NewTrie method or initialize the trie with a correct Root and trie again")
	}

	currentNode := *t.Root
	ss := strings.Split(s, "")
	for _, ch := range ss {
		if _, ok := currentNode.Children[ch]; ok == false {
			currentNode.Children[ch] = &TrieNode{Char: ch, Children: make(map[string]*TrieNode)}
		}

		currentNode = *currentNode.Children[ch]
	}

	return nil
}

func main() {
	fmt.Println("vim-go")
	t := Trie{}
	fmt.Println(t.Root == nil)
	fmt.Println(&t == nil)
}
