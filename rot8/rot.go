// Package rot8 provides utilities for rot8000 cyphering of Unicode text.
package rot8

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/rangetable"
)

const Bmp = 0x10000

var (
	// unicode.RangeTable of characters that will not be rotated.
	// This includes white space, control characters and surrogates.
	NonRotate = rangetable.Merge(
		unicode.Z,
		unicode.Cc,
		unicode.Cs,
		unicode.White_Space,
	)

	// Total number of non-rotatable characters.
	NonRotateCount = CountChars(NonRotate)

	// Total number of rotatable characters.
	RotateCount = Bmp - NonRotateCount

	// All rotatable characters.
	Runes []rune

	// Map of characters to its rotated counterpart's index in Runes.
	Map map[rune]int
)

func init() {
	Runes = make([]rune, RotateCount)
	Map = make(map[rune]int, RotateCount)

	rotate := RotateCount / 2

	for n, r := 0, 0; n < Bmp; n++ {
		c := rune(n)

		if !unicode.Is(NonRotate, c) {
			Runes[r] = c
			Map[c] = (r + rotate) % RotateCount
			r++
		}
	}
}

// Rotate string by about 0x8000 code points.
func Rotate(s string) string {
	var sb strings.Builder

	for _, c := range s {
		if r, ok := Map[c]; ok {
			sb.WriteRune(Runes[r])
		} else {
			sb.WriteRune(c)
		}
	}

	return sb.String()
}

// Count total characters in a unicode.RangeTable.
func CountChars(table *unicode.RangeTable) int {
	count := 0

	for _, r := range table.R16 {
		count += int((r.Hi-r.Lo)/r.Stride + 1)
	}

	for _, r := range table.R32 {
		count += int((r.Hi-r.Lo)/r.Stride + 1)
	}

	return count
}
