package y2016

import (
	"fmt"
	"strconv"
	"strings"
)

func isTriangle(a,b,c int) bool {
	if a+b > c && a+c > b && b+c > a  {
		return true
	}
	return false
}

func str2int(s string) int {
	n, err :=strconv.Atoi(s)
	if err != nil {
		fmt.Printf("err converting string=%s\n", err)
	}
	return n
}

func (m SharedType) Day01_A(testFile string) {
	lines := m.ReadFileLines(testFile)

	count := 0
	for _, l := range lines {
		s := strings.Fields(l)
		s0 := str2int(s[0])
		s1 := str2int(s[1])
		s2 := str2int(strings.TrimSpace(s[2]))

		if isTriangle(s0, s1, s2) {
			count++
		}
	}
	fmt.Println(count)
}

func add2matrix(v1, v2, v3 int, matrix [][]int, idx int) {
	matrix[0][idx] = v1
	matrix[1][idx] = v2
	matrix[2][idx] = v3
}

func howManyTriangles(matrix [][]int) int {
	c := 0
	for _, s := range matrix {
		if isTriangle(s[0], s[1], s[2]) {
			c++
		}
	}
	return c
}

func newEmptyMatrix() [][]int {
	m := make([][]int, 3)
	m[0] = make([]int, 3)
	m[1] = make([]int, 3)
	m[2] = make([]int, 3)

	return m
}

func (m SharedType) Day01_B(testFile string) {
	lines := m.ReadFileLines(testFile)

	count := 0
	acc := newEmptyMatrix()
	idxCounter := 0
	for _, l := range lines {
		s := strings.Fields(l)
		s0 := str2int(s[0])
		s1 := str2int(s[1])
		s2 := str2int(s[2])
		add2matrix(s0, s1, s2, acc, idxCounter)
		idxCounter++
		count += howManyTriangles(acc)

		if idxCounter >= 3 {
			acc = newEmptyMatrix()
			idxCounter = 0
		}
	}
	fmt.Println(count)
}
