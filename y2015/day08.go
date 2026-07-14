package y2015

import (
	"fmt"
)

/*
========== the C character is an ansii char ==========
"\xa8br\x8bjr\""                      -> CbrCjr" (len 7) |
"nq"
"zjrfcpbktjmrzgsz\xcaqsc\x03n\"huqab"
"daz\\zyyxddpwk"
"draes\xa2n\\g\x27ek\"lj\"\\viqych"
"nnx\\krnrfomdnt\x2flbl\xd2xpo\"cp\"k"
"kwdaapalq"
"u\"ptk"
"ckhorczuiudfjmmcc\\u\"wozqxibsfjma"

*/

func isHexDigit(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
}

func (m SharedType) Day08_A(testFile string) {
	lines := ReturnFileByLine(testFile)

	code := 0
	mem := 0
	for _, l := range lines {
		code += len(l)
		strMem := 0
		// fmt.Printf("LINE: %s | LEN_LITERAL: %d\n", l, len(l))
		for i := 0; i < len(l); i++ {
			c := l[i]
			if i == 0 || i == len(l)-1 {
				continue
			}

			if c == '\\' {
				// cases to cover:
				//   \xNN -> count as 1, advance 4
				// X \"   -> count as 1, advance 2
				// X \\   -> count as 1, advance 2
				// X \?   -> count as 2, advance 2 (invalid escape sequence)
				switch l[i+1] {
				case '"':
					strMem++
					i += 1
				case '\\':
					strMem++
					i += 1
				case 'x':
					if i+3 < len(l) && isHexDigit(l[i+2]) && isHexDigit(l[i+3]) {
						strMem++
						i += 3
					} else {
						strMem += 2
						i += 2
					}
				default:
					strMem += 2
					i += 1
				}
			} else {
				strMem++
			}
		}
		mem += strMem
		fmt.Printf("LINE: %s | LEN_LITERAL: %d | LEN_MEM: %d\n", l, len(l), strMem)
	}
	fmt.Printf("TOTAL CODE: %d | TOTAL MEM: %d | DIFF: %d\n", code, mem, code-mem)
}

func (m SharedType) Day08_B(testFile string) {
	lines := ReturnFileByLine(testFile)

	originalCount := 0
	encodedCount := 0
	for _, l := range lines {
		originalCount += len(l)
		encCount := 2
		for _, c := range l {
			switch c {
			case '"', '\\':
				encCount += 2
			default:
				encCount++
			}
		}
		encodedCount += encCount
		fmt.Printf("LINE: %s | LEN_ORIGINAL: %d | LEN_ENCODED: %d\n", l, len(l), encCount)
	}
	fmt.Printf("TOTAL ORIGINAL: %d | TOTAL ENCODED: %d | DIFF: %d\n", originalCount, encodedCount, encodedCount-originalCount)
}
