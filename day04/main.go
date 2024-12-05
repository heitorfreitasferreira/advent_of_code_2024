package main

import (
	"bufio"
	"fmt"
	"os"
)

var matrix []string = []string{}
var counter int = 0

func main() {

	file, err := os.Open("input")
	if err != nil {
		panic("Não achei o arquivo")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for i, line := range matrix {
		for j, char := range line {
			if char == 'X' {
				check(i, j)
			}

		}
	}
	fmt.Println(counter)
}

func check(i, j int) {
	dirs := [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}
	for _, dir := range dirs {
		if isWritten(dir, i, j, "XMAS") {
			counter++
		}
	}
}

func canGo(dir [2]int, i, j, offset int) bool {
	maxIOffset := i + dir[0]*offset
	maxJOffset := j + dir[1]*offset

	if maxIOffset < 0 || maxIOffset >= len(matrix) {
		return false
	}

	if maxJOffset < 0 || maxJOffset >= len(matrix[i]) {
		return false
	}
	return true
}

func isWritten(dir [2]int, i, j int, word string) bool {
	if !canGo(dir, i, j, len(word)-1) {
		return false
	}
	for k, char := range word {
		offset := arrMult(dir, k)
		if rune(matrix[i+offset[0]][j+offset[1]]) != char {
			return false
		}
	}
	return true
}

func arrMult(arr [2]int, escalar int) []int {
	a := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		a[i] = arr[i] * escalar
	}
	return a
}
