package main

import "fmt"

func main() {
	instr1 := ARMinstruction{
		Mnemonic: "subs",
		Operands: []string{
			"r4", "r4", "r4",
		},
	}

	fmt.Println("Original:", instr1.Mnemonic, instr1.Operands[0], instr1.Operands[1], instr1.Operands[2])
	var instr2 ARMinstruction = generalizeARMinstruction(instr1)
	fmt.Println("Generalize:", instr2.Mnemonic, instr2.Operands[0], instr2.Operands[1], instr2.Operands[2])
	var instr3 string = generatePolymorph(instr2)
	fmt.Println("Result:", instr3)

}
