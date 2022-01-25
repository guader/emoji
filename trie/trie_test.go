package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	rs := []rune("Hello World!")
	root := New()
	root.Insert(rs)

	println(root.String())

	search := append(rs, []rune(" Tire")...)
	for i := len(search); i > 0; i-- {
		target := search[:i]
		println(string(target), "(", root.Exist(target), ")", "(", string(root.Match(target)), ")")
	}
}
