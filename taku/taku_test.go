package taku

import (
	"testing"

	"fmt"

	"github.com/kyokomi/gomajan/mjp/pai"
)

func TestDoCalcPoint(t *testing.T) {
	testCase := []pai.MJP{
		pai.P2: 1,
		pai.P3: 1,
		pai.S7: 1,
		pai.S8: 1,
		pai.S9: 1,
		pai.M2: 1,
		pai.M3: 1,
		pai.M4: 1,
		pai.M5: 1,
		pai.M6: 1,
		pai.M7: 1,
		pai.S1: 2,
	}

	taku := NewTaku()
	for _, m := range testCase {
		taku.Players[1].PaiInc(m)
	}

	c := taku.RonCalcPoint(1, 2, pai.P1)
	fmt.Println(c.String())
}
