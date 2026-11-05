package y2016

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	real_input="ojvtpuvg"
	hex_prefix="00000"
	pass_len=8

	PASSONE = "abc3231929"
	PASSTWO = "abc5017308"
	PASSTHREE = "abc5278568"

	TEST_RESULT = "18f47a30"
)

var PROPER_HASHES = []string{"abc3231929", "abc5017308", "abc5278568"}
var N_PROPER_HASHES = []int{3231929, 5017308,5278568}

func decode2str(src string) string {
	h := md5.New()

	io.WriteString(h, src)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (m SharedType) Day05_A(testfile string) {
	pass := ""
	for i:=0;true;i++ {
		str := decode2str(real_input+strconv.Itoa(i))
		if str[:5] == hex_prefix {
			s := string(str[5])
			pass += s
			fmt.Println(i, str, s)

			if len(pass) >= pass_len {
				fmt.Println(pass)
				break
			}
		}
	}
}

func handleSixthChar(s []string, str string) (bool, []string) {
	idxStr := string(str[5])
	idx, err := strconv.Atoi(idxStr)
	if err != nil {
		return false, s
	}

	if idx > 7 || s[idx] != "" {
		return false, s
	}

	s[idx] = string(str[6])
	return true, s
}

func (m SharedType) Day05_B(testfile string) {
	sp := make([]string, pass_len)
	cout := 0
	for i:=0;true;i++ {
		str := decode2str(real_input+strconv.Itoa(i))

		if str[:5] == hex_prefix {
			success, sp := handleSixthChar(sp, str)
			if success {
				cout++
				fmt.Println(str, sp, cout)
			}

			if cout >= pass_len {
				pass := strings.Join(sp, "")
				fmt.Println(pass)
				break
			}
		}
	}
}
