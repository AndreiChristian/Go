package main

import (
	"fmt"
	"os"
)

func CliArgs() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

}

// another variant

func CliArgs2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = ""
	}
	fmt.Println(s)
}

// for confdtion{

// }
// this is the equivalent of a while loop
