package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

// Definisci un tipo che rappresenta i due valori associati a una chiave
type CoppiaValori struct {
	left  string
	right string
}

func main() {
	direzioni, fileLines := readFile()
	mappa := make(map[string]CoppiaValori)
	for _, line := range fileLines[1:] {
		n := strings.Index(line, "=")
		leftRight := strings.Split(line[n+3:len(line)-1], ",")
		leftRight[0] = strings.ReplaceAll(leftRight[0], " ", "")
		leftRight[1] = strings.ReplaceAll(leftRight[1], " ", "")
		header := strings.ReplaceAll(line[:n], " ", "")
		mappa[header] = CoppiaValori{left: leftRight[0], right: leftRight[1]}
	}
	getNext(mappa, "AAA", direzioni, 0, 0)
}

func getNext(mappa map[string]CoppiaValori, valore string, direz string, index, counter int) {
	if index >= len(direz) {
		index = 0
	}
	if valore == "ZZZ" {
		Println(counter)
		return
	}
	counter++
	switch string(direz[index]) {
	case "L":
		index++
		Println(mappa[valore].left)
		getNext(mappa, mappa[valore].left, direz, index, counter)
	case "R":
		index++
		Println(mappa[valore].right)
		getNext(mappa, mappa[valore].right, direz, index, counter)
	}
}

// mappa, un valore, indice dell'indicazione++,

func readFile() (string, []string) {
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
	return fileLines[0], fileLines[1:]
}
