package emoji

import (
	"testing"

	"github.com/guader/emoji/provider"
)

func TestCodePoints(t *testing.T) {
	r, err := New(provider.NewFileProvider(
		"./misc/emoji-sequences.txt",
		"./misc/emoji-zwj-sequences.txt",
	))
	if err != nil {
		panic(err)
	}
	emojis := r.FindAll("(⏩..⏬)(👨‍👩‍👧‍👦)(⏩⏬)123English中あ한국어")
	for _, e := range emojis {
		println(e.String())
	}
}
