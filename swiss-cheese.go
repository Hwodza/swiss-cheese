package main

import (
	"bufio"
	"fmt"
	"os"
)

func generate_line(inputCH chan string) {
	for i := range inputCH {
		fmt.Println(i)
	}
}

func main() {
	file, err := os.Open("lorem.txt")
	if err != nil {
		panic(err)
	}

	inputCH := make(chan string)

	scanner := bufio.NewScanner(file)

	go func() {
		for scanner.Scan() {
			inputCH <- scanner.Text()
		}
		close(inputCH)
	}()

	generate_line(inputCH)
}
