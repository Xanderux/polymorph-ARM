package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Xanderux/polymorph-ARM/sources"
)

func displayARMinstruction(message string, arm sources.ARMinstruction) {
	fmt.Print(message, " ", arm.Mnemonic, " ", arm.Operands[0], " ", arm.Operands[1])
	if len(arm.Operands) == 3 {
		fmt.Print(" ", arm.Operands[2])
	}
	fmt.Println()
}

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
