package emoji

import (
	"github.com/guader/emoji/sequence"
	"github.com/guader/emoji/trie"
)

type Emoji []rune

// String The string presentation of Emoji.
func (e Emoji) String() string {
	return string(e)
}

type CodePoints struct {
	tree *trie.Trie
}

// Match Find out all the emojis in a string,
// return nil when nothing found.
func (cp *CodePoints) Match(s string) []Emoji {
	var (
		rs     = []rune(s)
		length = len(rs)
		i      int
		emojis []Emoji
	)
	for i < length {
		emoji := cp.tree.Match(rs[i:])
		if emoji == nil {
			i++
		} else {
			emojis = append(emojis, emoji)
			i += len(emoji)
		}
	}
	return emojis
}

// New Load code points from official files into a trie.
func New(filenames ...string) (*CodePoints, error) {
	tree := trie.New()
	for _, filename := range filenames {
		rss, err := sequence.DecodeFile(filename)
		if err != nil {
			return nil, err
		}
		for _, rs := range rss {
			tree.Insert(rs)
		}
	}
	return &CodePoints{
		tree: tree,
	}, nil
}
