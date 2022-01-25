package emoji

import (
	"testing"
)

func TestCodePoints(t *testing.T) {
	cp, err := New("./misc/emoji-sequences.txt", "./misc/emoji-zwj-sequences.txt")
	if err != nil {
		panic(err)
	}
	emojis := cp.Match("(⏩..⏬)(👨‍👩‍👧‍👦)(⏩⏬)123English中あ한국어")
	for _, emoji := range emojis {
		println(emoji.String())
	}
}
