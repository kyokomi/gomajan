package main

import "github.com/kyokomi/gomajan/mjp"

func checkMentsu(nokori []Tehai) [][]mjp.MJP {
	// 面子
	mentsu := make([][]mjp.MJP, 0)

	// 残り牌からチェック
	tiles := make([]Tehai, 34)
	copy(tiles, nokori)

	// 面子候補
	temp := make([]mjp.MJP, 0)

	for _, t := range tiles {
		if t.val < 1 {
			continue
		}

		if len(temp) > 0 && temp[len(temp) - 1] == (t.pai - 1) {
			// 順子

			// 面子候補追加
			temp = append(temp, t.pai)
		} else if t.val >= 3 {
			// 暗刻

			// 面子候補リセット
			temp = make([]mjp.MJP, 0)
			for i := 0; i < 3; i++ {
				temp = append(temp, t.pai)
			}
		} else {
			// 面子候補リセット
			temp = make([]mjp.MJP, 0)
			temp = append(temp, t.pai)
		}

		// 面子完成
		if len(temp) == 3 {
			mentsu = append(mentsu, temp)
			temp = make([]mjp.MJP, 0)
		}
	}

	return mentsu
}

