package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

const (
	INIT_HEIGHT = 35
	INIT_LENGTH = 80
	NUM_CIRCLES = 2
)

type circle struct {
	name       string
	num_rows   int
	num_cols   int
	char_avail []int
	edges      [][]string
}

type empty_range struct {
	start int
	end   int
}

type row struct {
	empty []empty_range
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

func addHole(hole row, slice [][]bool, sr int, sc int) [][]bool {
	for i := sr; i < min(len(slice), len(hole.empty)+sr); i++ {
		for j := max(hole.empty[i-sr].start+sc, 0); j < (min(len(slice[i]), hole.empty[i-sr].end+sc)); j++ {
			slice[i][j] = true
		}
	}

	return slice
}

func createInitSlice(holes []row) [][]bool {
	slice := make([][]bool, INIT_HEIGHT)
	for i := range slice {
		slice[i] = make([]bool, INIT_LENGTH)
	}

	// Starting hole
	sr, sc := rand.Intn(INIT_HEIGHT), rand.Intn(INIT_LENGTH)
	nHole := rand.Intn(NUM_CIRCLES)
	slice[sr][sc] = true
	slice = addHole(holes[nHole], slice, sr, sc)

	return slice
}

func printSlice(slice [][]bool) {
	for _, row := range slice {
		for _, v := range row {
			if v {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func main() {
	holes := []row{
		{[]empty_range{{0, 2}, {-1, 3}, {0, 2}}},
		{[]empty_range{{0, 2}, {-4, 6}, {-8, 10}, {-10, 12}, {-11, 13}, {-12, 14}, {-12, 14}, {-12, 14}, {-11, 13}, {-10, 12}, {-8, 10}, {-4, 6}, {0, 2}}},
	}
	slice := createInitSlice(holes)
	printSlice(slice)

	// Static slice
	// Reading in Standard in
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		break
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//
	// cheeseSlice := []row{
	// 	{[]empty_range{{75, 78}}},
	// 	{[]empty_range{{74, 79}}},
	// 	{[]empty_range{{5, 8}, {75, 78}}},
	// 	{[]empty_range{{4, 9}}},
	// 	{[]empty_range{{5, 8}, {28, 31}, {59, 60}}},
	// 	{[]empty_range{{27, 32}, {55, 64}}},
	// 	{[]empty_range{{28, 31}, {51, 69}}},
	// 	{[]empty_range{{49, 71}}},
	// 	{[]empty_range{{48, 72}}},
	// 	{[]empty_range{{47, 73}}},
	// 	{[]empty_range{{47, 73}}},
	// 	{[]empty_range{{47, 73}}},
	// 	{[]empty_range{{18, 19}, {48, 72}}},
	// 	{[]empty_range{{14, 25}, {49, 71}}},
	// 	{[]empty_range{{10, 29}, {51, 69}}},
	// 	{[]empty_range{{8, 31}, {55, 64}}},
	// 	{[]empty_range{{7, 32}, {59, 60}}},
	// 	{[]empty_range{{6, 33}}},
	// 	{[]empty_range{{6, 33}, {69, 72}}},
	// 	{[]empty_range{{6, 33}, {68, 73}}},
	// 	{[]empty_range{{7, 32}, {69, 72}}},
	// 	{[]empty_range{{8, 31}}},
	// 	{[]empty_range{{10, 29}, {38, 41}}},
	// 	{[]empty_range{{14, 25}, {37, 42}}},
	// 	{[]empty_range{{18, 19}, {38, 41}, {66, 67}}},
	// 	{[]empty_range{{62, 71}}},
	// 	{[]empty_range{{58, 75}}},
	// 	{[]empty_range{{56, 77}}},
	// 	{[]empty_range{{55, 78}}},
	// 	{[]empty_range{{54, 79}}},
	// 	{[]empty_range{{14, 17}, {54, 79}}},
	// 	{[]empty_range{{13, 18}, {54, 79}}},
	// 	{[]empty_range{{14, 17}, {55, 78}}},
	// 	{[]empty_range{{56, 77}}},
	// 	{[]empty_range{{58, 75}}},
	// }
	// for i, line := range cheeseSlice {
	// 	out := "################################################################################"
	//
	// 	for _, slice := range line.empty {
	// 		replace := "                                                                                "
	// 		if slice.end > len(lines[i]) {
	// 			replace = lines[i] + replace[len(lines[i]):]
	// 		} else {
	// 			replace = lines[i]
	// 		}
	// 		out = out[:slice.start] + replace[slice.start:slice.end] + out[slice.end:]
	// 		if len(lines[i]) > 80 {
	// 			out = out + lines[i][80:]
	// 		}
	// 	}
	// 	fmt.Println(cheeseify(out))
	// }
	// for _, l := range lines {
	// 	fmt.Println(l)
	// }
}
