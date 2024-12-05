package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string) []int {
	report := []int{}
	partial := strings.Split(line, " ")
	for _, str := range partial {
		n, _ := strconv.Atoi(str)
		report = append(report, n)
	}
	return report

}

func processesReport(levels []int) int {
	ascending := levels[0] < levels[1]
	for i := 0; i < len(levels)-1; i++ {
		var max, min int
		if ascending {
			max = levels[i+1]
			min = levels[i]
		} else {
			max = levels[i]
			min = levels[i+1]
		}
		diff := max - min
		if diff < 1 || diff > 3 {
			return 0
		}
	}
	return 1
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic("Não achei o arquivo")
	}

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		report := parseLine(line)
		another := processesReport(report)
		if another == 0 {
			for i := 0; i < len(report); i++ {
				reportWithoutIesim := removeIPos(i, report)
				if try := processesReport(reportWithoutIesim); try == 1 {
					another = 1
					break
				}
			}
		}
		total += another
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(total)
	test := []int{0, 1, 2, 3, 4}
	for i := range len(test) {
		heyyy := removeIPos(i, test)
		fmt.Println(heyyy)
	}
}

func removeIPos(pos int, arr []int) []int {
	returnable := []int{}
	for i, v := range arr {
		if pos == i {
			continue
		}
		returnable = append(returnable, v)
	}
	return returnable
}
