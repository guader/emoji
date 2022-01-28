package provider

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// DecodeText Decode a text to a slice of code points,
// text should be in format of `RUNE_HEX..RUNE_HEX` as a code point range,
// or `RUNE_HEX RUNE_HEX ...` as a code point,
// return nil when no code points found.
func DecodeText(text string) [][]rune {
	rss := make([][]rune, 0)
	if strings.Contains(text, "..") {
		// handle code point range
		codePointRange := strings.SplitN(text, "..", 2)
		if len(codePointRange) != 2 {
			return nil
		}
		from, err := strconv.ParseInt(codePointRange[0], 16, 46)
		if err != nil {
			return nil
		}
		to, err := strconv.ParseInt(codePointRange[1], 16, 64)
		if err != nil {
			return nil
		}
		for ; from <= to; from++ {
			rss = append(rss, []rune{rune(from)})
		}
	} else {
		// handle code point
		runeStrs := strings.Split(text, " ")
		rs := make([]rune, 0)
		for _, runeStr := range runeStrs {
			r, err := strconv.ParseInt(runeStr, 16, 64)
			if err != nil {
				return nil
			}
			rs = append(rs, rune(r))
		}
		rss = append(rss, rs)
	}
	return rss
}

// DecodeFile Decode file content to a slice of code points,
// file should be provided by the official site: https://www.unicode.org/Public/emoji/.
func DecodeFile(filename string) ([][]rune, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rss := make([][]rune, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		fields := strings.SplitN(line, ";", 2)
		if len(fields) != 2 {
			continue
		}
		rss = append(rss, DecodeText(strings.TrimSpace(fields[0]))...)
	}
	return rss, nil
}

type FileProvider struct {
	filenames []string
}

func (p *FileProvider) CodePoints() ([][]rune, error) {
	var codePoints [][]rune
	for _, filename := range p.filenames {
		rss, err := DecodeFile(filename)
		if err != nil {
			return nil, err
		}
		codePoints = append(codePoints, rss...)
	}
	return codePoints, nil
}

func NewFileProvider(filenames ...string) *FileProvider {
	return &FileProvider{filenames: filenames}
}
