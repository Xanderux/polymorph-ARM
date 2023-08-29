package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	instr1 := ARMinstruction{
		Mnemonic: "movs",
		Operand1: "r7",
		Operand2: "#1",
		Operand3: "",
	}

	fmt.Println("Mnemonic:", instr1.Mnemonic, instr1.Operand1, instr1.Operand2, instr1.Operand3)
	var instr2 ARMinstruction = generalizeARMinstruction(instr1)

	fmt.Println("Mnemonic:", instr2.Mnemonic, instr2.Operand1, instr2.Operand2, instr2.Operand3)

}
