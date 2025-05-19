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
