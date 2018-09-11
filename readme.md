# gotrie
[![Build Status](https://travis-ci.org/ru-lai/gotrie.svg?branch=master)](https://travis-ci.org/ru-lai/gotrie)

## Install
```
$ go get github.com/ru-lai/gotrie
```

## Usage
```
import (
	trie "github.com/ru-lai/gotrie"
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
