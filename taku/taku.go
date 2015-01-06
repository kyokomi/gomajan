package taku

import (
	"github.com/kyokomi/gomajan/mjp"
	"github.com/kyokomi/gomajan/mjp/pai"
	"github.com/kyokomi/gomajan/mjp/tehai"
	"github.com/kyokomi/gomajan/taku/calc"
	"github.com/kyokomi/gomajan/taku/hora"
	"github.com/kyokomi/gomajan/taku/oyako"
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
	// Ba 場
	Ba BaType
	// Kyoku 局
	Kyoku int
	// Honba 本場
	Honba int
	// Jyunme 順目
	Jyunme int
	// Dora ドラ
	Dora []pai.MJP
	// UraDora 裏ドラ
	UraDora []pai.MJP
	// Players プレイヤー
	Players []mjp.Player
	// Nokori 残り牌
	Nokori []tehai.Tehai
}

// NewTaku 対局1回分の卓を生成する
func NewTaku() *Taku {

	var p [4]mjp.Player
	for i := 0; i < len(p); i++ {
		p[i] = mjp.NewPlayer(i + 1)
	}

	return &Taku{
		Ba:      TonBa,
		Kyoku:   1,
		Honba:   0,
		Jyunme:  0,
		Dora:    []pai.MJP{pai.M1}, // TODO: ランダムにする
		UraDora: []pai.MJP{pai.M1}, // TODO: ランダムにする
		Players: p[:],
		Nokori:  tehai.NewTakuPai(),
	}
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
