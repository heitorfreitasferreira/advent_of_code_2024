package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func massageData() ([][2]int, [][]int) {
	inputBytes, err := os.ReadFile("input")
	if err != nil {
		panic("erro ao ler o arquivo")
	}
	rules := strings.Split(string(inputBytes), "\n\n")
	precedencesStr := strings.Split(rules[0], "\n")
	updatesStr := strings.Split(rules[1], "\n")

	precedences := make([][2]int, len(precedencesStr))
	for i, prec := range precedencesStr {
		numsStr := strings.Split(prec, "|")
		a, _ := strconv.Atoi(numsStr[0])
		b, _ := strconv.Atoi(numsStr[1])

		precedences[i][0] = a
		precedences[i][1] = b
	}

	updates := make([][]int, len(updatesStr))
	for i, upd := range updatesStr {
		parsed := []int{}
		for _, numStr := range strings.Split(upd, ",") {
			num, _ := strconv.Atoi(numStr)
			parsed = append(parsed, num)
		}
		updates[i] = parsed
	}

	return precedences, updates

}
func main() {
	precedences, updates := massageData()
	mustPrecede := make(map[int]map[int]bool)
	mustSucceed := make(map[int]map[int]bool)

	for _, tuple := range precedences {
		a, b := tuple[0], tuple[1]

		if mustPrecede[b] == nil {
			mustPrecede[b] = make(map[int]bool)
		}
		if mustSucceed[a] == nil {
			mustSucceed[a] = make(map[int]bool)
		}
		// fixando B, se A aparecer tem que ser antes que B
		mustPrecede[b][a] = true
		// C A B -> OK
		// B A C -> RUIM
		mustSucceed[a][b] = true
	}

	isValid := func(update []int) bool {

		for i, pivot := range update {
			left := update[:i]
			right := update[i+1:]

			for _, v := range right {
				if mustPrecede[pivot][v] {
					return false
				}
			}
			for _, v := range left {
				if mustSucceed[pivot][v] {
					return false
				}
			}
		}
		return true
	}

	validUpdates := [][]int{}
	invalidUpdates := [][]int{}
	for _, updt := range updates {
		if isValid(updt) {
			validUpdates = append(validUpdates, updt)
		} else {
			invalidUpdates = append(invalidUpdates, updt)
		}
	}

	accum := 0

	for _, upd := range validUpdates {
		accum += upd[len(upd)/2]
	}
	fmt.Println(accum)

}
