package main

import (
	"math/rand"
	"strconv"
)

type ARMinstruction struct {
	Mnemonic string
	Operands []string
}

func generalizeARMinstruction(arm ARMinstruction) *ARMinstruction {

	operands := arm.Operands
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
	if len(arm.Operands) == 3 {
		generalizedInstruction := ARMinstruction{
			Mnemonic: arm.Mnemonic,
			Operands: []string{"$r" + strconv.Itoa(int_operand[0]),
				"$r" + strconv.Itoa(int_operand[1]),
				"$r" + strconv.Itoa(int_operand[2]),
			},
		}
		return &generalizedInstruction
	}
	if len(arm.Operands) == 2 {
		generalizedInstruction := ARMinstruction{
			Mnemonic: arm.Mnemonic,
			Operands: []string{"$r" + strconv.Itoa(int_operand[0]),
				"$r" + strconv.Itoa(int_operand[1]),
			},
		}
		return &generalizedInstruction
	}

	return nil

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
		arm.Operands[0] + ", " + arm.Operands[1]

	if len(arm.Operands) == 3 {
		str_equi = str_equi + ", " + arm.Operands[2]
	}

	if contains(equivalence, str_equi) {
		return equivalence[str_equi][rand.Intn(len(equivalence[str_equi]))]

	}
	return ""
}
