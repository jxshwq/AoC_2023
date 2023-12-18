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
	var result int
	for _, line := range fileLines {

		// rimuovo tutta la parte prima dei due punti compresi
		index := strings.Index(line, ":")
		if index != -1 {
			line = line[index+1:]
		}

		// divido la stringa in due parti, i miei numeri e i numeri da giocare, convertendoli poi in array di int
		game := strings.Split(line, "|")
		result += scratchNumbers(getNums(strings.Split(game[0], " "), strings.Split(game[1], " ")))
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
			if wins == 0 {
				wins = 1
			} else {
				wins *= 2
			}
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
