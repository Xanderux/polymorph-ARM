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
		"subs $r0, $r0, $r0": {
			"movs $r0, #0",
			"eors $r0, $r0, $r0",
			"ands $r0, $r0, #0",
			"bics $r0, $r0, $r0",
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

func isARMInstruction(ins string) string {
	// SUB R4, R5, #4
	regex := `(MOV|SUBS|ADDS){1,4}((,)?\s(R(1[0-5]|[0-9])|#[0-F])){1,3}`
	re := regexp.MustCompile(regex)
	matches := re.FindString(ins)
	return matches
}

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
		result := isARMInstruction(strings.ToUpper(str))
		if result == "" {
			file.WriteString(str + "\n")
		} else {
			ins_gen := GeneralizeARMinstruction(stringToARMinstruction(result))

			new_ins := GeneratePolymorph(*ins_gen)

			file.WriteString(new_ins + "\n")

		}
	}
}
