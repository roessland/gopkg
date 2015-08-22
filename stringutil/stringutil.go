package stringutil

import "unicode/utf8"
import "strings"

// Pad a string to a certain length with some padding character.
// If the string is longer than the desired length, nothing is done.
// UTF-8 friendly.
func PadLeft(str string, totLength int, pad rune) string {
    strLen := utf8.RuneCountInString(str)

    if strLen > totLength {
        // String is long enough -- do nothing
        return str
    } else {
        // Prepend the padding string
        return strings.Repeat(string(pad), totLength - strLen) + str
    }
}

// Pad a string to a certain length with some padding character.
// If the string is longer than the desired length, nothing is done.
// UTF-8 friendly.
func PadRight(str string, totLength int, pad rune) string {
    strLen := utf8.RuneCountInString(str)

    if strLen > totLength {
        // String is long enough -- do nothing
        return str
    } else {
        // Append the padding string
        return str + strings.Repeat(string(pad), totLength - strLen)
    }
}
