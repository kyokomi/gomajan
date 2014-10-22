package main

import (
	"fmt"
	"github.com/kyokomi/gomajan/mjp"
)

type Player struct {
	// 手牌
	tiles map[mjp.MJP]int
	// フーロ
	foo [][]mjp.MJP
}

func NewPlayer() Player {
	p := Player{}
	// 33種類
	p.tiles = make(map[mjp.MJP]int, 33)
	// 最大4フーロ
	p.foo = make([][]mjp.MJP, 4)

	// TODO: ランダムな牌を設定する

	return p
}

func (p Player) isKokushimusou() bool {
	if p.tiles[mjp.M1] >= 1 &&
		p.tiles[mjp.M9] >= 1 &&
		p.tiles[mjp.S1] >= 1 &&
		p.tiles[mjp.S9] >= 1 &&
		p.tiles[mjp.P1] >= 1 &&
		p.tiles[mjp.P9] >= 1 &&
		p.tiles[mjp.TON] >= 1 &&
		p.tiles[mjp.NAN] >= 1 &&
		p.tiles[mjp.SHA] >= 1 &&
		p.tiles[mjp.PEI] >= 1 &&
		p.tiles[mjp.HAK] >= 1 &&
		p.tiles[mjp.HAT] >= 1 &&
		p.tiles[mjp.CHN] >= 1 {
		return true;
	}
	return false;
}

func (p Player) String() string {
	var tehai string
	// 手牌を表示
	for key, val := range p.tiles {
		if val <= 0 {
			continue
		}
		for i := 0; i < val; i++ {
			tehai += key.String() + " "
		}
	}
	return tehai
}

func main() {

	p := NewPlayer()
	// ピンフ、一気通貫
	p.tiles[mjp.M1] = 1
	p.tiles[mjp.M2] = 1
	p.tiles[mjp.M3] = 1
	p.tiles[mjp.M4] = 1
	p.tiles[mjp.M5] = 1
	p.tiles[mjp.M6] = 1
	p.tiles[mjp.M7] = 1
	p.tiles[mjp.M8] = 1
	p.tiles[mjp.M9] = 1
	p.tiles[mjp.P1] = 1
	p.tiles[mjp.P2] = 1
	p.tiles[mjp.P3] = 1
	p.tiles[mjp.S2] = 2

	fmt.Println(yakuCheck(p))
}

func yakuCheck(p Player) string {
	fmt.Println(p)

	// TODO: 少牌判定

	// TODO: 多牌判定

	// 特殊役の判定

	// TODO: 国士無双判定
	if p.isKokushimusou() {
		// 確定
		return "国士無双 役満"
	}

	// TODO: 七対子判定

	// 手牌を雀頭と面子に分解する

	// TODO: 雀頭の判定

	// TODO: 面子の判定
	// TODO: 対子と順子どっち優先？

	return "ノーテン 役なし"
}

