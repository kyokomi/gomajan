package main

import (
	"fmt"

	"github.com/kyokomi/gomajan/mjp"
	"github.com/kyokomi/gomajan/mjp/pai"
)

func main() {

	p := mjp.NewPlayer(nil, nil)
	// ピンフ、一気通貫
	p.TehaiSet(pai.M1, 1)
	p.TehaiSet(pai.M2, 1)
	p.TehaiSet(pai.M3, 1)
	p.TehaiSet(pai.M4, 1)
	p.TehaiSet(pai.M5, 1)
	p.TehaiSet(pai.M6, 1)
	p.TehaiSet(pai.M7, 1)
	p.TehaiSet(pai.M8, 1)
	p.TehaiSet(pai.M9, 1)
	p.TehaiSet(pai.P1, 1)
	p.TehaiSet(pai.P2, 1)
	p.TehaiSet(pai.P3, 1)
	p.TehaiSet(pai.S2, 2)

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
