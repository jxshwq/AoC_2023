package main

import (
	"bufio"
	. "fmt"
	"math/big"
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
	var headers []string
	for val := range mappa {
		if string(val[len(val)-1]) == "A" {
			headers = append(headers, val)
		}
	}
	Println(headers)
	for _, header := range headers {
		getNext(mappa, header, direzioni, 0, 0)
	}
	Println(result)
	Println(calculateLCM(result))
}

var result []int

func calculateLCM(numbers []int) *big.Int {
	lcm := big.NewInt(int64(numbers[0]))
	for i := 1; i < len(numbers); i++ {
		temp := big.NewInt(int64(numbers[i]))
		gcd := gcd(lcm, temp)
		lcm.Mul(lcm, temp)
		lcm.Div(lcm, gcd)
	}
	return lcm
}

func gcd(a, b *big.Int) *big.Int {
	gcd := new(big.Int)
	gcd.GCD(nil, nil, a, b)
	return gcd
}

func getNext(mappa map[string]CoppiaValori, valore string, direz string, index, counter int) {
	if index >= len(direz) {
		index = 0
	}
	if string(valore[len(valore)-1]) == "Z" {
		result = append(result, counter)
		return
	}
	counter++
	switch string(direz[index]) {
	case "L":
		index++
		getNext(mappa, mappa[valore].left, direz, index, counter)
	case "R":
		index++
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
