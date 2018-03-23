// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General
// Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program.  If not, see <http://www.gnu.org/licenses/>.

package linkify

import "unicode"

var (
	unreserved = [256]bool{'-': true, '.': true, '_': true, '~': true}
	basicPunct = [256]bool{'.': true, ',': true, '?': true, '!': true, ';': true, ':': true}
	subdelims  = [256]bool{'!': true, '$': true, '&': true, '\'': true, '(': true, ')': true,
		'*': true, '+': true, ',': true, ';': true, '=': true}
	emailcs = [256]bool{'a': true, 'b': true, 'c': true, 'd': true, 'e': true, 'f': true, 'g': true,
		'h': true, 'i': true, 'j': true, 'k': true, 'l': true, 'm': true, 'n': true, 'o': true,
		'p': true, 'q': true, 'r': true, 's': true, 't': true, 'u': true, 'v': true, 'w': true,
		'x': true, 'y': true, 'z': true, 'A': true, 'B': true, 'C': true, 'D': true, 'E': true,
		'F': true, 'G': true, 'H': true, 'I': true, 'J': true, 'K': true, 'L': true, 'M': true,
		'N': true, 'O': true, 'P': true, 'Q': true, 'R': true, 'S': true, 'T': true, 'U': true,
		'V': true, 'W': true, 'X': true, 'Y': true, 'Z': true, '0': true, '1': true, '2': true,
		'3': true, '4': true, '5': true, '6': true, '7': true, '8': true, '9': true, '!': true,
		'#': true, '$': true, '%': true, '&': true, '\'': true, '*': true, '+': true, '/': true,
		'=': true, '?': true, '^': true, '_': true, '`': true, '{': true, '|': true, '}': true,
		'~': true, '-': true}
)

func isAllowedInEmail(r rune) bool {
	return r < 0x7f && emailcs[r]
}

func isLetterOrDigit(r rune) bool {
	return unicode.In(r, unicode.Letter, unicode.Digit)
}

func isPunctOrSpaceOrControl(r rune) bool {
	return r == '<' || r == '>' || unicode.In(r, unicode.Punct, unicode.Space, unicode.Cc)
}

func isUnreserved(r rune) bool {
	return (r < 0x7f && unreserved[r]) || isLetterOrDigit(r)
}

func isSubDelimiter(r rune) bool {
	return r < 0x7f && subdelims[r]
}
