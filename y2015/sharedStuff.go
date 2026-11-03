package y2015

import (
	"bufio"
	"fmt"
	"os"
)

type SharedType struct{}

func ReturnFileContentByChar(testFile string) string {
	myFile, err := os.Open(testFile)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(myFile)
	fileScanner.Split(bufio.ScanWords)

	str := ""
	for fileScanner.Scan() {
		str += fileScanner.Text()
	}

	myFile.Close()
	return str
}

func ReturnFileByLine(testFile string) []string {
	myFile, err := os.Open(testFile)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(myFile)
	fileScanner.Split(bufio.ScanLines)

	str := []string{}
	for fileScanner.Scan() {
		str = append(str, fileScanner.Text())
	}

	return str
}

func Factorial(n int) int {
	if n <= 1 {
		return n
	} else {
		return n*Factorial(n-1)
	}

}
