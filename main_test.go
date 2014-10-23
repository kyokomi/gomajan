package main

import (
	"fmt"
	"testing"

	"github.com/kyokomi/gomajan/mjp"
)

type TestCase struct {
	in map[mjp.MJP]int
	out string
}

func TestYakuCheck(t *testing.T) {

	testCases := []TestCase{
		TestCase{
			in: map[mjp.MJP]int{
				mjp.S1 : 2,
				mjp.S9 : 1,
				mjp.P1 : 1,
				mjp.P9 : 1,
				mjp.M1 : 1,
				mjp.M9 : 1,
				mjp.TON: 1,
				mjp.NAN: 1,
				mjp.SHA: 1,
				mjp.PEI: 1,
				mjp.HAK: 1,
				mjp.HAT: 1,
				mjp.CHN: 1,
			},
			out: "国士無双",
		},
		TestCase{
			in: map[mjp.MJP]int{
				mjp.S1 : 2,
				mjp.S2 : 2,
				mjp.M3 : 2,
				mjp.M4 : 2,
				mjp.P5 : 2,
				mjp.P6 : 2,
				mjp.TON: 2,
			},
			out: "七対子",
		},
		// ソーズ清一色
		TestCase{
			in: map[mjp.MJP]int{
				mjp.S1 : 2,
				mjp.S2 : 2,
				mjp.S3 : 2,
				mjp.S4 : 1,
				mjp.S5 : 1,
				mjp.S6 : 1,
				mjp.S7 : 3,
				mjp.S9 : 2,
			},
			out: "清一色",
		},
		// マンズ清一色
		TestCase{
			in: map[mjp.MJP]int{
				mjp.M1 : 2,
				mjp.M2 : 2,
				mjp.M3 : 2,
				mjp.M4 : 1,
				mjp.M5 : 1,
				mjp.M6 : 1,
				mjp.M7 : 3,
				mjp.M9 : 2,
			},
			out: "清一色",
		},
		// ピンズ清一色
		TestCase{
			in: map[mjp.MJP]int{
				mjp.P1 : 2,
				mjp.P2 : 2,
				mjp.P3 : 2,
				mjp.P4 : 1,
				mjp.P5 : 1,
				mjp.P6 : 1,
				mjp.P7 : 3,
				mjp.P9 : 2,
			},
			out: "清一色",
		},
	}

	for _, testCase := range testCases {
		p := NewPlayer()
		p.tiles = testCase.in

		if yaku := yakuCheck(p); yaku != testCase.out {
			t.Error(testCase.out, " error 手牌 => ", p)
		} else {
			fmt.Println(" => ", yaku)
		}
	}
}
