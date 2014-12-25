package mjp

import "github.com/kyokomi/gomajan/mjp/pai"

// Tehai 手牌1枚
type Tehai struct {
	pai pai.MJP
	val int
}

// NewTehai 手牌作成
func NewTehai(tehai map[pai.MJP]int) []Tehai {
	tiles := make([]Tehai, pai.PaiSize())
	for i := 0; i < pai.PaiSize(); i++ {
		tiles[i].pai = pai.MJP(i)

		if tehai != nil && tehai[tiles[i].pai] > 0 {
			tiles[i].val = tehai[tiles[i].pai]
		} else {
			tiles[i].val = 0
		}
	}

	return tiles
}

func checkMentsu(nokori []Tehai) [][]pai.MJP {
	// 面子
	var mentsu [][]pai.MJP

	// 残り牌からチェック
	tiles := make([]Tehai, pai.PaiSize())
	copy(tiles, nokori)

	// 面子候補
	var temp []pai.MJP

	for _, t := range tiles {
		if t.val < 1 {
			continue
		}

		if len(temp) > 0 && temp[len(temp)-1] == (t.pai-1) {
			// 順子

			// 面子候補追加
			temp = append(temp, t.pai)
		} else if t.val >= 3 {
			// 暗刻

			// 面子候補リセット
			temp = make([]pai.MJP, 0)
			for i := 0; i < 3; i++ {
				temp = append(temp, t.pai)
			}
		} else {
			// 面子候補リセット
			temp = make([]pai.MJP, 0)
			temp = append(temp, t.pai)
		}

		// 面子完成
		if len(temp) == 3 {
			mentsu = append(mentsu, temp)
			temp = make([]pai.MJP, 0)
		}
	}

	return mentsu
}
