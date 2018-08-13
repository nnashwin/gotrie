package main

import (
	"reflect"
	"testing"
)

func TestNewTrie(t *testing.T) {
	expected := &Trie{Root: &TrieNode{Char: "", Children: make(map[string]*TrieNode)}}
	actual := NewTrie()

	if reflect.DeepEqual(expected, actual) == false {
		t.Error("The NewTrie function did not create a correct trie")
	}
}

func TestAddString(t *testing.T) {
	tr := NewTrie()
	ss := "tyler"

	tr.AddString(ss)

	node := tr.Root
	for _, ch := range ss {
		// Convert to string to use correctly in the Children map
		strCh := string(ch)
		if _, ok := node.Children[strCh]; ok == false {
			t.Errorf("The AddString function did not write the correct string to the trie. The Children map should contains %s and did not", strCh)
		}
		node = (node.Children[strCh])
	}
}
