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
}
