package main

import (
	"fmt"

	"github.com/kyokomi/gomajan/mjp"
)

func main() {

	p := mjp.NewPlayer(nil, nil)
	// ピンフ、一気通貫
	p.TehaiSet(mjp.M1, 1)
	p.TehaiSet(mjp.M2, 1)
	p.TehaiSet(mjp.M3, 1)
	p.TehaiSet(mjp.M4, 1)
	p.TehaiSet(mjp.M5, 1)
	p.TehaiSet(mjp.M6, 1)
	p.TehaiSet(mjp.M7, 1)
	p.TehaiSet(mjp.M8, 1)
	p.TehaiSet(mjp.M9, 1)
	p.TehaiSet(mjp.P1, 1)
	p.TehaiSet(mjp.P2, 1)
	p.TehaiSet(mjp.P3, 1)
	p.TehaiSet(mjp.S2, 2)

	yakuCheck(p)
}

func yakuCheck(p mjp.Player) {
	fmt.Println(p)

	// TODO: 少牌判定

	// TODO: 多牌判定

	/////////////////////
	y := p.NewYakuCheck()
	fmt.Println(y.String())

	// TODO: 雀頭の判定

	// TODO: 面子の判定
	// TODO: 対子と順子どっち優先？
}
