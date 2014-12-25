package mjp

import (
	"fmt"
	"github.com/kyokomi/gomajan/mjp/pai"
)


// Player プレイヤー
type Player struct {
	// 手牌
	tiles []Tehai
	// フーロ
	foos []Foo
	// 一応持っておく
	yaku *YakuCheck
}

// YakuCheck 役チェック結果
type YakuCheck struct {
	// 面子
	mentsu [][]pai.MJP
	// 鳴き面子
	nakiMentsu [][]pai.MJP
	// 雀頭
	jyanto pai.MJP
	// 面子外残り
	nokori []Tehai
	// 役
	yakus []Yaku
}

func (y YakuCheck) Yakus() []Yaku {
	return y.yakus
}

func (y YakuCheck) String() string {

	var yakus string
	for _, yaku := range y.Yakus() {
		yakus += (" " + yaku.Name)
		if yaku.Name == M_国士無双.Name || yaku.Name == M_七対子.Name {
			return fmt.Sprintf("役 %s", yakus)
		}
	}

	var mentsu string
	for _, m := range y.mentsu {
		for _, p := range m {
			mentsu += (" " + p.String())
		}
		mentsu += " |"
	}

	for _, m := range y.nakiMentsu {
		if len(m) == 0 {
			continue
		}
		mentsu += " ("
		for _, p := range m {
			mentsu += p.String()
		}
		mentsu += ") |"
	}

	var nokori string
	for _, n := range y.nokori {
		if n.val >= 1 {
			nokori += (" " + n.pai.String())
		}
	}
	if nokori == "" {
		nokori = " なし"
	}

	return fmt.Sprintf("雀頭 %s 面子|%s 残り%s => 役 %s", y.jyanto, mentsu, nokori, yakus)
}

// NewPlayer プレイヤー作成
func NewPlayer(tiles []Tehai, foos []Foo) Player {
	p := Player{}
	// 33種類
	if tiles == nil {
		p.tiles = NewTehai(nil)
	} else {
		p.tiles = tiles
	}

	// 最大4フーロ
	if foos == nil {
		p.foos = make([]Foo, 4)
	} else {
		p.foos = foos
	}

	p.yaku = nil
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

func (p *Player) TehaiSet(m pai.MJP, v int) {
	p.tiles[m].val = v
}

// NewYakuCheck 役チェック
func (p Player) NewYakuCheck() *YakuCheck {
	y := YakuCheck{}

	// 面子
	y.mentsu = make([][]pai.MJP, 0)

	// 鳴き面子
	y.nakiMentsu = make([][]pai.MJP, 0)

	// 残り牌（テンパイ判定用）
	y.nokori = make([]Tehai, pai.PaiSize())
	copy(y.nokori, p.tiles)

	// 面子がひとつも出来ない場合、判定終わり
	for {
		men := checkMentsu(y.nokori)
		if len(men) == 0 {
			break
		}

		for _, m := range men {
			// 完成した面子を更新
			y.mentsu = append(y.mentsu, m)

			// 残り牌を更新
			for _, p := range m {
				y.nokori[p].val--
			}
		}
	}

	// 鳴いてる時点で面子確定
	for _, f := range p.foos {
		y.nakiMentsu = append(y.nakiMentsu, f.Mentsu())
	}

	// 手牌を雀頭と面子に分解する
	// 面子作成後の残り牌から雀頭を作成
	for _, n := range y.nokori {
		// 雀頭
		if n.val == 2 {
			y.jyanto = n.pai
			y.nokori[n.pai].val -= 2
			break
		}
	}

	yakuman := p.yakuManCheck()
	if len(yakuman) != 0 {
		// 役満確定したらチェック終わり
		y.yakus = yakuman
	} else {
		/////////////////////
		// 通常役の判定
		y.yakus = p.yakuCheck()
	}

	return &y
}
