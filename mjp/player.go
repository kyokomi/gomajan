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

// MentuCheck 面子判定結果
type MentuCheck struct {
	agari pai.MJP // あがり牌
	mentsu [][]pai.MJP // 面子
	nakiMentsu [][]pai.MJP // 鳴き面子
	jyanto pai.MJP // 雀頭
	nokori []Tehai // 面子外残り
}

// YakuCheck 役判定結果
type YakuCheck struct {
	mentsuCheck MentuCheck // 面子判定結果
	yakus       []Yaku     // 役
}

// Yakus getter yakus
func (y YakuCheck) Yakus() []Yaku {
	return y.yakus
}

func (y YakuCheck) String() string {

	var yakus string
	for _, yaku := range y.Yakus() {
		yakus += (" " + yaku.Name)
		if yaku.Name == 国士無双.Name || yaku.Name == 七対子.Name {
			return fmt.Sprintf("役 %s", yakus)
		}
	}

	mc := y.mentsuCheck

	var mentsu string
	for _, m := range mc.mentsu {
		for _, p := range m {
			mentsu += (" " + p.String())
		}
		mentsu += " |"
	}

	for _, m := range mc.nakiMentsu {
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
	for _, n := range mc.nokori {
		if n.val >= 1 {
			nokori += (" " + n.pai.String())
		}
	}
	if nokori == "" {
		nokori = " なし"
	}

	return fmt.Sprintf("雀頭 %s 面子|%s 残り%s => 役 %s", mc.jyanto, mentsu, nokori, yakus)
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

// TehaiSet 手牌設定
func (p *Player) TehaiSet(m pai.MJP, v int) {
	p.tiles[m].val = v
}

func (p Player) newMentuCheck() MentuCheck {
	var mc MentuCheck

	// 面子
	mc.mentsu = make([][]pai.MJP, 0)

	// 鳴き面子
	mc.nakiMentsu = make([][]pai.MJP, 0)

	// 残り牌（テンパイ判定用）
	mc.nokori = make([]Tehai, pai.PaiSize())
	copy(mc.nokori, p.tiles)

	// 面子がひとつも出来ない場合、判定終わり
	for {
		men := checkMentsu(mc.nokori)
		if len(men) == 0 {
			break
		}

		for _, m := range men {
			// 完成した面子を更新
			mc.mentsu = append(mc.mentsu, m)

			// 残り牌を更新
			for _, p := range m {
				mc.nokori[p].val--
			}
		}
	}

	// 鳴いてる時点で面子確定
	for _, f := range p.foos {
		mc.nakiMentsu = append(mc.nakiMentsu, f.Mentsu())
	}

	// 手牌を雀頭と面子に分解する
	// 面子作成後の残り牌から雀頭を作成
	for _, n := range mc.nokori {
		// 雀頭
		if n.val == 2 {
			mc.jyanto = n.pai
			mc.nokori[n.pai].val -= 2
			break
		}
	}

	return mc
}

// NewYakuCheck 役チェック
func (p Player) NewYakuCheck() *YakuCheck {
	y := YakuCheck{}

	y.mentsuCheck = p.newMentuCheck()
	// TODO: こっちにもsetしないといけないのがイマイチ
	p.yaku = &y

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
