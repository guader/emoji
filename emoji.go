package emoji

import (
	"github.com/guader/emoji/trie"
)

type Emoji []rune

type Provider interface {
	CodePoints() (codePoints [][]rune, err error)
}

// String The string presentation of Emoji.
func (e Emoji) String() string {
	return string(e)
}

type Repository struct {
	tree *trie.Trie
}

// FindOne Find the first emoji in a string,
// return nil when no emoji found.
func (r *Repository) FindOne(s string) Emoji {
	var (
		rs     = []rune(s)
		length = len(rs)
		i      int
	)
	for i < length {
		emoji := r.tree.MatchLong(rs[i:])
		if emoji != nil {
			return emoji
		}
		i++
	}
	return nil
}

// FindAll Find out all the emojis in a string,
// return nil when no emoji found.
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
func New(p Provider) (*Repository, error) {
	tree := trie.New()
	codePoints, err := p.CodePoints()
	if err != nil {
		return nil, err
	}
	for _, codePoint := range codePoints {
		tree.Insert(codePoint)
	}
	return &Repository{tree: tree}, nil
}
