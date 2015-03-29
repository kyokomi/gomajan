package taku

import (
	"testing"

	"fmt"

	"github.com/kyokomi/gomajan/pai"
)

func TestDoCalcPoint(t *testing.T) {
	taku := NewTaku()
	c := taku.RonCalcPoint(1, 2, pai.P1)
	fmt.Println(c.String())
	fmt.Println("サイコロ:", taku.Sai)
	fmt.Println("ドラ表示:", taku.Dora())
	fmt.Println("裏ドラ表示:", taku.UraDora())

	for _, p := range taku.Players {
		fmt.Printf("player%d 配牌: %s\n", p.PlayerID(), p.String())
	}

	ueshita := 2
	nextFunc := func(playerID int) pai.MJP {
		if ueshita == 2 {
			taku.Add列And山越しFunc(1)
			ueshita = 0
		}

		p := taku.Next(ueshita, playerID)
		ueshita++
		return p
	}

	for {
		p := nextFunc(1)
		if p == pai.NonePai {
			break
		}

		fmt.Println("引いた牌: ", p)

		taku.Players[1].PaiInc(p)
		yakuCheck := taku.Players[1].NewYakuCheck(p)
		if yakuCheck.Is和了() {
			fmt.Println("アガリ: ", yakuCheck.String())
			break
		}
		fmt.Println("チェック: ", yakuCheck.String())

		for _, noko := range yakuCheck.MentsuCheck().Nokori() {
			if noko.Pai.Type() == pai.NoneType || noko.Val == 0 {
				continue
			}

			fmt.Println("捨てた牌: ", noko.Pai)
			taku.Players[1].PaiDec(noko.Pai)
			break
		}
	}
}
