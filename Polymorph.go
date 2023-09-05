package main

import (
	"math/rand"
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
	int_operand := [3]int{-1, -1, -1}

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
		Operand1: "$r" + strconv.Itoa(int_operand[0]),
		Operand2: "$r" + strconv.Itoa(int_operand[1]),
		Operand3: "$r" + strconv.Itoa(int_operand[2]),
	}

	return generalizedInstruction

}

func contains(slice map[string][]string, value string) bool {
	for src, _ := range slice {
		if src == value {
			return true
		}

	}
	return false

}

func generatePolymorph(arm ARMinstruction) string {
	equivalence := map[string][]string{
		"subs $r0, $r0, $r0": {
			"subs r4, r4, r4",
			"mov r4, #0",
			"eor r4, r4, r4",
			"bic r4, r4, r4",
			"and r4, r4, #0",
		},
	}
	var str_equi = arm.Mnemonic + " " +
		arm.Operand1 + ", " + arm.Operand2 + ", " + arm.Operand3

	if contains(equivalence, str_equi) {
		return equivalence[str_equi][rand.Intn(len(equivalence[str_equi]))]

	}
	return ""
}
