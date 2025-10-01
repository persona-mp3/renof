package soundex

import (
	"fmt"
	"strings"
)

/*
	The whole point for the Soundex Encdoing is to help with Fuzzy Finding.
	All stored documentation, by file names are going to encoded. For example
			ReduceFunc.md encoded as -> R321
			Remove.md -> R510

	And then in the source file/STORE it will be mapped to the path. A highoverview
	would be something like this:
			R321 -> DOCS/views/lib/src/ReduceFunc.md
			R510 -> DOCS/methods/src/Remove.md

	Following that, we implement the Leveschtein Distance Algorithm to get all related/useful
	results to the user as best as possible. Read more: https://en.wikipedia.org/wiki/Levenshtein_distance

	-------------------------------------------
		Rules for American Soundex Encoding
	-------------------------------------------
	Full Reference: https://en.wikipedia.org/wiki/Soundex
	1. The first character must be preserved and all vowels including "y", "w", and "h" are ignored
	2. Replace the other consonants with the following value encoding:
		- b, f, p, v -> 1
		- c, g, j, k, q, s, x, z -> 2
		- d, t -> 3
		- l -> 4
		- m, n -> 5
		- r -> 6
	3a. If two or more letters with the same number are adjacent (before step 1), retain the first letter
	3b. if two letters have the same number are seperated by h,w,y, they should be coded as a single entity
		else if seperated by a vowel, they should be coded twice
	4. If the digits after the first letter is less than 3, pad it with 0's else retain the first 3

*/

var tableEnc = map[rune]string{
	'b': "1", 'f': "1", 'p': "1", 'v': "1",
	'c': "2", 'g': "2", 'j': "2", 'k': "2",
	'q': "2", 's': "2", 'x': "2", 'z': "2",
	'd': "3", 't': "3", 'l': "4",
	'm': "5", 'n': "5", 'r': "6",
}

var delimiters = "hwy"

const (
	CODEC_LEN = 4
)

func isDelimiter(s rune) bool {
	return strings.ContainsRune(delimiters, s)
}

func SoundEnc(s string) string {
	if len(s) < 1 {
		return ""
	} else if len(s) < 2 {
		return s
	}

	var builder strings.Builder
	clone := strings.ToLower(s)

	// Rule 3a.
	_char1, _char2 := rune(clone[0]), rune(clone[1])
	_codec1, _codec2 := tableEnc[_char1], tableEnc[_char2]
	start := 1
	if (len(_codec1) != 0 && len(_codec2) != 0) && _codec1 == _codec2 {
		// so we convert the _char1 back to uppercase, and update where to start from
		char1 := strings.ToUpper(string(_char1))
		builder.WriteString(char1)
		start = 2
	} else {
		builder.WriteRune(rune(s[0]))
	}
	s = strings.ToLower(s[1:])

	next := start + 1
	// Rule 2.
	for ; start < len(s); start++ {
		if next >= len(s) {
			next = 0
		}
		currChar := rune(s[start])
		val, ok := tableEnc[currChar]
		if !ok {
			continue
		}
		// Rule 3b.
		_char1, _char2 = rune(s[start-1]), rune(s[next])
		_codec1, _codec2 = tableEnc[_char1], tableEnc[_char2]

		if isDelimiter(currChar) && _codec1 == _codec2 {
			// coded as a single entry
			builder.WriteString(_codec2)
			continue
		}
		builder.WriteString(val)
	}

	codec := builder.String()
	if len(codec) > CODEC_LEN {
		codec = codec[:CODEC_LEN]
	} else if len(codec) < CODEC_LEN {
		for i := range CODEC_LEN - len(codec) {
			_ = i
			codec += "0"
		}
	}
	fmt.Println("Final Codec -> ", codec)
	return codec
}
