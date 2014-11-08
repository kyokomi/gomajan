package main

import "github.com/kyokomi/gomajan/mjp"

type Tehai struct {
	pai mjp.MJP
	val int
}

type Player struct {
	// 手牌
	tiles []Tehai
	// フーロ
	foo [][]mjp.MJP
}

func NewPlayer() Player {
	p := Player{}
	// 33種類
	p.tiles = NewTehai(nil)
	// 最大4フーロ
	p.foo = make([][]mjp.MJP, 4)

	// TODO: ランダムな牌を設定する

	return p
}

func (p Player) String() string {
	var tehaiStr string
	// 手牌を表示
	for _, tehai := range p.tiles {
		if tehai.val <= 0 {
			continue
		}
		for i := 0; i < tehai.val; i++ {
			tehaiStr += tehai.pai.String() + " "
		}
	}
	return tehaiStr
}

func NewTehai(tehai map[mjp.MJP]int) []Tehai {
	tiles := make([]Tehai, 34)
	for i := 0; i < 34; i++ {
		tiles[i].pai = mjp.MJP(i)

		if tehai != nil && tehai[tiles[i].pai] > 0 {
			tiles[i].val = tehai[tiles[i].pai]
		} else {
			tiles[i].val = 0
		}
	}

	return tiles
}
