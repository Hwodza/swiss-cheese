package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strings"
)

const (
	INIT_SLICE_HEIGHT = 35
	SLICE_LENGTH      = 80
	NUM_CIRCLES       = 2
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

type hole struct {
	empty []empty_range
	prob float64
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

func addHole(hole hole, slice [][]bool, sr int, sc int) [][]bool {
	for i := sr; i < min(len(slice), len(hole.empty)+sr); i++ {
		for j := max(hole.empty[i-sr].start+sc, 0); j < (min(len(slice[i]), hole.empty[i-sr].end+sc)); j++ {
			slice[i][j] = true
		}
	}

	return slice
}

func bfs(slice [][]bool, sr, sc int) float64 {
	rows := len(slice)
	if rows == 0 {
		log.Fatal("BFS given an empty array")
	}
	if sr < 0 || sr >= rows || sc < 0 || sc >= SLICE_LENGTH {
		log.Fatal("BFS starting point out of bounds")
	}
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, SLICE_LENGTH)
	}

	queue := [][2]int{{sr, sc}}
	visited[sr][sc] = true

	for len(queue) > 0 {
		r, c := queue[0][0], queue[0][1]
		queue = queue[1:]

		if slice[r][c] {
			return math.Abs(float64(sr-r)) + math.Abs(float64(sc-c))
		}

		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < SLICE_LENGTH && !visited[nr][nc] {
				visited[nr][nc] = true
				queue = append(queue, [2]int{nr, nc})
			}
		}
	}
	return SLICE_LENGTH
}

func chooseHole(holes []hole) hole {
	total := 0.0
	for _, hole := range holes {
		total += hole.prob
	}
	r := rand.Float64() * total

	accum := 0.0
	for _, hole := range holes {
		accum += hole.prob
		if r <= accum {
			return hole
		}
	}
	return holes[0]
}

func createInitSlice(holes []hole) [][]bool {
	slice := make([][]bool, INIT_SLICE_HEIGHT)
	for i := range slice {
		slice[i] = make([]bool, SLICE_LENGTH)
	}

	// Starting hole
	sr, sc := rand.Intn(INIT_SLICE_HEIGHT), rand.Intn(SLICE_LENGTH)
	slice[sr][sc] = true
	slice = addHole(chooseHole(holes), slice, sr, sc)
	maxDistance := float64(SLICE_LENGTH / 1)
	for i := range INIT_SLICE_HEIGHT - 1 {
		for j := range SLICE_LENGTH - 1 {
			if slice[i][j] {
				continue
			}
			distance := bfs(slice, i, j)
			prob := distance / maxDistance

			if prob > 1 {
				prob = 1
			}
			if rand.Float64() < prob {
				slice[i][j] = true
				slice = addHole(chooseHole(holes), slice, i, j)
			}
		}
	}

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
	holes := []hole{
		{[]empty_range{{0, 2}, {-1, 3}, {0, 2}}, .8},
		{[]empty_range{{0, 2}, {-4, 6}, {-8, 10}, {-10, 12}, {-11, 13}, {-12, 14}, {-12, 14}, {-12, 14}, {-11, 13}, {-10, 12}, {-8, 10}, {-4, 6}, {0, 2}}, .4},
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
