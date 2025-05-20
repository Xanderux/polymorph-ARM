package tests

import (
	"testing"

	"github.com/Xanderux/polymorph-ARM/sources"
)

func TestIsARMInstruction(t *testing.T) {
	armIns := []string{
		"subs r1, r1, r1",
		"movs r7, #8",
		"eors r0, r0, r0",
		"ands r0, r0, #0",
		"bics r0, r0, r0",
		"adds r7, #3",
	}
	nonArmIns := []string{
		"my string",
		"movv r0, r4",
	}

	for _, ins := range armIns {
		result := sources.IsARMInstruction(ins)
		if result == "" {
			t.Errorf("Should be detected as ARM instruction: %s", ins)
		}
	}

	for _, ins := range nonArmIns {
		result := sources.IsARMInstruction(ins)
		if result != "" {
			t.Errorf("Should NOT be detected as ARM instruction: %s", ins)
		}
	}
}

func TestGeneralizeARMinstruction(t *testing.T) {
	tests := []struct {
		input    sources.ARMinstruction
		expected sources.ARMinstruction
	}{
		{
			input: sources.ARMinstruction{
				Mnemonic: "subs",
				Operands: []string{"r1", "r1", "r1"},
			},
			expected: sources.ARMinstruction{
				Mnemonic: "subs",
				Operands: []string{"$r0", "$r0", "$r0"},
			},
		},
		{
			input: sources.ARMinstruction{
				Mnemonic: "subs",
				Operands: []string{"r2", "r3", "r2"},
			},
			expected: sources.ARMinstruction{
				Mnemonic: "subs",
				Operands: []string{"$r0", "$r1", "$r0"},
			},
		},
		{
			input: sources.ARMinstruction{
				Mnemonic: "movs",
				Operands: []string{"r7", "#8"},
			},
			expected: sources.ARMinstruction{
				Mnemonic: "movs",
				Operands: []string{"$r0", "$imm0"},
			},
		},
	}

	for i, test := range tests {
		result := sources.GeneralizeARMinstruction(test.input)
		if result == nil {
			t.Fatalf("Test %d: GeneralizeARMinstruction returned nil", i)
		}

		if result.Mnemonic != test.expected.Mnemonic {
			t.Errorf("Test %d: Expected mnemonic %s, got %s", i, test.expected.Mnemonic, result.Mnemonic)
		}

		for j := range test.expected.Operands {
			if result.Operands[j] != test.expected.Operands[j] {
				t.Errorf("Test %d: Expected operand %d to be %s, got %s", i, j, test.expected.Operands[j], result.Operands[j])
			}
		}
	}
}
