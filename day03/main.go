package main

import (
	"fmt"
	"os"
	"strconv"
)

func isDigit(r byte) bool {
	return map[byte]bool{
		'0': true,
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
	}[r]
}

var doing bool = true
var state int = 0
var counter int64 = 0
var p1, p2 string

func consume(char byte) {
	flush := func() {
		state = 0
		p1 = ""
		p2 = ""
	}
	switch state {
	case -7:
		if char == ')' {
			doing = false
			flush()
		}
	case -6:
		if char == '(' {
			state--
		} else {
			flush()
		}
	case -5:
		if char == 't' {
			state--
		} else {
			flush()
		}
	case -4:
		if char == '\'' {
			state--
		} else {
			flush()
		}
	case -3:
		if char == ')' {
			doing = true
			flush()
		}
	case -2:
		if char == '(' {
			state--
		} else if char == 'n' {
			state = -4
		} else {
			flush()
		}
	case -1:
		if char == 'o' {
			state--
		} else {
			flush()
		}
	case 0:
		if char == 'm' && doing {
			state++
		} else if char == 'd' {
			state--
		} else {
			flush()
		}
	case 1:
		if char == 'u' {
			state++
		} else {
			flush()
		}
	case 2:
		if char == 'l' {
			state++
		} else {
			flush()
		}
	case 3:
		if char == '(' {
			state++
		} else {
			flush()
		}
	case 4:
		if isDigit(char) {
			p1 += string(char)
		} else if char == ',' {
			state++
		} else {
			flush()
		}
	case 5:
		if isDigit(char) {
			p2 += string(char)
		} else if char == ')' {
			n1, err := strconv.Atoi(p1)
			if err != nil {
				fmt.Println(p1, p2)
				panic("deu ruim")
			}
			n2, err := strconv.Atoi(p2)
			if err != nil {
				panic("deu ruim 2")
			}
			counter += int64(n1) * int64(n2)
			flush()
		} else {
			flush()
		}

	default:
		flush()
	}
}

func main() {

	file, err := os.Open("input")
	if err != nil {
		panic("n abriu")
	}
	defer file.Close()

	buffer := make([]byte, 4096)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			break
		}
		for i := 0; i < n; i++ {
			consume(buffer[i])
		}
	}
	fmt.Println(counter)
}
