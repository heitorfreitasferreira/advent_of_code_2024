package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	id1 := []int{}
	id2 := []int{}
	file, err := os.Open("input")
	if err != nil {
		panic("Não achei o arquivo")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strsValues := strings.Split(line, "   ")
		v1, err := strconv.Atoi(strsValues[0])
		if err != nil {
			panic("não consegui converter para int")
		}
		v2, err := strconv.Atoi(strsValues[len(strsValues)-1])
		if err != nil {
			panic("não consegui converter para int")
		}

		id1 = append(id1, v1)

		id2 = append(id2, v2)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// slices.Sort(id1)
	// slices.Sort(id2)
	// totalDistance := 0
	// for i := 0; i < len(id1); i++ {
	// 	maior := id1[i]
	// 	menor := id2[i]
	// 	if id2[i] > id1[i] {
	// 		maior = id2[i]
	// 		menor = id1[i]
	// 	}
	// 	totalDistance += maior - menor
	// }
	// fmt.Println(id1[:10])
	// fmt.Println(id2[:10])

	// fmt.Println(totalDistance) // 2344935

	frequency := make(map[int]int)

	for _, v := range id2 {
		frequency[v]++
	}

	total := 0
	for _, id := range id1 {
		total += id * frequency[id]
	}

	fmt.Println(total)
}
