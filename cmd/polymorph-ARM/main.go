package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Xanderux/polymorph-ARM/sources"
)

func main() {
	input := flag.String("i", "", "ARM assembly source file")
	output := flag.String("o", "", "ARM assembly output file")

	flag.Parse()

	if *input == "" || *output == "" {
		fmt.Println("Invalid usage")
		flag.Usage()
		os.Exit(1)
	}

	sources.PolymorphEngine(*input, *output)

}
