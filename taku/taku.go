package taku

import (
	"fmt"

	"github.com/kyokomi/gomajan/pai"
	"github.com/kyokomi/gomajan/player"
	"github.com/kyokomi/gomajan/sai"
	"github.com/kyokomi/gomajan/taku/calc"
	"github.com/kyokomi/gomajan/taku/hora"
	"github.com/kyokomi/gomajan/taku/oyako"
	"github.com/kyokomi/gomajan/tehai"
	"github.com/kyokomi/gomajan/yama"
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

	// 取得した牌のマスク
	YamaMask yama.YamaMask
	// 現在の山
	PlayYamaIndex int
	// 現在の列
	PlayRetuIndex int
}

// NewTaku 対局1回分の卓を生成する
func NewTaku() *Taku {
	y := yama.New()
	// debug log
	yama.DebugLog(y)

	taku := &Taku{
		Yama:     y,
		YamaMask: yama.YamaMask{},
		Sai:      sai.DoubleDiceRoll(),
		Ba:       TonBa,
		Kyoku:    1,
		Honba:    0,
		Jyunme:   0,
		Nokori:   tehai.NewTakuPai(),
	}

	tehais := taku.配牌()

	var p [4]player.Player
	for i := 0; i < len(p); i++ {
		p[i] = player.NewPlayer(i+1, tehais[i])
	}
	taku.Players = p[:]

	// debug log
	yama.DebugMaskLog(taku.YamaMask)

	return taku
}

func (t *Taku) 配牌() [4][]tehai.Tehai {
	yamaIdx := 山Index(t.Sai)
	retu := t.Sai.Sum() - 1

	// 4人分配牌

	add列And山越しFunc := func(addRetu int) {
		retu += addRetu
		if retu > 16 {
			retu = 0
			yamaIdx--
			if yamaIdx < 0 {
				yamaIdx = 3
			}
		}
	}

	var p [4][14]pai.MJP
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			add列And山越しFunc(1)
			p[j][(i*4)+0] = t.Yama[yamaIdx][0][retu]
			t.YamaMask[yamaIdx][0][retu] = j + 1
			p[j][(i*4)+1] = t.Yama[yamaIdx][1][retu]
			t.YamaMask[yamaIdx][1][retu] = j + 1

			add列And山越しFunc(1)
			p[j][(i*4)+2] = t.Yama[yamaIdx][0][retu]
			t.YamaMask[yamaIdx][0][retu] = j + 1
			p[j][(i*4)+3] = t.Yama[yamaIdx][1][retu]
			t.YamaMask[yamaIdx][1][retu] = j + 1
		}
	}

	// ちょんちょんとちょん
	p[0][12] = t.Yama[yamaIdx][0][retu]
	t.YamaMask[yamaIdx][0][retu] = 1
	p[1][12] = t.Yama[yamaIdx][1][retu]
	t.YamaMask[yamaIdx][1][retu] = 2
	add列And山越しFunc(1)
	p[2][12] = t.Yama[yamaIdx][0][retu]
	t.YamaMask[yamaIdx][0][retu] = 3
	p[3][12] = t.Yama[yamaIdx][1][retu]
	t.YamaMask[yamaIdx][1][retu] = 4
	add列And山越しFunc(1)
	p[0][13] = t.Yama[yamaIdx][0][retu]
	t.YamaMask[yamaIdx][0][retu] = 1

	var playerTehai [4][]tehai.Tehai
	for idx, pais := range p {
		playerTehai[idx] = tehai.NewTehai(nil)
		for _, pp := range pais {
			if pp.Type() == pai.NoneType {
				continue
			}
			playerTehai[idx][pp].Val++
		}
	}

	t.PlayYamaIndex = yamaIdx
	t.PlayRetuIndex = retu

	return playerTehai
}

// Dora ドラ表示牌を返す
func (t Taku) Dora() []pai.MJP {
	yamaIdx, retu := t.doraIndex()

	// TODO: カンの有無確認

	return []pai.MJP{t.Yama[yamaIdx][0][retu]}
}

// UraDora 裏ドラ表示牌を返す
func (t Taku) UraDora() []pai.MJP {
	yamaIdx, retu := t.doraIndex()

	// TODO: カンの有無確認する

	return []pai.MJP{t.Yama[yamaIdx][1][retu]}
}

// TODO: 残念コードすぎる・・・
func (t Taku) doraIndex() (yamaIdx, retu int) {
	yamaIdx = 山Index(t.Sai)
	// -1はindex, ドラは3枚目なので-2
	retu = t.Sai.Sum() - 1 - 2
	// retuが無いときは、隣の山へ
	if retu < 0 {
		// -1はindex
		retu = len(t.Yama[0]) - 1
		yamaIdx--
		// 山一周したら後ろに戻る
		if yamaIdx < 0 {
			// -1はindex
			yamaIdx = len(t.Yama) - 1
		}
	}
	return
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
