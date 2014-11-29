package main

import (
	"fmt"
	"testing"

	"github.com/kyokomi/gomajan/mjp"
)

type TestCase struct {
	in  []Tehai
	out []string
}

func TestYakuCheck(t *testing.T) {
	testCases := []TestCase{
		TestCase{
			in: NewTehai(map[mjp.MJP]int{
				mjp.S1:  2,
				mjp.S9:  1,
				mjp.P1:  1,
				mjp.P9:  1,
				mjp.M1:  1,
				mjp.M9:  1,
				mjp.TON: 1,
				mjp.NAN: 1,
				mjp.SHA: 1,
				mjp.PEI: 1,
				mjp.HAK: 1,
				mjp.HAT: 1,
				mjp.CHN: 1,
			}),
			out: []string{"国士無双"},
		},
		TestCase{
			in: NewTehai(map[mjp.MJP]int{
				mjp.S1:  2,
				mjp.S2:  2,
				mjp.M3:  2,
				mjp.M4:  2,
				mjp.P5:  2,
				mjp.P6:  2,
				mjp.TON: 2,
			}),
			out: []string{"七対子"},
		},
		// ソーズ清一色
		TestCase{
			in: NewTehai(map[mjp.MJP]int{
				mjp.S1: 2,
				mjp.S2: 2,
				mjp.S3: 2,
				mjp.S4: 1,
				mjp.S5: 1,
				mjp.S6: 1,
				mjp.S7: 3,
				mjp.S9: 2,
			}),
			out: []string{"清一色"},
		},
		// マンズ清一色
		TestCase{
			in: NewTehai(map[mjp.MJP]int{
				mjp.M1: 2,
				mjp.M2: 2,
				mjp.M3: 2,
				mjp.M4: 1,
				mjp.M5: 1,
				mjp.M6: 1,
				mjp.M7: 3,
				mjp.M9: 2,
			}),
			out: []string{"清一色"},
		},
		// ピンズ清一色
		TestCase{
			in: NewTehai(map[mjp.MJP]int{
				mjp.P1: 2,
				mjp.P2: 2,
				mjp.P3: 2,
				mjp.P4: 1,
				mjp.P5: 1,
				mjp.P6: 1,
				mjp.P7: 3,
				mjp.P9: 2,
			}),
			out: []string{"清一色"},
		},
		// ソーズ混一色
		TestCase{
			in: NewTehai(map[mjp.MJP]int{
				mjp.S1: 2,
				mjp.S2: 2,
				mjp.S3: 2,
				mjp.S4: 1,
				mjp.S5: 1,
				mjp.S6: 1,
				mjp.S7: 3,
				mjp.HAK: 2,
			}),
			out: []string{"混一色"},
		},
		// マンズ混一色
		TestCase{
			in: NewTehai(map[mjp.MJP]int{
				mjp.M1: 2,
				mjp.M2: 2,
				mjp.M3: 2,
				mjp.M4: 1,
				mjp.M5: 1,
				mjp.M6: 1,
				mjp.M7: 3,
				mjp.TON: 2,
			}),
			out: []string{"混一色"},
		},
		// ピンズ混一色
		TestCase{
			in: NewTehai(map[mjp.MJP]int{
				mjp.P1: 2,
				mjp.P2: 2,
				mjp.P3: 2,
				mjp.P4: 1,
				mjp.P5: 1,
				mjp.P6: 1,
				mjp.P7: 3,
				mjp.CHN: 2,
			}),
			out: []string{"混一色"},
		},
		// 断么九
		TestCase{
			in: NewTehai(map[mjp.MJP]int{
				mjp.P2: 1,
				mjp.P3: 1,
				mjp.P4: 1,
				mjp.S4: 1,
				mjp.S5: 1,
				mjp.S6: 1,
				mjp.M7: 3,
				mjp.P8: 2,
			}),
			out: []string{"断么九"},
		},
		// 緑一色
		TestCase{
			in: NewTehai(map[mjp.MJP]int{
				mjp.S2: 2,
				mjp.S3: 2,
				mjp.S4: 2,
				mjp.S6: 3,
				mjp.HAT: 3,
				mjp.S8: 2,
			}),
			out: []string{"緑一色"},
		},
		// 大三元
		TestCase{
			in: NewTehai(map[mjp.MJP]int{
				mjp.S2: 1,
				mjp.S3: 1,
				mjp.S4: 1,
				mjp.P5: 2,
				mjp.HAK: 3,
				mjp.HAT: 3,
				mjp.CHN: 3,
			}),
			out: []string{"大三元"},
		},
		// 字一色
		TestCase{
			in: NewTehai(map[mjp.MJP]int{
				mjp.TON: 2,
				mjp.NAN: 2,
				mjp.SHA: 2,
				mjp.PEI: 2,
				mjp.HAK: 2,
				mjp.HAT: 2,
				mjp.CHN: 2,
			}),
			out: []string{"字一色"},
		},
		// TODO: 副露考慮したら、あとで四暗刻なくす
		// 大四喜
		TestCase{
			in: NewTehai(map[mjp.MJP]int{
				mjp.TON: 3,
				mjp.NAN: 3,
				mjp.SHA: 3,
				mjp.PEI: 3,
				mjp.P1: 2,
			}),
			out: []string{"大四喜", "四暗刻"},
		},
		// 小四喜
		TestCase{
			in: NewTehai(map[mjp.MJP]int{
				mjp.TON: 3,
				mjp.NAN: 3,
				mjp.SHA: 3,
				mjp.PEI: 2,
				mjp.P1: 1,
				mjp.P2: 1,
				mjp.P3: 1,
			}),
			out: []string{"小四喜"},
		},
		// 四暗刻
		TestCase{
			in: NewTehai(map[mjp.MJP]int{
				mjp.TON: 3,
				mjp.S1: 3,
				mjp.S5: 3,
				mjp.P9: 3,
				mjp.CHN: 2,
			}),
			out: []string{"四暗刻"},
		},
		// 三暗刻
		TestCase{
			in: NewTehai(map[mjp.MJP]int{
				mjp.TON: 3,
				mjp.S1: 3,
				mjp.S5: 3,
				mjp.P7: 1,
				mjp.P8: 1,
				mjp.P9: 1,
				mjp.CHN: 2,
			}),
			out: []string{"三暗刻"},
		},
		// TODO: 副露考慮したら、あとで四暗刻なくす
		// 清老頭
		TestCase{
			in: NewTehai(map[mjp.MJP]int{
				mjp.S1: 2,
				mjp.P1: 3,
				mjp.P9: 3,
				mjp.M1: 3,
				mjp.M9: 3,
			}),
			out: []string{"清老頭", "四暗刻"},
		},
		// 小三元
		TestCase{
			in: NewTehai(map[mjp.MJP]int{
				mjp.S2: 1,
				mjp.S3: 1,
				mjp.S4: 1,
				mjp.P5: 1,
				mjp.P6: 1,
				mjp.P7: 1,
				mjp.HAK: 3,
				mjp.HAT: 3,
				mjp.CHN: 2,
			}),
			out: []string{"小三元"},
		},
	}

	for _, testCase := range testCases {
		p := NewPlayer()
		p.tiles = testCase.in

		if yaku := yakuCheck(p); len(yaku) != len(testCase.out) {
			t.Error(testCase.out, "!= ", yaku, " error 手牌 => ", p)
		} else {
			fmt.Println(" => ", yaku)
		}
	}
}
