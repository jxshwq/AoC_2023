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

	var result int

	for _, line := range fileLines {
		mappe := parseGameString(line)
		result += check(mappe)

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

func check(maps []map[string]int) int {
	piuPiccoli := make(map[string]int)
	power := 1
	for _, singleMap := range maps { //itero su tutte le mappe, prendendone una alla volta
		for color, quantity := range singleMap { // for range sulla mappa singola
			_, ok := piuPiccoli[color]              // se il colore è presente nella mappa "piùpiccoli" E la quantità è più
			if ok && quantity > piuPiccoli[color] { // grande di quella dello stesso colore nella mappa
				piuPiccoli[color] = quantity // allora il valore nella mappa di quel colore viene settato alla quantità corrente
			} else if !ok { // invece se nella mappa "più piccoli" non è presente il valore
				piuPiccoli[color] = quantity // setto nella mappa ""più piccoli" il valore corrente
			}
		}
	}
	Println(piuPiccoli)
	for _, quantity := range piuPiccoli {
		power *= quantity
	}
	Println(power)
	return power
}
