package main

import (
	"bufio"
	. "fmt"
	"os"
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
	// var result int
	mappa := make(map[int]int)
	for i, line := range fileLines {
		var winsOfTheGame int

		// rimuovo tutta la parte prima dei due punti compresi
		index := strings.Index(line, ":")
		if index != -1 {
			line = line[index+1:]
		}

		// divido la stringa in due parti, i miei numeri e i numeri da giocare, convertendoli poi in array di int
		game := strings.Split(line, "|")
		winsOfTheGame = scratchNumbers(getNums(strings.Split(game[0], " "), strings.Split(game[1], " ")))
		mappa[i+1] = winsOfTheGame
	}

	var quantities []int
	for i := 0; i < len(mappa); i++ {
		quantities = append(quantities, 1)
	}
	Println(mappa, "\n", quantities)

	for j := 1; j < len(mappa); j++ {
		for i := j; i < j+mappa[j]; i++ {
			if _, ok := mappa[i]; ok {
				quantities[i] += quantities[j-1]
			}
		}
		Println(quantities)
	}
	var result int

	Println(quantities, mappa)
	for _, v := range quantities {
		result += v
	}
	Println(result)

}

func scratchNumbers(myNumbers, luckyNumbers []int) int {
	var wins int
	mappa := make(map[int]int)

	// inserisco in una mappa, con valore zero, tutte le keys corrisponidenti ai numeri fortnuati
	for _, v := range luckyNumbers {
		mappa[v] = 0
	}

	// itero sui miei fortunati, se ognuno di essi corrisponde ad una chiave già presente nella mappa, aumento il valore di esso
	for _, v := range myNumbers {
		if _, ok := mappa[v]; ok {
			mappa[v]++
		}
	}

	// ritorno il valore "wins" che raddoppia ogni volta che viene trovata una cifra vincente
	for k, v := range mappa {
		if k != 0 && v != 0 {
			wins++
		}
	}
	return wins
}

// funzione che trasforma gli array con i numeri in formati string in array di int, semplificandone l'uso
func getNums(m, l []string) (myNumbers, luckyNumbers []int) {
	for _, v := range m {
		num, _ := strconv.Atoi(v)
		myNumbers = append(myNumbers, num)
	}
	for _, v := range l {
		num, _ := strconv.Atoi(v)
		luckyNumbers = append(luckyNumbers, num)
	}
	return myNumbers, luckyNumbers
}
