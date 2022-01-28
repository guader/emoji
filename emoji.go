package emoji

import (
	"github.com/guader/emoji/trie"
)

type Emoji []rune

type Provider interface {
	CodePoints() ([][]rune, error)
}

// String The string presentation of Emoji.
func (e Emoji) String() string {
	return string(e)
}

type Repository struct {
	tree *trie.Trie
}

// FindAll Find out all the emojis in a string,
// return nil when nothing found.
func (r *Repository) FindAll(s string) []Emoji {
	var (
		rs     = []rune(s)
		length = len(rs)
		i      int
		emojis []Emoji
	)
	for i < length {
		emoji := r.tree.MatchLong(rs[i:])
		if emoji == nil {
			i++
		} else {
			emojis = append(emojis, emoji)
			i += len(emoji)
		}
	}
	return emojis
}

// New Load code points from Provider into a trie.
func New(provider Provider) (*Repository, error) {
	tree := trie.New()
	codePoints, err := provider.CodePoints()
	if err != nil {
		return nil, err
	}
	for _, codePoint := range codePoints {
		tree.Insert(codePoint)
	}
	return &Repository{tree: tree}, nil
}
