package trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	root := New()
	root.Insert([]rune("Hello"))
	root.Insert([]rune("Hello World"))
	fmt.Println("worlds: Hello, World")

	str := []rune("Hello World!")
	for i := len(str); i > 0; i-- {
		source := str[:i]
		fmt.Printf("findIn: %s, exist: %t, short: %s, long: %s\n",
			string(source), root.Exist(source), string(root.MatchShort(source)), string(root.MatchLong(source)))
	}
}
