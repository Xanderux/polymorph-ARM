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
	instr1 := sources.ARMinstruction{
		Mnemonic: "subs",
		Operands: []string{
			"r4", "r4", "r4",
		},
	}

	displayARMinstruction("Original:", instr1)
	var instr2 = sources.GeneralizeARMinstruction(instr1)
	displayARMinstruction("Generalize:", *instr2)
	var instr3 string = sources.GeneratePolymorph(*instr2)
	fmt.Println("Result:", instr3)

	sources.PolymorphEngine("shellcode-904.c", "shellcode-904_new.c")

}
