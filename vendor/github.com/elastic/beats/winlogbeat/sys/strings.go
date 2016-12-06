package sys

import (
	"fmt"
	"strings"
	"unicode/utf16"
)

// UTF16BytesToString returns a string that is decoded from the UTF-16 bytes.
// The byte slice must be of even length otherwise an error will be returned.
// The integer returned is the offset to the start of the next string with
// buffer if it exists, otherwise -1 is returned.
func UTF16BytesToString(b []byte) (string, int, error) {
	if len(b)%2 != 0 {
		return "", 0, fmt.Errorf("Slice must have an even length (length=%d)", len(b))
	}

	offset := -1

	// Find the null terminator if it exists and re-slice the b.
	if nullIndex := indexNullTerminator(b); nullIndex > 0 {
		if len(b) > nullIndex+2 {
			offset = nullIndex + 2
		}

		b = b[:nullIndex]
	}

	s := make([]uint16, len(b)/2)
	for i := range s {
		s[i] = uint16(b[i*2]) + uint16(b[(i*2)+1])<<8
	}

	return string(utf16.Decode(s)), offset, nil
}

// indexNullTerminator returns the index of a null terminator within a buffer
// containing UTF-16 encoded data. If the null terminator is not found -1 is
// returned.
func indexNullTerminator(b []byte) int {
	if len(b) < 2 {
		return -1
	}

	for i := 0; i < len(b); i += 2 {
		if b[i] == 0 && b[i+1] == 0 {
			return i
		}
	}

	return -1
}

// RemoveWindowsLineEndings replaces carriage return line feed (CRLF) with
// line feed (LF) and trims any newline character that may exist at the end
// of the string.
func RemoveWindowsLineEndings(s string) string {
	s = strings.Replace(s, "\r\n", "\n", -1)
	return strings.TrimRight(s, "\n")
}
