package main

import (
	"bufio"
	. "fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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

	re := regexp.MustCompile(`Game (\d+):`)

	var result int

	for _, line := range fileLines {
		match, _ := strconv.Atoi((re.FindStringSubmatch(line))[1])
		mappe := parseGameString(line)
		if check(mappe) {
			result += match
		}
	}
	Println(result)
}

func parseGameString(input string) []map[string]int {

	colonIndex := strings.Index(input, ":")
	input = input[colonIndex+1:]

	gameParts := strings.Split(input, ";")

	var result []map[string]int

	re := regexp.MustCompile(`(\d+)\s*([a-zA-Z]+)`)

	for _, part := range gameParts {
		matches := re.FindAllStringSubmatch(part, -1)

		row := make(map[string]int)

		for _, match := range matches {
			quantity, _ := strconv.Atoi(match[1])
			color := match[2]
			row[color] = quantity
		}
		result = append(result, row)
	}
	return result
}

func check(m []map[string]int) bool {
	disponibility := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	for _, v := range m {
		for color, quantity := range v {
			_, ok := disponibility[color]
			if ok && quantity > disponibility[color] {
				return false
			}
		}
	}
	return true
}
