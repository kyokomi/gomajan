package main

import (
	"testing"

	"fmt"
	"reflect"

	"github.com/kyokomi/gomajan/mjp"
)

type TestCase struct {
	in     []mjp.Tehai
	inFoos []mjp.Foo
	out    []mjp.Yaku
}

func TestYakuCheck(t *testing.T) {
	testCases := []TestCase{
		TestCase{
			in: mjp.NewTehai(map[mjp.MJP]int{
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
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M_国士無双},
		},
		TestCase{
			in: mjp.NewTehai(map[mjp.MJP]int{
				mjp.S1:  2,
				mjp.S2:  2,
				mjp.M3:  2,
				mjp.M4:  2,
				mjp.P5:  2,
				mjp.P6:  2,
				mjp.TON: 2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M_七対子},
		},
		// ソーズ清一色
		TestCase{
			in: mjp.NewTehai(map[mjp.MJP]int{
				mjp.S1: 2,
				mjp.S2: 2,
				mjp.S3: 2,
				mjp.S4: 1,
				mjp.S5: 1,
				mjp.S6: 1,
				mjp.S7: 3,
				mjp.S9: 2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M_清一色},
		},
		// マンズ清一色
		TestCase{
			in: mjp.NewTehai(map[mjp.MJP]int{
				mjp.M1: 2,
				mjp.M2: 2,
				mjp.M3: 2,
				mjp.M4: 1,
				mjp.M5: 1,
				mjp.M6: 1,
				mjp.M7: 3,
				mjp.M9: 2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M_清一色},
		},
		// ピンズ清一色
		TestCase{
			in: mjp.NewTehai(map[mjp.MJP]int{
				mjp.P1: 2,
				mjp.P2: 2,
				mjp.P3: 2,
				mjp.P4: 1,
				mjp.P5: 1,
				mjp.P6: 1,
				mjp.P7: 3,
				mjp.P9: 2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M_清一色},
		},
		// ソーズ混一色
		TestCase{
			in: mjp.NewTehai(map[mjp.MJP]int{
				mjp.S1:  2,
				mjp.S2:  2,
				mjp.S3:  2,
				mjp.S4:  1,
				mjp.S5:  1,
				mjp.S6:  1,
				mjp.S7:  3,
				mjp.HAK: 2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M_混一色},
		},
		// マンズ混一色
		TestCase{
			in: mjp.NewTehai(map[mjp.MJP]int{
				mjp.M1:  2,
				mjp.M2:  2,
				mjp.M3:  2,
				mjp.M4:  1,
				mjp.M5:  1,
				mjp.M6:  1,
				mjp.M7:  3,
				mjp.TON: 2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M_混一色},
		},
		// ピンズ混一色
		TestCase{
			in: mjp.NewTehai(map[mjp.MJP]int{
				mjp.P1:  2,
				mjp.P2:  2,
				mjp.P3:  2,
				mjp.P4:  1,
				mjp.P5:  1,
				mjp.P6:  1,
				mjp.P7:  3,
				mjp.CHN: 2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M_混一色},
		},
		// 断么九
		TestCase{
			in: mjp.NewTehai(map[mjp.MJP]int{
				mjp.P2: 1,
				mjp.P3: 1,
				mjp.P4: 1,
				mjp.S4: 1,
				mjp.S5: 1,
				mjp.S6: 1,
				mjp.M7: 3,
				mjp.S7: 3,
				mjp.P8: 2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M_断么九},
		},
		// 緑一色
		TestCase{
			in: mjp.NewTehai(map[mjp.MJP]int{
				mjp.S2:  2,
				mjp.S3:  2,
				mjp.S4:  2,
				mjp.S6:  3,
				mjp.HAT: 3,
				mjp.S8:  2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M_緑一色},
		},
		// 大三元
		TestCase{
			in: mjp.NewTehai(map[mjp.MJP]int{
				mjp.S2:  1,
				mjp.S3:  1,
				mjp.S4:  1,
				mjp.P5:  2,
				mjp.HAT: 3,
				mjp.CHN: 3,
			}),
			inFoos: []mjp.Foo{mjp.NewFooPon(mjp.HAK)},
			out:    []mjp.Yaku{mjp.M_大三元},
		},
		// 字一色
		TestCase{
			in: mjp.NewTehai(map[mjp.MJP]int{
				mjp.TON: 3,
				mjp.NAN: 3,
				mjp.HAT: 3,
				mjp.CHN: 2,
			}),
			inFoos: []mjp.Foo{mjp.NewFooPon(mjp.HAK)},
			out:    []mjp.Yaku{mjp.M_字一色},
		},
		// 大四喜
		TestCase{
			in: mjp.NewTehai(map[mjp.MJP]int{
				mjp.TON: 3,
				mjp.NAN: 3,
				mjp.PEI: 3,
				mjp.P1:  2,
			}),
			inFoos: []mjp.Foo{mjp.NewFooPon(mjp.SHA)},
			out:    []mjp.Yaku{mjp.M_大四喜},
		},
		// 小四喜
		TestCase{
			in: mjp.NewTehai(map[mjp.MJP]int{
				mjp.TON: 3,
				mjp.NAN: 3,
				mjp.PEI: 2,
				mjp.P1:  1,
				mjp.P2:  1,
				mjp.P3:  1,
			}),
			inFoos: []mjp.Foo{mjp.NewFooPon(mjp.SHA)},
			out:    []mjp.Yaku{mjp.M_小四喜},
		},
		// 四暗刻
		TestCase{
			in: mjp.NewTehai(map[mjp.MJP]int{
				mjp.TON: 3,
				mjp.S1:  3,
				mjp.S5:  3,
				mjp.P9:  3,
				mjp.CHN: 2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M_四暗刻},
		},
		// 三暗刻
		TestCase{
			in: mjp.NewTehai(map[mjp.MJP]int{
				mjp.TON: 3,
				mjp.S1:  3,
				mjp.S5:  3,
				mjp.P7:  1,
				mjp.P8:  1,
				mjp.P9:  1,
				mjp.CHN: 2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M_三暗刻},
		},
		// 清老頭
		TestCase{
			in: mjp.NewTehai(map[mjp.MJP]int{
				mjp.S1: 2,
				mjp.P1: 3,
				mjp.P9: 3,
				mjp.S9: 3,
			}),
			inFoos: []mjp.Foo{mjp.NewFooPon(mjp.M1)},
			out:    []mjp.Yaku{mjp.M_清老頭},
		},
		// 小三元
		TestCase{
			in: mjp.NewTehai(map[mjp.MJP]int{
				mjp.S2:  1,
				mjp.S3:  1,
				mjp.S4:  1,
				mjp.P5:  1,
				mjp.P6:  1,
				mjp.P7:  1,
				mjp.HAK: 3,
				mjp.HAT: 3,
				mjp.CHN: 2,
			}),
			inFoos: nil,
			out:    []mjp.Yaku{mjp.M_小三元},
		},
	}

	for _, testCase := range testCases {
		p := mjp.NewPlayer(testCase.in, testCase.inFoos)

		y := p.NewYakuCheck()

		if reflect.DeepEqual(y.Yakus(), testCase.out) {
			fmt.Println(" => ", y.String())
		} else {
			t.Error(testCase.out, "!= ", y.Yakus(), " error 手牌 => ", p)
		}
	}
}
