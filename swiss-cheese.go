package main

import (
	"bufio"
	"fmt"
	"os"
)

type circle struct {
	name       string
	num_rows   int
	num_cols   int
	char_avail []int
	edges      [][]string
}

func main() {
	var c circle
	fmt.Println(c)
	circle_options := [...]string{
		`
                   ### *** *** ###
               *##                 ##*
           *##                         ##*
        *##                               ##*
      *##                                   ##*
    *##                                       ##*
   *##                                         ##*
  *##                                           ##*
 *##                                             ##*
 *##                                             ##*
 *##                                             ##*
 *##                                             ##*
 *##                                             ##*
  *##                                           ##*
   *##                                         ##*
    *##                                       ##*
      *##                                   ##*
        *#                                ##*
           *##                         ##*
               *##                 ##*
                   *** ### ### *** `,
		`
 ,#####, 
##*   '##
#(     )#
##,   ,##
 "#####" 

		`,
		`
                      ,,ggddY""""Ybbgg,,
                 ,agd""'              '""bg,
              ,gdP"                       "Ybg,
            ,dP"                             "Yb,
          ,dP"         _,,ddP"""Ybb,,_         "Yb,
         ,8"         ,dP"'         '"Yb,         "8,
        ,8'        ,d"                 "b,        '8,
       ,8'        d"                     "b        '8,
       d'        d'        ,gPPRg,        'b        'b
       8         8        dP'   'Yb        8         8
       8         8        8)     (8        8         8
       8         8        Yb     dP        8         8
       8         Y,        "8ggg8"        ,P         8
       Y,         Ya                     aP         ,P
       '8,         "Ya                 aP"         ,8'
        '8,          "Yb,_         _,dP"          ,8'
         '8a           '""YbbgggddP""'           a8'
          'Yba                                 adP'
            "Yba                             adY"
              '"Yba,                     ,adP"'
                 '"Y8ba,             ,ad8P"'
                      ''""YYbaaadPP""''`,
	}
	for i := 0; i < len(circle_options); i++ {
		fmt.Println(circle_options[i])
	}
	fmt.Println("Hello, World!")

	// Reading in standard in
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanBytes)

	var input []byte
	for scanner.Scan() {
		input = append(input, scanner.Bytes()...)
	}

	var lines []string
	const lineLength = 80

	for i := 0; i < len(input); i += lineLength {
		end := i + lineLength
		if end > len(input) {
			end = len(input)
		}
		lines = append(lines, string(input[i:end]))
	}
	for i, line := range lines {
		fmt.Printf("line %d, %q\n", i, line)
	}
}
