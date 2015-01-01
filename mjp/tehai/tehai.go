package tehai

import "github.com/kyokomi/gomajan/mjp/pai"

// Tehai 手牌1枚
type Tehai struct {
	Pai pai.MJP
	Val int
}

// NewTehai 手牌作成
func NewTehai(tehai map[pai.MJP]int) []Tehai {
	tiles := make([]Tehai, pai.PaiSize())
	for i := 0; i < pai.PaiSize(); i++ {
		tiles[i].Pai = pai.MJP(i)

		if tehai != nil && tehai[tiles[i].Pai] > 0 {
			tiles[i].Val = tehai[tiles[i].Pai]
		} else {
			tiles[i].Val = 0
		}
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

		if len(temp) > 0 && temp[len(temp)-1] == (t.Pai-1) {
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
		}
	}

	return mentsu
}
