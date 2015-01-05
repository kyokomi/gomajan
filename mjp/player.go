package mjp

import (
	"github.com/kyokomi/gomajan/mjp/foo"
	"github.com/kyokomi/gomajan/mjp/pai"
	"github.com/kyokomi/gomajan/mjp/tehai"
)

// Player プレイヤー
type Player struct {
	// プレイヤーを一意に識別するID
	playerID int
	// 手牌
	tiles []tehai.Tehai
	// フーロ
	foos []foo.Foo
	// 一応持っておく
	yaku *YakuCheck
}

func (p Player) String() string {
	var tehaiStr string
	// 手牌を表示
	for _, t := range p.tiles {
		if t.Val <= 0 {
			continue
		}
		for i := 0; i < t.Val; i++ {
			tehaiStr += t.Pai.String() + " "
		}
	}
	return tehaiStr
}

// TehaiSet 手牌設定
func (p *Player) TehaiSet(m pai.MJP, v int) {
	p.tiles[m].Val = v
}

// NewPlayer プレイヤー作成
func NewPlayer(tiles []tehai.Tehai, foos []foo.Foo) Player {
	p := Player{}
	// 33種類
	if tiles == nil {
		p.tiles = tehai.NewTehai(nil)
	} else {
		p.tiles = tiles
	}

	// 最大4フーロ
	if foos == nil {
		p.foos = make([]foo.Foo, 4)
	} else {
		p.foos = foos
	}

	p.yaku = nil
	// TODO: ランダムな牌を設定する

	return p
}
