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

	fileLines := readFile()

	time, distance := getNumbers(fileLines)
	Println(time, distance)

	result := 1
	for v, sec := range time {
		var counter int
		for i := 1; i < sec; i++ {
			if ((sec - i) * i) > distance[v] {
				counter++
			}
		}
		result *= counter
	}
	Println(result)

}

func readFile() []string {
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
	return fileLines
}

func getNumbers(fileLines []string) ([]int, []int) {
	fileLines[0] = strings.ReplaceAll(fileLines[0], " ", "")
	fileLines[1] = strings.ReplaceAll(fileLines[1], " ", "")
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(fileLines[0], -1)
	var time []int
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err == nil {
			time = append(time, num)
		}
	}
	matches = re.FindAllString(fileLines[1], -1)
	var distance []int
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err == nil {
			distance = append(distance, num)
		}
	}
	return time, distance
}
