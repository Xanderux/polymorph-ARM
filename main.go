package main

import (
	"fmt"
	"strings"
)

func displayARMinstruction(message string, arm ARMinstruction) {
	fmt.Print(message, " ", arm.Mnemonic, " ", arm.Operands[0], " ", arm.Operands[1])
	if len(arm.Operands) == 3 {
		fmt.Print(" ", arm.Operands[2])
	}
	fmt.Println()
}

func main() {
	instr1 := ARMinstruction{
		Mnemonic: "subs",
		Operands: []string{
			"r4", "r4", "r4",
		},
	}

	displayARMinstruction("Original:", instr1)
	var instr2 = generalizeARMinstruction(instr1)
	displayARMinstruction("Generalize:", *instr2)
	var instr3 string = generatePolymorph(*instr2)
	fmt.Println("Result:", instr3)

	content := readLineByLine("shellcode-904.c")

	for _, str := range content {
		if isARMInstruction(strings.ToUpper(str)) {
			fmt.Println("ARM detected :" + str)
		}
	}

}
