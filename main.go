package main

import "fmt"

func main() {
	instr1 := ARMinstruction{
		Mnemonic: "subs",
		Operand1: "r4",
		Operand2: "r4",
		Operand3: "r4",
	}

	fmt.Println("Original:", instr1.Mnemonic, instr1.Operand1, instr1.Operand2, instr1.Operand3)
	var instr2 ARMinstruction = generalizeARMinstruction(instr1)
	fmt.Println("Generalize:", instr2.Mnemonic, instr2.Operand1, instr2.Operand2, instr2.Operand3)
	var instr3 string = generatePolymorph(instr2)
	fmt.Println("Result:", instr3)

}
