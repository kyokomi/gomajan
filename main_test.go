package main

import (
	"testing"

	"fmt"

	"github.com/kyokomi/gomajan/mjp"
	"github.com/kyokomi/gomajan/mjp/foo"
	"github.com/kyokomi/gomajan/mjp/pai"
	"github.com/kyokomi/gomajan/mjp/tehai"
)

type TestCase struct {
	in     []tehai.Tehai
	inFoos []foo.Foo
	out    []string
}

func TestYakuCheck(t *testing.T) {
	testCases := []TestCase{
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.S1:  2,
				pai.S9:  1,
				pai.P1:  1,
				pai.P9:  1,
				pai.M1:  1,
				pai.M9:  1,
				pai.TON: 1,
				pai.NAN: 1,
				pai.SHA: 1,
				pai.PEI: 1,
				pai.HAK: 1,
				pai.HAT: 1,
				pai.CHN: 1,
			}),
			inFoos: nil,
			out:    []string{"国士無双"},
		},
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.S1:  2,
				pai.S2:  2,
				pai.M3:  2,
				pai.M4:  2,
				pai.P5:  2,
				pai.P6:  2,
				pai.TON: 2,
			}),
			inFoos: nil,
			out:    []string{"七対子"},
		},
		// ソーズ清一色
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.S1: 1,
				pai.S2: 1,
				pai.S3: 2,
				pai.S4: 1,
				pai.S5: 1,
				pai.S6: 3,
				pai.S7: 3,
				pai.S9: 2,
			}),
			inFoos: nil,
			out:    []string{"清一色"},
		},
		// マンズ清一色
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.M1: 1,
				pai.M2: 2,
				pai.M3: 2,
				pai.M4: 2,
				pai.M5: 1,
				pai.M6: 1,
				pai.M7: 3,
				pai.M9: 2,
			}),
			inFoos: nil,
			out:    []string{"清一色"},
		},
		// ピンズ清一色
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.P1: 1,
				pai.P2: 2,
				pai.P3: 2,
				pai.P4: 2,
				pai.P5: 1,
				pai.P6: 1,
				pai.P7: 3,
				pai.P9: 2,
			}),
			inFoos: nil,
			out:    []string{"清一色"},
		},
		// ソーズ混一色
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.S1:  1,
				pai.S2:  2,
				pai.S3:  2,
				pai.S4:  2,
				pai.S5:  1,
				pai.S6:  1,
				pai.S7:  3,
				pai.HAK: 2,
			}),
			inFoos: nil,
			out:    []string{"混一色"},
		},
		// マンズ混一色
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.M1:  1,
				pai.M2:  2,
				pai.M3:  2,
				pai.M4:  2,
				pai.M5:  1,
				pai.M6:  1,
				pai.M7:  3,
				pai.TON: 2,
			}),
			inFoos: nil,
			out:    []string{"混一色"},
		},
		// ピンズ混一色
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.P1:  1,
				pai.P2:  2,
				pai.P3:  2,
				pai.P4:  2,
				pai.P5:  1,
				pai.P6:  1,
				pai.P7:  3,
				pai.CHN: 2,
			}),
			inFoos: nil,
			out:    []string{"混一色"},
		},
		// 断么九
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.P2: 1,
				pai.P3: 1,
				pai.P4: 1,
				pai.S4: 1,
				pai.S5: 1,
				pai.S6: 1,
				pai.M7: 3,
				pai.S7: 3,
				pai.P8: 2,
			}),
			inFoos: nil,
			out:    []string{"断么九"},
		},
		// 断么九（クイタン）
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.P2: 1,
				pai.P3: 1,
				pai.P4: 1,
				pai.S4: 1,
				pai.S5: 1,
				pai.S6: 1,
				pai.P8: 2,
			}),
			inFoos: []foo.Foo{
				foo.NewFooPon(foo.Toimen, pai.M8),
				foo.NewFooPon(foo.SimoCha, pai.M2),
			},
			out: []string{"断么九"},
		},
		// 緑一色
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.S2: 2,
				pai.S3: 2,
				pai.S4: 2,
				pai.S6: 3,
				pai.S8: 2,
			}),
			inFoos: []foo.Foo{
				foo.NewFooPon(foo.KamiCha, pai.HAT),
			},
			out: []string{"緑一色"},
		},
		// 大三元
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.S2:  1,
				pai.S3:  1,
				pai.S4:  1,
				pai.P5:  2,
				pai.HAT: 3,
				pai.CHN: 3,
			}),
			inFoos: []foo.Foo{foo.NewFooPon(foo.Toimen, pai.HAK)},
			out:    []string{"大三元"},
		},
		// 字一色
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.TON: 3,
				pai.NAN: 3,
				pai.HAT: 3,
				pai.CHN: 2,
			}),
			inFoos: []foo.Foo{foo.NewFooPon(foo.Toimen, pai.HAK)},
			out:    []string{"字一色"},
		},
		// 大四喜
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.TON: 3,
				pai.NAN: 3,
				pai.PEI: 3,
				pai.P1:  2,
			}),
			inFoos: []foo.Foo{foo.NewFooPon(foo.Toimen, pai.SHA)},
			out:    []string{"大四喜"},
		},
		// 小四喜
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.TON: 3,
				pai.NAN: 3,
				pai.PEI: 2,
				pai.P1:  1,
				pai.P2:  1,
				pai.P3:  1,
			}),
			inFoos: []foo.Foo{foo.NewFooPon(foo.Toimen, pai.SHA)},
			out:    []string{"小四喜"},
		},
		// 四暗刻
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.TON: 3,
				pai.S1:  3,
				pai.S5:  3,
				pai.CHN: 2,
			}),
			inFoos: []foo.Foo{
				foo.NewFooAnnKan(pai.P9),
			},
			out: []string{"四暗刻"},
		},
		// 三暗刻
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.TON: 3,
				pai.S1:  3,
				pai.P7:  1,
				pai.P8:  1,
				pai.P9:  1,
				pai.CHN: 2,
			}),
			inFoos: []foo.Foo{
				foo.NewFooAnnKan(pai.S5),
			},
			out: []string{"三暗刻"},
		},
		// 清老頭
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.S1: 2,
				pai.P1: 3,
				pai.P9: 3,
				pai.S9: 3,
			}),
			inFoos: []foo.Foo{foo.NewFooPon(foo.Toimen, pai.M1)},
			out:    []string{"清老頭"},
		},
		// 混老頭
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.S1:  2,
				pai.P1:  3,
				pai.HAK: 3,
			}),
			inFoos: []foo.Foo{
				foo.NewFooPon(foo.Toimen, pai.M1),
				foo.NewFooPon(foo.Toimen, pai.TON),
			},
			out: []string{"対々和", "混老頭"},
		},
		// 小三元
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.S2:  1,
				pai.S3:  1,
				pai.S4:  1,
				pai.P5:  1,
				pai.P6:  1,
				pai.P7:  1,
				pai.HAT: 3,
				pai.CHN: 2,
			}),
			inFoos: []foo.Foo{
				foo.NewFooPon(foo.KamiCha, pai.HAK),
			},
			out: []string{"小三元"},
		},
		// 四槓子
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.CHN: 2,
			}),
			inFoos: []foo.Foo{
				foo.NewFooMinKan(foo.Toimen, pai.M1),
				foo.NewFooAnnKan(pai.M2),
				foo.NewFooAnnKan(pai.M3),
				foo.NewFooAnnKan(pai.M4),
			},
			out: []string{"四槓子"},
		},
		// 三槓子
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.CHN: 2,
				pai.P1:  1,
				pai.P2:  1,
				pai.P3:  1,
			}),
			inFoos: []foo.Foo{
				foo.NewFooMinKan(foo.Toimen, pai.M1),
				foo.NewFooAnnKan(pai.M3),
				foo.NewFooAnnKan(pai.M4),
			},
			out: []string{"三槓子"},
		},
		// 一気通貫
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.P4: 1,
				pai.P5: 1,
				pai.P6: 1,
				pai.P7: 1,
				pai.P8: 1,
				pai.P9: 3,
				pai.S2: 3,
			}),
			inFoos: []foo.Foo{
				foo.NewFooChe(pai.P1, pai.P2, pai.P3),
			},
			out: []string{"一気通貫"},
		},
		// 二盃口
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.P4: 2,
				pai.P5: 2,
				pai.P6: 2,
				pai.S7: 2,
				pai.S8: 2,
				pai.S9: 2,
				pai.S1: 2,
			}),
			inFoos: nil,
			out:    []string{"二盃口"},
		},
		// 一盃口
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.P4: 2,
				pai.P5: 2,
				pai.P6: 2,
				pai.S7: 1,
				pai.S8: 1,
				pai.S9: 1,
				pai.M3: 3,
				pai.S2: 2,
			}),
			inFoos: nil,
			out:    []string{"一盃口"},
		},
		// 純全帯么九
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.P1: 1,
				pai.P2: 1,
				pai.P3: 1,
				pai.S7: 1,
				pai.S8: 1,
				pai.S9: 1,
				pai.M9: 3,
				pai.S1: 2,
			}),
			inFoos: []foo.Foo{
				foo.NewFooChe(pai.P7, pai.P8, pai.P9),
			},
			out:    []string{"純全帯么九"},
		},
		// 混全帯么九
		TestCase{
			in: tehai.NewTehai(map[pai.MJP]int{
				pai.P1: 1,
				pai.P2: 1,
				pai.P3: 1,
				pai.S7: 1,
				pai.S8: 1,
				pai.S9: 1,
				pai.HAK: 3,
				pai.S1: 2,
			}),
			inFoos: []foo.Foo{
				foo.NewFooChe(pai.P7, pai.P8, pai.P9),
			},
			out:    []string{"混全帯么九"},
		},
	}

	for _, testCase := range testCases {
		p := mjp.NewPlayer(testCase.in, testCase.inFoos)

		yakuCheck := p.NewYakuCheck()

		// 役の数
		if len(yakuCheck.Yakus()) != len(testCase.out) {
			t.Error(testCase.out, "!= ", yakuCheck.Yakus(), " 役数error 手牌 => ", p)
			continue
		}

		// 役の内容（順番もチェックしてる）
		ok := true
		for idx, y := range yakuCheck.Yakus() {
			if y.Name != testCase.out[idx] {
				t.Error(testCase.out, "!= ", yakuCheck.Yakus(), " 役内容error 手牌 => ", p)
				ok = false
				break
			}
		}
		if ok {
			fmt.Println(" => ", yakuCheck.String())
		}
	}
}
