package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	up      = '^'
	down    = 'v'
	right   = '>'
	left    = '<'
	empty   = '.'
	blocked = '#'
)

var (
	n = [2]int{-1, 0}
	s = [2]int{1, 0}
	w = [2]int{0, -1}
	e = [2]int{0, 1}
)

type guard struct {
	i, j               int
	dir                [2]int
	obstacles, visited [][]bool
}

func (g *guard) next_dir() {
	switch g.dir {
	case n:
		g.dir = e
	case e:
		g.dir = s
	case s:
		g.dir = w
	case w:
		g.dir = n
	default:
		panic("invalid dir")
	}
}

func newGuard(input string) *guard {
	lines := strings.Split(input, "\n")
	junin := guard{
		obstacles: make([][]bool, len(lines)),
		visited:   make([][]bool, len(lines)),
	}

	for i, line := range lines {
		junin.obstacles[i] = make([]bool, len(line))
		junin.visited[i] = make([]bool, len(line))
		for j, chr := range line {
			switch chr {
			case blocked:
				junin.obstacles[i][j] = true
			case left:
				junin.j = j
				junin.i = i
				junin.dir = w
				junin.visited[i][j] = true
			case right:
				junin.j = j
				junin.i = i
				junin.dir = e
				junin.visited[i][j] = true
			case up:
				junin.j = j
				junin.i = i
				junin.dir = n
				junin.visited[i][j] = true
			case down:
				junin.j = j
				junin.i = i
				junin.dir = s
				junin.visited[i][j] = true
			}
		}
	}
	return &junin
}
func (junin *guard) out_of_bounds() bool {
	return junin.i < 0 || junin.j < 0 || junin.i >= len(junin.visited) || junin.j >= len(junin.visited[0])
}
func (junin *guard) will_be_out_of_bounds() bool {
	nextI := junin.i + junin.dir[0]
	nextJ := junin.j + junin.dir[1]
	return nextI < 0 || nextJ < 0 || nextI >= len(junin.obstacles) || nextJ >= len(junin.obstacles[nextI])
}

func (junin *guard) update() int {

	// Marca o lugar como visitado
	junin.visited[junin.i][junin.j] = true

	// Checa se o prox move vai tirar do mapa
	if junin.will_be_out_of_bounds() {
		return junin.count()
	}

	// muda direção se tem obstaculo
	if junin.obstacles[junin.i+junin.dir[0]][junin.j+junin.dir[1]] {
		junin.next_dir()
		return junin.update()
	}

	// faz ele andar pra frente
	junin.i += junin.dir[0]
	junin.j += junin.dir[1]
	return junin.update()
}
func (junin *guard) count() int {
	couter := 0
	for _, line := range junin.visited {
		for _, went := range line {
			if went {
				couter++
			}
		}
	}
	return couter
}
func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		panic("erro ao ler o arquivo")
	}
	junin := newGuard(string(input))
	fmt.Println(junin.update())
}
