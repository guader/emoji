# Emoji

Load official emoji sequence files from site: https://www.unicode.org/Public/emoji/ into a tree,

each rune of a code point is stored as a node of the tree,

then we could find out all the code points within a string.

## Example

```go
package main

import "github.com/guader/emoji"

func main() {
	cp, err := emoji.New("./misc/emoji-sequences.txt", "./misc/emoji-zwj-sequences.txt")
	if err != nil {
		panic(err)
	}
	emojis := cp.Match("(⏩..⏬)(👨‍👩‍👧‍👦)(⏩⏬)123English中あ한국어")
	for _, e := range emojis {
		println(e.String())
	}
	/* Output:
		⏩
		⏬
		👨‍👩‍👧‍👦
		⏩
		⏬
	*/
}
```
