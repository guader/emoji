package trie

import (
	"fmt"
)

type Trie struct {
	matched  bool
	children map[rune]*Trie
}

// String The string presentation in json format of Trie.
func (t *Trie) String() string {
	s := "{"
	s += fmt.Sprintf(`"match":%t,`, t.matched)
	for n, t2 := range t.children {
		s += fmt.Sprintf(`"%d":`, n)
		s += t2.String()
		s += ","
	}
	s += "}"
	return s
}

// New New Trie root.
func New() *Trie {
	return &Trie{
		matched:  false,
		children: make(map[rune]*Trie),
	}
}

// Insert Insert a rune slice to trie,
// each rune is stored as a node of the trie,
// the last node of the rune slice would be set to be matched.
func (t *Trie) Insert(rs []rune) {
	node := t
	for _, r := range rs {
		// insert node when rune not found in the trie
		if node.children[r] == nil {
			node.children[r] = New()
		}
		// prepare the node for next rune
		node = node.children[r]
	}
	// set the last node to be matched
	node.matched = true
}

// Exist Test whether a rune slice exists in the trie as a matched slice.
func (t *Trie) Exist(rs []rune) bool {
	node := t
	var ok bool
	for _, r := range rs {
		node, ok = node.children[r]
		// all the runes should be found
		if !ok {
			return false
		}
	}
	return node.matched
}

// MatchShort Find a rune slice in the trie,
// return the shortest matched rune slice.
func (t *Trie) MatchShort(rs []rune) []rune {
	var (
		node  = t
		ok    bool
		runes []rune
	)
	// find the runes in the trie
	for _, r := range rs {
		node, ok = node.children[r]
		if !ok {
			break
		}
		if node.matched {
			runes = append(runes, r)
			break
		} else {
			// minus stands for a rune unmatched
			runes = append(runes, -r)
		}
	}

	if len(runes) == 0 || runes[len(runes)-1] < 0 {
		return nil
	}
	return runes
}

// MatchLong Find a rune slice in the trie greedily,
// return the longest matched rune slice.
func (t *Trie) MatchLong(rs []rune) []rune {
	var (
		node  = t
		ok    bool
		runes []rune
	)
	// find the runes in the trie greedily
	for _, r := range rs {
		node, ok = node.children[r]
		if !ok {
			break
		}
		if node.matched {
			runes = append(runes, r)
		} else {
			// minus stands for a rune unmatched
			runes = append(runes, -r)
		}
	}

	var result []rune
	// pick up the longest matched runes
	for i := len(runes) - 1; i >= 0; i-- {
		// continue until any positive(matched) rune exists
		if runes[i] < 0 {
			continue
		}
		// i shows the index of the last matched rune
		// load the runes to the result and finish finding
		for _, r := range runes[:i+1] {
			if r < 0 {
				result = append(result, -r)
			} else {
				result = append(result, r)
			}
		}
		break
	}
	return result
}
