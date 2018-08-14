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

	// Adds a dollar sign to indicate the string is over
	currentNode.Children["$"] = &TrieNode{Char: "$", Children: make(map[string]*TrieNode)}

	return nil
}

func (t *Trie) DoesContain(s string) (isContained bool) {
	isContained = true
	if t.Root == nil {
		return
	}

	currentNode := *t.Root
	s += "$"
	ss := strings.Split(s, "")
	for _, ch := range ss {
		if _, ok := currentNode.Children[ch]; ok == false {
			isContained = false
			return
		}

		currentNode = *currentNode.Children[ch]
	}

	return
}

func main() {
	fmt.Println("vim-go")
	t := Trie{}
	fmt.Println(t.Root == nil)
	fmt.Println(&t == nil)
}
