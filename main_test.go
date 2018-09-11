package main

import (
	"reflect"
	"sort"
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
	// It adds strings correctly to the data structure
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

	// Correctly adds a "$" to signal the ending of a word
	if _, ok := node.Children["$"]; ok == false {
		t.Errorf("The AddString function did not add the correct ending character \"$\" to the data structure.")
	}
}

func TestDoesContain(t *testing.T) {
	tr := NewTrie()
	ss := "tyler"
	// "tyler" should not be contained in the trie when nothing has been added (trie.Root == nil)
	expected := false
	actual := tr.DoesContain(ss)

	tr.AddString(ss)

	// after adding it to the trie, "tyler" should be contained in the trie
	expected = true
	actual = tr.DoesContain(ss)

	if actual != expected {
		t.Errorf("DoesContain should return true for strings contained in the trie.  String: %s", ss)
	}

	// tyl, which has not been added to the trie, should not be contained
	ss = "tyl"
	expected = false
	actual = tr.DoesContain(ss)

	if actual != expected {
		t.Errorf("DoesContain should return false for strings not contained in the trie.  String: %s", ss)
	}
}

func TestGetStartsWith(t *testing.T) {
	tr := NewTrie()
	tr.AddString("cookies")
	tr.AddString("coconuts")
	tr.AddString("collate")
	tr.AddString("coconut")
	tr.AddString("chocolate")
	expected := []string{"coconut", "coconuts", "cookies", "collate"}
	actual := tr.GetStartsWith("co")
	sort.Strings(actual)
	sort.Strings(expected)

	if reflect.DeepEqual(expected, actual) == false {
		t.Errorf("The GetStartsWith method did not return all strings that start with the prefix.  Expected: %q\nActual: %q\n", expected, actual)
	}
}
