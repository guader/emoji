# Emoji

Load official emoji sequence files from site: https://www.unicode.org/Public/emoji/ into a tree,

each rune of a code point is stored as a node of the tree,

then we could find out all the code points within a string.

## Example

```go
package main

import (
	"github.com/guader/emoji"
	"github.com/guader/emoji/provider"
)

func main() {
	r, err := emoji.New(provider.NewFileProvider(
		"./misc/emoji-sequences.txt",
		"./misc/emoji-zwj-sequences.txt",
	))
	if err != nil {
		panic(err)
	}
	emojis := r.FindAllEmojis("(β©..β¬)(π¨βπ©βπ§βπ¦)(β©β¬)123EnglishδΈ­γνκ΅­μ΄")
	for _, e := range emojis {
		println(e.String())
	}
	/* Output:
		β©
		β¬
		π¨βπ©βπ§βπ¦
		β©
		β¬
	*/
}
```
