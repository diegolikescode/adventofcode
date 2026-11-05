package y2016

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type crypticMsg struct {
	RawLine string
	Line []string
	SectorID int
	CheckSum []string
}

type CurrProcess string
const (
	LINE CurrProcess = "line"
	SECTOR_ID CurrProcess = "sectorID"
	CHECKSUM CurrProcess = "checksum"
	NONE CurrProcess = "none"
)

const (
	A_BYTE_VALUE int = 97
	Z_BYTE_VALUE int = 122
	ALPHABET_SIZE int = 26
)

func extractLineData(line string, m *SharedType) crypticMsg {
	cm := crypticMsg{
		RawLine: "",
		Line: []string{},
		SectorID: 0,
		CheckSum: []string{},
	}

	var processing = NONE
	// for i:=len(line)-1;i>=0;i-- {
	// 	c := line[i]
	sectorStr := ""
	for _, c := range line {

		if processing == NONE && !m.IsNumber(c) {
			processing = LINE
		}

		if processing == LINE && m.IsNumber(c) {
			processing = SECTOR_ID
		}

		if processing == SECTOR_ID && c == '[' {
			processing = CHECKSUM
			continue
		}

		switch (processing) {
		case LINE:
			if c == '-' {
				cm.RawLine += "-"
				continue
			}

			cm.RawLine += string(c)
			cm.Line = append(cm.Line, string(c))
			continue

		case SECTOR_ID:
			sectorStr += string(c)

		case CHECKSUM:
			if c == ']' {
				continue
			}
			cm.CheckSum = append(cm.CheckSum, string(c))
		}
	}
	v, err := strconv.Atoi(sectorStr)
	if err != nil {
		fmt.Println("err decoding SectorID", err)
	} else {
		cm.SectorID = v
	}

	return cm
}

func isDecoy(line, checksum []string) bool {

	lineMap := make(map[string]int)
	for _, c := range line {
		_, exist := lineMap[c]
		if exist {
			lineMap[c] += 1
		} else {
			lineMap[c] = 1
		}
	}

	letters := make([]string, 0, len(lineMap))
	for k := range lineMap {
		letters = append(letters, k)
	}

	sort.Slice(letters, func(i, j int) bool {
		if lineMap[letters[i]] != lineMap[letters[j]] {
			return lineMap[letters[i]] > lineMap[letters[j]]
		}
		return letters[i] < letters[j]
	})
	// fmt.Println(letters, checksum)
	expected := strings.Join(checksum, "")
	given := strings.Join(letters[:5], "")
	return expected == given
}

func (m SharedType) Day04_A(testFile string) {
	lines := m.ReadFileLines(testFile)

	// fubrjhqlf-edvnhw-dftxlvlwlrq-803[wjvzd]
	// {[f u b r j h q l f e d v n h w d f t x l v l w l r q] 803 [w j v z d]}
	acc := 0
	for _, l := range lines {
		cms := extractLineData(l, &m)
		if isDecoy(cms.Line, cms.CheckSum) {
			acc += cms.SectorID
			// fmt.Println(l)
		}
	}
	fmt.Println(acc)
}

func decryptMsg(line string, sectorID int) string {
	// math.Ceil()
	fullStr := ""
	for _, c := range line {
		if c == '-' {
			fullStr += " "
			continue
		}
		remain := sectorID % ALPHABET_SIZE // is that correct?
		dec := int(c) + remain
		if dec > Z_BYTE_VALUE {
			dec -= ALPHABET_SIZE
		}
		fullStr += string(rune(dec))
	}

	return fullStr
}

func (m SharedType) Day04_B(testFile string) {
	lines := m.ReadFileLines(testFile)

	decrypted := ""
	for _, l := range lines {
		cms := extractLineData(l, &m)
		if isDecoy(cms.Line, cms.CheckSum) {
			// fmt.Println(l)
			msg := decryptMsg(cms.RawLine, cms.SectorID)
			// decrypted = append(decrypted, )
			fmt.Println(msg, cms.SectorID)
			if strings.Contains(msg, "northpole") {
				decrypted = fmt.Sprintf("===== MSG=%s | SECTOR_ID=%d =====", msg, cms.SectorID)
			}
		}
	}
	fmt.Println(decrypted)
}
