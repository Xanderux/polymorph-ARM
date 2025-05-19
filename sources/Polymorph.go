package sources

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ARMinstruction struct {
	Mnemonic string
	Operands []string
}

func GeneralizeARMinstruction(arm ARMinstruction) *ARMinstruction {

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

func GeneratePolymorph(arm ARMinstruction) string {
	equivalence := map[string][]string{
		"SUBS $r0 $r0 $r0": {
			"MOVS $r0 #0",
			"EORS $r0 $r0 $r0",
			"ANDS $r0 $r0 #0",
			"BICS $r0 $r0 $r0",
		},
	}
	var str_equi = arm.Mnemonic + " " +
		arm.Operands[0] + " " + arm.Operands[1]

	if len(arm.Operands) == 3 {
		str_equi = str_equi + " " + arm.Operands[2]
	}

	if contains(equivalence, str_equi) {
		return equivalence[str_equi][rand.Intn(len(equivalence[str_equi]))]

	}
	// fail, return the base polymorph
	return str_equi
}

func ARMinstructionToString(arm ARMinstruction) string {
	if len(arm.Operands) < 2 {
		return ""
	}
	result := arm.Mnemonic + " " + arm.Operands[0] + " " + arm.Operands[1]
	if len(arm.Operands) == 3 {
		result += " " + arm.Operands[2]
	}
	return result
}

func PolymorphToInstruction(poly_str string, base_ins ARMinstruction) string {
	poly_ins := stringToARMinstruction(poly_str)

	// to do : "subs $r0, $r0, $r0" et ands $r0, $r0, #0 ne devrait pas passer
	if len(poly_ins.Operands) == len(base_ins.Operands) {
		new_ins := ARMinstruction{
			Mnemonic: poly_ins.Mnemonic,
			Operands: base_ins.Operands,
		}
		return ARMinstructionToString(new_ins)
	}
	return ""
}

func IsARMInstruction(ins string) string {
	// SUB R4, R5, #4
	regex := `(?i)(mov(s)|add(s)|sub(s)|eor(s)|and(s)|orr|bic(s)|cmp|cmn|ldr|str|bx|bl|b|bne|beq|blx){1,4}((,)?\s(R(1[0-5]|[0-9])|#[0-F])){1,3}`
	re := regexp.MustCompile(regex)
	matches := re.FindString(ins)
	return matches
}

// Convert an assembly instruction into a struct
func stringToARMinstruction(src string) ARMinstruction {
	src = strings.ReplaceAll(src, ",", "")
	slice := strings.Split(src, " ")
	operand := slice[0]
	operands := slice[1:]
	instr1 := ARMinstruction{
		Mnemonic: operand,
		Operands: operands,
	}
	return instr1

}

func PolymorphEngine(inputPath string, outputPath string) {
	content := readLineByLine(inputPath)

	file, err := os.Create(outputPath)

	if err != nil {
		fmt.Println("Can't open file" + outputPath)

	}

	for _, str := range content {

		result := IsARMInstruction(strings.ToUpper(str))

		if result == "" {
			file.WriteString(str + "\n")
		} else {
			// fetch the base instruction
			base_ins := stringToARMinstruction(result)
			// generalize it
			gen_ins := GeneralizeARMinstruction(base_ins)
			poly_ins := GeneratePolymorph(*gen_ins)
			new_ins := PolymorphToInstruction(poly_ins, base_ins)

			file.WriteString(new_ins + "\n")
		}
	}
}
