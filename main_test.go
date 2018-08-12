package main

import (
	"reflect"
	"testing"
)

func TestInsertChar(t *testing.T) {
	expected := &TrieNode{Char: "t"}
	actual := TrieNode{}
	actual.insertChar("t")
	if reflect.DeepEqual(actual.Char, expected.Char) == false {
		t.Error("The insertChar method did not successfully insert the character to the node")
	}
}

func TestNewTrie(t *testing.T) {
	expected := &Trie{Root: &TrieNode{Char: "", Next: nil}}
	actual := NewTrie()

	if reflect.DeepEqual(expected, actual) == false {
		t.Error("The NewTrie function did not create a correct trie")
	}
}
