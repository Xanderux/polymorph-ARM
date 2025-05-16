package main

import (
	"fmt"

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
	sources.PolymorphEngine("shellcode-904.c", "shellcode-904_new.c")
}
