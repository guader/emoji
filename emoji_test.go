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

	const str = "(⏩..⏬)(👨‍👩‍👧‍👦)(⏩⏬)123English中あ한국어"

	emojis := r.FindAll(str)
	for _, e := range emojis {
		println("all: ", e.String())
	}

	e := r.FindOne(str)
	println("one: ", e.String())
}
