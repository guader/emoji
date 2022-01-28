package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/guader/emoji/provider"
)

func main() {
	fp := provider.NewFileProvider(
		"./misc/emoji-sequences.txt",
		"./misc/emoji-zwj-sequences.txt",
	)
	codePoints, err := fp.CodePoints()
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("./provider/slice_provider_code_points.go", os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString("package provider\nvar sliceProviderCodePoints = [][]rune{"); err != nil {
		panic(err)
	}

	for _, cp := range codePoints {
		var cpstrs []string
		for _, r := range cp {
			cpstrs = append(cpstrs, fmt.Sprintf("0x%X", r))
		}
		if _, err = f.WriteString(fmt.Sprintf("	{%s},\n", strings.Join(cpstrs, ","))); err != nil {
			panic(err)
		}
	}
	if _, err = f.WriteString("}\n"); err != nil {
		panic(err)
	}
}
