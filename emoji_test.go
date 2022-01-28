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

	emojis := r.FindAllEmojis(str)
	for _, e := range emojis {
		println("all:", e.String())
	}

	e := r.FindOneEmoji(str)
	println("one:", e.String())

	println("contains:", r.ContainsEmoji(str))
	println("contains:", r.ContainsEmoji("(123English中あ한국어)"))
}
