package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

var matrix []string = []string{}
var counter int = 0

//go:embed test
var input string

func main() {

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for i := 1; i < len(matrix)-1; i++ {
		for j := 1; j < len(matrix[i])-1; j++ {
			overlap(i, j)
		}
	}
	fmt.Println(counter)
}

func check(i, j int) {
	dirs := [][2]int{
		// {-1, 0},
		// {1, 0},
		// {0, -1},
		// {0, 1},
		// {-1, -1},
		{-1, 1},
		// {1, -1},
		{1, 1},
	}
	var vibe = true
	for _, dir := range dirs {
		if !isWritten(dir, i, j, "MAS") {
			vibe = false
			break
		}
	}
	if vibe {
		counter++
	}
}

func overlap(i, j int) {
	kerneis := [][][]rune{
		{
			{'M', ' ', 'S'},
			{' ', 'A', ' '},
			{'M', ' ', 'S'},
		},
	}
	for range 3 {
		kerneis = append(kerneis, rotate(kerneis[len(kerneis)-1]))
	}

	for _, kernel := range kerneis {
		match := true
		for ki := range kernel {
			for kj := range kernel[ki] {
				if kernel[ki][kj] == ' ' {
					continue
				}
				if kernel[ki][kj] != rune(matrix[i-1+ki][j-1+kj]) {
					match = false
					break
				}
			}
			if !match {
				break
			}
		}
		if match {
			counter++
		}
	}
}

func rotate(kernel [][]rune) [][]rune {
	n := len(kernel)
	rotated := make([][]rune, n)
	for i := range rotated {
		rotated[i] = make([]rune, n)
	}
	for i := range kernel {
		for j := range kernel[i] {
			rotated[j][n-1-i] = kernel[i][j]
		}
	}
	return rotated
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

func arrMult(arr [2]int, escalar int) []int {
	a := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		a[i] = arr[i] * escalar
	}
	return a
}
