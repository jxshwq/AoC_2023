package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"unicode"
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
	for lineIndex, line := range fileLines {
		var digitFound bool
		var indexStart, indexFinish int
		for index, char := range line {
			if unicode.IsNumber(char) && !digitFound {
				digitFound = true
				indexStart = index
			}
			if index == len(line)-1 && unicode.IsNumber(char) {
				digitFound = false
				indexFinish = index
				result += getValue(fileLines, lineIndex, indexStart, indexFinish)
			}
			if !unicode.IsNumber(char) && digitFound {
				digitFound = false
				indexFinish = index - 1
				result += getValue(fileLines, lineIndex, indexStart, indexFinish)
			}
		}
	}
	Println(result)
}

func getValue(fileLines []string, lineIndex, indexStart, indexFinish int) int {
	var contorno string
	var result int
	Println(lineIndex, indexStart, indexFinish)
	if lineIndex > 0 {
		contorno += string(fileLines[lineIndex-1][indexStart : indexFinish+1])
		if indexStart > 0 {
			contorno += string(fileLines[lineIndex-1][indexStart-1])
		}
		if indexFinish < len(fileLines[0])-1 {
			contorno += string(fileLines[lineIndex-1][indexFinish+1])
		}
	}
	if lineIndex < len(fileLines)-1 {
		contorno += string(fileLines[lineIndex+1][indexStart : indexFinish+1])
		if indexStart > 0 {
			contorno += string(fileLines[lineIndex+1][indexStart-1])
		}
		if indexFinish < len(fileLines[0])-1 {
			contorno += string(fileLines[lineIndex+1][indexFinish+1])
		}
	}
	if indexStart > 0 {
		contorno += string(fileLines[lineIndex][indexStart-1])
	}
	if indexFinish < len(fileLines[0])-1 {
		contorno += string(fileLines[lineIndex][indexFinish+1])
	}
	for _, char := range contorno {
		if char != '.' && !unicode.IsDigit(char) {
			value, _ := strconv.Atoi(fileLines[lineIndex][indexStart : indexFinish+1])
			result += value
			return result
		}
	}
	return 0
}
