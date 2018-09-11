# go-trie
[![Build Status](https://travis-ci.org/ru-lai/go-trie.svg?branch=master)](https://travis-ci.org/ru-lai/go-trie)

## Install
```
$ go get github.com/ru-lai/go-trie
```

## Usage
```
import (
	trie "github.com/ru-lai/go-trie"
)

tr := trie.NewTrie()
tr.AddString("crush")
tr.AddString("crushes")

tr.DoesContain("crush")
// => true

tr.DoesContain("crushed")
// => false

tr.GetStartsWith("cr")
// => []string{"crush", "crushed"}
```
