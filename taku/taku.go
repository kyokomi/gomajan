package taku

import (
	"github.com/kyokomi/gomajan/pai"
	"github.com/kyokomi/gomajan/player"
	"github.com/kyokomi/gomajan/sai"
	"github.com/kyokomi/gomajan/taku/calc"
	"github.com/kyokomi/gomajan/taku/hora"
	"github.com/kyokomi/gomajan/taku/oyako"
	"github.com/kyokomi/gomajan/tehai"
	"github.com/kyokomi/gomajan/yama"
	"fmt"
)

// BaType 場区分
type BaType int

const (
	// TonBa 東場
	TonBa BaType = (0 + iota)
	// NanBa 南場
	NanBa
)

// Taku 麻雀卓
type Taku struct {
	Yama yama.Yama
	Sai  sai.Sai
	// Ba 場
	Ba BaType
	// Kyoku 局
	Kyoku int
	// Honba 本場
	Honba int
	// Jyunme 順目
	Jyunme int
	// Players プレイヤー
	Players []player.Player
	// Nokori 残り牌
	Nokori []tehai.Tehai
}

// NewTaku 対局1回分の卓を生成する
func NewTaku() *Taku {
	var p [4]player.Player
	for i := 0; i < len(p); i++ {
		p[i] = player.NewPlayer(i + 1)
	}

	return &Taku{
		Yama:    yama.New(),
		Sai:     sai.DoubleDiceRoll(),
		Ba:      TonBa,
		Kyoku:   1,
		Honba:   0,
		Jyunme:  0,
		Players: p[:],
		Nokori:  tehai.NewTakuPai(),
	}
}

func (t Taku) Dora() []pai.MJP {
	yamaIdx := 山Index(t.Sai)

	retu := t.Sai.Sum()-1-2
	if retu < 0 {
		retu = len(t.Yama[0]) - 1
		yamaIdx--
		if yamaIdx < 0 {
			yamaIdx = len(t.Yama) - 1
		}
	}

	// TODO: カンの有無確認

	return []pai.MJP{t.Yama[yamaIdx][0][retu]}
}

// UraDora 裏ドラ判定はリーチしてたら読んでいい
func (t Taku) UraDora() []pai.MJP {
	yamaIdx := 山Index(t.Sai)

	retu := t.Sai.Sum()-1-2
	if retu < 0 {
		retu = len(t.Yama[0]) - 1
		yamaIdx--
		if yamaIdx < 0 {
			yamaIdx = len(t.Yama) - 1
		}
	}

	// TODO: カンの有無確認

	return []pai.MJP{t.Yama[yamaIdx][1][retu]}
}

func 山Index(s sai.Sai) int {
	switch s.Sum() {
	case 2, 6, 10:
		return 1
	case 3, 7, 11:
		return 2
	case 4, 8, 12:
		return 3
	case 5, 9:
		return 0
	}
	panic(fmt.Errorf("%s = %d", "不正なサイコロです", s.Sum()))
}

// RonCalcPoint ロン得点計算
func (t Taku) RonCalcPoint(playerID, targetID int, agariPai pai.MJP) *calc.CalcPoint {
	var c calc.CalcPoint

	c.Hora = hora.Ron
	c.TargetID = targetID

	if t.Kyoku == playerID {
		c.Oyako = oyako.Oya
	} else {
		c.Oyako = oyako.Ko
	}

	p := t.Players[playerID]

	// 基本役の生成
	yakuCheck := p.NewYakuCheck(agariPai)
	c.Yakus = yakuCheck.Map()

	// TODO: ドラ、赤ドラ、特殊上がり系（嶺上開花など）を付与する

	// TODO: 符計算
	c.Fu = 30

	// TODO: 得点計算
	switch c.Hora {
	case hora.Ron:
		c.TokutenRon = 3900
	case hora.Tsumo:
		c.TokutenTsumo = [2]int{2000, 1000}
	}

	return &c
}

// TsumoCalcPoint ツモ得点計算
func (t Taku) TsumoCalcPoint(playerIdx, agariPai pai.MJP) *calc.CalcPoint {
	var c calc.CalcPoint

	// TODO: 未実装

	return &c
}

// TODO: 河底撈魚
func (t Taku) isLastAttack() bool {
	return false
}

// TODO: 海底摸月
func (t Taku) isLastAttackAll() bool {
	return false
}
