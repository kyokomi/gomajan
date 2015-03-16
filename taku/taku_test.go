package taku

import (
	"testing"

	"fmt"

	"github.com/kyokomi/gomajan/pai"
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
	fmt.Println("サイコロ:", taku.Sai)
	fmt.Println("ドラ表示:", taku.Dora())
	fmt.Println("裏ドラ表示:", taku.UraDora())

	yamaIdx := 山Index(taku.Sai)
	retu := taku.Sai.Sum()-1

	// TODO: 4人分配牌してみる

	var p [4][14]pai.MJP
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			calRetu := retu+1
			if calRetu > 16 {
				calRetu = 0
				yamaIdx--
				if yamaIdx < 0 {
					yamaIdx = 3
				}
			}

			p[j][(i*4)+0] = taku.Yama[yamaIdx][0][calRetu]
			p[j][(i*4)+1] = taku.Yama[yamaIdx][1][calRetu]

			calRetu += 1
			if calRetu > 16 {
				calRetu = 0
				yamaIdx--
				if yamaIdx < 0 {
					yamaIdx = 3
				}
			}

			p[j][(i*4)+2] = taku.Yama[yamaIdx][0][calRetu]
			p[j][(i*4)+3] = taku.Yama[yamaIdx][1][calRetu]

			retu = calRetu
		}
	}

	// ちょんちょん
	p[0][12] = taku.Yama[yamaIdx][0][retu]
	p[0][13] = taku.Yama[yamaIdx][0][retu+2]
	// ちょん
	p[1][12] = taku.Yama[yamaIdx][1][retu]
	p[2][12] = taku.Yama[yamaIdx][0][retu+1]
	p[3][12] = taku.Yama[yamaIdx][1][retu+1]

	fmt.Println("player1 配牌:", p[0])
	fmt.Println("player2 配牌:", p[1])
	fmt.Println("player3 配牌:", p[2])
	fmt.Println("player4 配牌:", p[3])
}
