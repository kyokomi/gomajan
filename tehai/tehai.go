package tehai

import "github.com/kyokomi/gomajan/pai"

// Tehai 手牌1枚
type Tehai struct {
	Pai pai.MJP
	Val int
}

// NewTehai 手牌作成
func NewTehai(tehai map[pai.MJP]int) []Tehai {
	tiles := newTehai(0)

	if tehai == nil {
		return tiles
	}

	for m, count := range tehai {
		tiles[m].Val = count
	}

	return tiles
}

// NewTakuPai 1局の卓用の牌を作成
func NewTakuPai() []Tehai {
	// 全種類4枚
	return newTehai(4)
}

func newTehai(val int) []Tehai {
	tiles := make([]Tehai, pai.PaiSize())
	for i := 0; i < pai.PaiSize(); i++ {
		tiles[i].Pai = pai.MJP(i)
		tiles[i].Val = val
	}
	return tiles
}

// CheckTehaiMentsu 手牌から面子を作成
func CheckTehaiMentsu(tehai []Tehai) [][]pai.MJP {
	// 面子
	var mentsu [][]pai.MJP

	// 残り牌からチェック
	tiles := make([]Tehai, pai.PaiSize())
	copy(tiles, tehai)

	// 面子候補
	var temp []pai.MJP

	for _, t := range tiles {
		if t.Val < 1 {
			continue
		}

		// 連続している
		// 同じ種類
		// 字牌以外
		if len(temp) > 0 && temp[len(temp)-1] == (t.Pai-1) &&
			temp[0].Type() == t.Pai.Type() &&
			t.Pai.Type() != pai.KType && t.Pai.Type() != pai.GType {
			// 順子

			// 面子候補追加
			temp = append(temp, t.Pai)
		} else if t.Val >= 3 {
			// 暗刻

			// 面子候補リセット
			temp = make([]pai.MJP, 0)
			for i := 0; i < 3; i++ {
				temp = append(temp, t.Pai)
			}
		} else {
			// 面子候補リセット
			temp = make([]pai.MJP, 0)
			temp = append(temp, t.Pai)
		}

		// 面子完成
		if len(temp) == 3 {
			mentsu = append(mentsu, temp)
			temp = make([]pai.MJP, 0)

			// TODO: 2枚の牌の片方が順子に組み込まれると1牌余って
			// 次の順子に組み合わせられないので一旦抜ける
			break
		}
	}

	return mentsu
}
