package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("input")
	if err != nil {
		Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()
	var result int
	for _, line := range fileLines {
		valueOfLine, _ := strconv.Atoi(getNumbers(line))
		result += valueOfLine
	}
	Println(result)
}

func getNumbers(str string) string {
	var subStrings []string

	for i := 0; i < len(str); i++ {
		for j := i + 1; j <= len(str); j++ {
			subString := str[i:j]
			if checkIsNumber(subString) {
				subStrings = append(subStrings, subString)
			}
		}
	}
	n1, n2 := trasformNumbersToInt(subStrings[0], subStrings[len(subStrings)-1])
	return Sprintf("%d%d", n1, n2)
}

var numbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func checkIsNumber(str string) bool {
	if len(str) == 1 {
		if _, err := strconv.Atoi(str); err == nil {
			return true
		}
	} else {
		for _, word := range numbers {
			if word == str {
				return true
			}
		}
	}
	return false
}

func trasformNumbersToInt(str1, str2 string) (n1, n2 int) {
	if len(str1) == 1 {
		n1, _ = strconv.Atoi(str1)
	} else {
		for index, word := range numbers {
			if word == str1 {
				n1 = index + 1
			}
		}
	}
	if len(str2) == 1 {
		n2, _ = strconv.Atoi(str2)
	} else {
		for index, word := range numbers {
			if word == str2 {
				n2 = index + 1
			}
		}
	}
	return
}
