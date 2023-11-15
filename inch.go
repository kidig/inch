package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	flag.Parse()

	if inch := flag.Arg(0); inch != "" {
		if inch, err := strconv.ParseFloat(inch, 32); err == nil {
			fmt.Printf("%s cm\n", strconv.FormatFloat(inch*2.5, 'f', -1, 32))
		} else {
			fmt.Println("Wrong value")
		}
	}
}
