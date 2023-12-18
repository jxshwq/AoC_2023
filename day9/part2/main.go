package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Apri il file
	file, err := os.Open("/home/josh/Coding/adventOfCode/2023/day9/part2/input")
	if err != nil {
		fmt.Println("Errore nell'apertura del file:", err)
		return
	}
	defer file.Close()

	// Leggi il file riga per riga
	scanner := bufio.NewScanner(file)
	var rows [][]int
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, " ")
		row := make([]int, len(numbers))
		for i, numStr := range numbers {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Errore nella conversione dell'intero:", err)
				return
			}
			row[i] = num
		}
		rows = append(rows, row)
	}

	// Somma le righe per ottenere un singolo array

	var result []int
	for _, row := range rows {
		for i, num := range row {
			if i >= len(result) {
				result = append(result, num)
			} else {
				result[i] += num
			}
		}
	}

	var n int
	n = (getResult(getNext(result, 1)))
	fmt.Println(n)
}

// funzione che calcola il "previous value" per ogni riga di input
func getResult(result []int) int {
	for i := len(result) - 1; i > 0; i-- {
		result[i-1] = result[i-1] - result[i]
	}
	return result[0]
}

// funzione ricorsiva che calcola il "next value" per ogni riga di input
func getNext(result []int, counter int) []int {
	for i := len(result) - 1; i >= counter; i-- {
		result[i] = result[i] - result[i-1]
	}
	counter++
	if result[len(result)-1] == 0 {
		fmt.Println(result)
		return result
	} else {
		return getNext(result, counter)
	}
}
