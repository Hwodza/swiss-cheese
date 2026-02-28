package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	SWISS_CHEESE_HEIGHT = 35
	SWISS_CHEESE_WIDTH  = 80
)

func createSwissSlice(cheeseCH chan []bool) {
	cheeseSlice := make([][]bool, SWISS_CHEESE_HEIGHT)
	for i := range cheeseSlice {
		cheeseSlice[i] = make([]bool, SWISS_CHEESE_WIDTH)
	}
	for i := range cheeseSlice {
		cheeseCH <- cheeseSlice[i]
	}
	close(cheeseCH)
}

func generateLine(inputCH chan string, cheeseCH chan []bool) {
	out := "                                                                                "
	for i := range inputCH {
		runes := []rune(i)
		out = ""
		cheese := <-cheeseCH
		if cheese == nil {
			fmt.Println(i)
			continue
		}
		for j, v := range cheese {
			if v {
				if j >= len(runes) {
					out += " "
				} else {
					out += string(runes[j])
				}
			} else {
				out += "#"
			}
		}

		fmt.Println(cheeseify(out))
	}
}

func cheeseify(s string) string {
	const (
		yellow = "\033[33m"
		reset  = "\033[0m"
	)

	var b strings.Builder
	inHash := false

	for _, ch := range s {
		if ch == '#' {
			if !inHash {
				b.WriteString(yellow)
				inHash = true
			}
		} else {
			if inHash {
				b.WriteString(reset)
				inHash = false
			}
		}
		b.WriteRune(ch)
	}

	if inHash {
		b.WriteString(reset)
	}

	return b.String()
}

func main() {
	file, err := os.Open("lorem.txt")
	if err != nil {
		panic(err)
	}

	inputCH := make(chan string)
	cheeseCH := make(chan []bool)

	scanner := bufio.NewScanner(file)

	go func() {
		for scanner.Scan() {
			inputCH <- scanner.Text()
		}
		close(inputCH)
	}()

	go createSwissSlice(cheeseCH)

	generateLine(inputCH, cheeseCH)
}
