package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	r := []rune(str)
	var unpackedStr string
	var numRepeat int
	backslash := false

	for pos, ch := range r {
		if ch == '\\' && !backslash {
			backslash = true
			continue
		}

		if unicode.IsLetter(ch) && backslash {
			return "", ErrInvalidString
		}

		if backslash {
			unpackedStr += string(ch)
			backslash = false
			continue
		}

		if unicode.IsDigit(ch) && pos == 0 {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(ch) && unicode.IsDigit(r[pos-1]) && r[pos-2] != '\\' {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(ch) {
			numRepeat, _ = strconv.Atoi(string(ch))
			if numRepeat == 0 {
				unpackedStr = unpackedStr[:len(unpackedStr)-1]
				continue
			} else if numRepeat > 1 {
				unpackedStr += strings.Repeat(string(r[pos-1]), numRepeat-1)
				continue
			}
		}
		unpackedStr += string(ch)
	}
	return unpackedStr, nil
}
