package main

import (
	"strconv"
)

type ARMinstruction struct {
	Mnemonic string
	Operand1 string
	Operand2 string
	Operand3 string
}

func generalizeARMinstruction(arm ARMinstruction) ARMinstruction {

	operands := []string{arm.Operand1, arm.Operand2, arm.Operand3}
	operands_int := make(map[string][]int)
	var int_operand [3]int

	// associate operand to indices
	for index, operand := range operands {
		if operand != "" {
			operands_int[operand] = append(operands_int[operand], index)
		}
	}

	// reverse map to an array
	var actual int = 0
	for _, index := range operands_int {
		for _, i := range index {
			int_operand[i] = actual
		}
		actual++
	}

	// TODO : fix when Operand3 is not set (default 0 :/ )
	generalizedInstruction := ARMinstruction{
		Mnemonic: arm.Mnemonic,
		Operand1: "$" + strconv.Itoa(int_operand[0]),
		Operand2: "$" + strconv.Itoa(int_operand[1]),
		Operand3: "$" + strconv.Itoa(int_operand[2]),
	}

	return generalizedInstruction

}

/*
func generatePolymorph(ARMinstruction) {
	//ARMinstruction.
}*/
