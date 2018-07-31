package easygo

import (
	"fmt"
	"github.com/cespare/xxhash"
	"strconv"
	"strings"
)

// StringReverse reversed a string
func StringReverse(s string) (reversedString string) {
	for _, v := range s {
		reversedString = string(v) + reversedString
	}
	return
}

// StringStrip strips chars from a string
func StringStrip(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}

// StringToXXHash Retuns a hash for a String
func StringToXXHash(s string) string {
	x := xxhash.Sum64String(s)
	return fmt.Sprintf("%02x", x)
}

// StringToInt Converts to a Pointer
func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}
