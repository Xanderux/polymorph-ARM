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

	// Maps to assign a unique ID to each register and each immediate
	regMap := make(map[string]int)
	immMap := make(map[string]int)
	nextRegID, nextImmID := 0, 0

	// Slice to store the assigned IDs per operand position
	intOperand := make([]int, len(operands))

	// Iterate through the operands in order
	for i, op := range operands {
		if strings.HasPrefix(op, "#") {
			// Immediate operand
			if _, seen := immMap[op]; !seen {
				immMap[op] = nextImmID
				nextImmID++
			}
			intOperand[i] = immMap[op]
		} else {
			// Register operand (or treated as register if not immediate)
			if _, seen := regMap[op]; !seen {
				regMap[op] = nextRegID
				nextRegID++
			}
			intOperand[i] = regMap[op]
		}
	}

	// Build the generalized operands slice
	generalized := make([]string, len(operands))
	for i, op := range operands {
		if strings.HasPrefix(op, "#") {
			generalized[i] = "$imm" + strconv.Itoa(intOperand[i])
		} else {
			generalized[i] = "$r" + strconv.Itoa(intOperand[i])
		}
	}

	// Return the generalized instruction
	return &ARMinstruction{
		Mnemonic: arm.Mnemonic,
		Operands: generalized,
	}
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
	// TODO 1 : change struct into [][]string
	// TODO 2 : one instruction can be equivalent to 2 others
	equivalence := map[string][]string{
		"SUBS $r0 $r0 $r0": {
			"MOVS $r0 #0",
			"EORS $r0 $r0 $r0",
			"ANDS $r0 $r0 #0",
			"BICS $r0 $r0 $r0",
		},
	}
	str_equi := arm.Mnemonic
	for _, op := range arm.Operands {
		str_equi = str_equi + " " + op
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

func PolymorphToInstruction(polyStr string, baseIns ARMinstruction) string {
	polyIns := stringToARMinstruction(polyStr)

	results := make([]string, 0)

	for _, op := range polyIns.Operands {
		if strings.HasPrefix(op, "$r") {
			// extract index
			indexStr := strings.TrimPrefix(op, "$r")
			index, err := strconv.Atoi(indexStr)
			if err != nil || index < 0 || index >= len(baseIns.Operands) {
				// failed
				return ""
			}
			results = append(results, baseIns.Operands[index])
		} else {
			// immediate value
			results = append(results, op)
		}
	}

	newIns := ARMinstruction{
		Mnemonic: polyIns.Mnemonic,
		Operands: results,
	}

	return ARMinstructionToString(newIns)
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
			fmt.Println(result)
			// fetch the base instruction
			base_ins := stringToARMinstruction(result)
			// generalize it
			gen_ins := GeneralizeARMinstruction(base_ins)
			poly_ins := GeneratePolymorph(*gen_ins)
			new_ins := PolymorphToInstruction(poly_ins, base_ins)
			if new_ins != "" {
				file.WriteString(new_ins + "\n")
			} else {
				file.WriteString(result + "\n")
			}

		}
	}
}
