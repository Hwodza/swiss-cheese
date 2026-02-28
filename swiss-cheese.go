package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const (
	swissCheeseHeight = 35
	swissCheeseWidth  = 80
)

type Dimension struct {
	left, right int
}

type Hole struct {
	prob       float64
	dimensions []Dimension
}

type swissCheeseSlice struct {
	slice [swissCheeseHeight][swissCheeseWidth]bool
}

func (cheeseSlice swissCheeseSlice) createSwissSlice(cheeseCH chan []bool) {

	// Location for starting hole
	sr, sc := rand.Intn(swissCheeseHeight), rand.Intn(swissCheeseWidth)
	cheeseSlice.slice[sr][sc] = true
	for i := range cheeseSlice.slice {
		cheeseCH <- cheeseSlice.slice[i][:]
	}
	close(cheeseCH)
}

func generateCheeseOverlay(overlayCH chan []bool) {
	slice := swissCheeseSlice{}
	cheeseCH := make(chan []bool)
	go slice.createSwissSlice(cheeseCH)
	for v := range cheeseCH {
		overlayCH <- v
	}
	close(overlayCH)

}

// Takes in input text (inputCH) and cheese slice by line (overlayCH), then combines the two by overwriting text with cheese
func generateLine(inputCH chan string, overlayCH chan []bool) {
	out := "                                                                                "
	for i := range inputCH {
		runes := []rune(i)
		out = ""
		cheese := <-overlayCH
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

// Function that will color cheese to orange, works by checking for #, and then coloring them
func cheeseify(s string) string {
	const (
		yellow = "\033[33m"
		reset  = "\033[0m"
	)

	var b strings.Builder
	inSwissSlice := false

	for _, ch := range s {
		if ch == '#' {
			if !inSwissSlice {
				b.WriteString(yellow)
				inSwissSlice = true
			}
		} else {
			if inSwissSlice {
				b.WriteString(reset)
				inSwissSlice = false
			}
		}
		b.WriteRune(ch)
	}

	if inSwissSlice {
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
	overlayCH := make(chan []bool)

	scanner := bufio.NewScanner(file)

	go func() {
		for scanner.Scan() {
			inputCH <- scanner.Text()
		}
		close(inputCH)
	}()

	go generateCheeseOverlay(overlayCH)

	generateLine(inputCH, overlayCH)
}
