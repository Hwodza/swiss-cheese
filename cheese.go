package main

import "fmt"

func main() {
	var circle_options = [...]string{
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
}
