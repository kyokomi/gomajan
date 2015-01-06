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

// PaiInc 手牌加える
func (p *Player) PaiInc(m pai.MJP) {
	p.tiles[m].Val++
}

// PaiDec 手牌捨てる
func (p *Player) PaiDec(m pai.MJP) {
	p.tiles[m].Val--
}

// NewPlayer プレイヤー作成
func NewPlayer(playerID int) Player {
	return newPlayer(playerID, nil, nil)
}

func newPlayer(playerID int, tiles []tehai.Tehai, foos []foo.Foo) Player {
	p := Player{}
	p.playerID = playerID

	// 33種類
	if tiles == nil {
		// TODO: ランダムな牌を設定する（卓側か?）
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

	return p
}
