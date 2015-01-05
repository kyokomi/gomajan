package taku

import (
	"github.com/kyokomi/gomajan/mjp"
	"github.com/kyokomi/gomajan/mjp/pai"
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
}

// DoCalcPoint 得点計算
func (t Taku) DoCalcPoint(playerIdx int) *calc.CalcPoint {

	// 役
	//	hoge 1翻
	// 	fuga 1翻
	//
	// 点数
	//	30符 2翻 2900点
	//	40符 3翻 5200点（子: 1300点 親: 2600点）

	// TODO: sample
	return &calc.CalcPoint{
		Oyako: oyako.Ko,
		Hora:  hora.Ron,
		Yakus: map[string]int{
			"断么九":  1,
			"三色同順": 2,
		},
		Fu:           30,
		TokutenRon:   3900,
		TokutenTsumo: [2]int{2000, 1000},
	}
}

// TODO: 河底撈魚
func (t Taku) isLastAttack() bool {

	return false
}

// TODO: 海底摸月
func (t Taku) isLastAttackAll() bool {

	return false
}
