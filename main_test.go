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
