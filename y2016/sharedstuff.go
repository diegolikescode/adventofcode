package y2016

import (
	"os"
	"strings"
)

type SharedType struct {}

func (s *SharedType) ReadFile(filename string) string {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func (s *SharedType) ReadFileLines(filename string) []string {
	fileContent := s.ReadFile(filename)

	lines := strings.Split(fileContent, "\n")
	splitted := []string{}
	for _, l := range lines {
		if len(strings.TrimSpace(l)) > 0 {
			splitted = append(splitted, l)
		}
	}
	return splitted
}
