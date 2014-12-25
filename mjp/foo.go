package mjp

import "github.com/kyokomi/gomajan/mjp/pai"

// FooType フーロ区分
type FooType int

const (
	// NoneFoo 初期値
	NoneFoo FooType = (0 + iota)
	// Che チー
	Che
	// Pon ポン
	Pon
	// MinKan 明カン
	MinKan
	// AnnKan 暗カン
	AnnKan
)

// DareType 誰から鳴いたかの区分
type DareType int

const (
	// NoneDare 初期値
	NoneDare DareType = (0 + iota)
	// JiCha 自家
	JiCha
	// KamiCha 上家
	KamiCha
	// Toimen 対面
	Toimen
	// SimoCha 下家
	SimoCha
)

// Foo 鳴き牌
type Foo struct {
	dare DareType
	// どの牌で鳴いたか
	nakiPai pai.MJP
	// 鳴いた牌（全部）
	mentsu  []pai.MJP
	fooType FooType
}

// FooType getter fooType
func (f Foo) FooType() FooType {
	return f.fooType
}

// Mentsu getter mentsu
func (f Foo) Mentsu() []pai.MJP {
	return f.mentsu
}

// NakiPai getter nakiPai
func (f Foo) NakiPai() pai.MJP {
	return f.nakiPai
}

// NewFooPon ポン
func NewFooPon(dare DareType, nakiPai pai.MJP) Foo {
	return Foo{
		dare:    dare,
		nakiPai: nakiPai,
		mentsu:  []pai.MJP{nakiPai, nakiPai, nakiPai},
		fooType: Pon,
	}
}

// NewFooAnnKan 暗槓
func NewFooAnnKan(nakiPai pai.MJP) Foo {
	return Foo{
		dare:    JiCha,
		nakiPai: nakiPai,
		mentsu:  []pai.MJP{nakiPai, nakiPai, nakiPai, nakiPai},
		fooType: AnnKan,
	}
}

// NewFooMinKan 明槓
func NewFooMinKan(dare DareType, nakiPai pai.MJP) Foo {
	return Foo{
		dare:    dare,
		nakiPai: nakiPai,
		mentsu:  []pai.MJP{nakiPai, nakiPai, nakiPai, nakiPai},
		fooType: MinKan,
	}
}
