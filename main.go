package main

import (
	"fmt"
	"github.com/kyokomi/gomajan/mjp"
)

func main() {

	p := NewPlayer()
	// ピンフ、一気通貫
	p.tiles[mjp.M1].val = 1
	p.tiles[mjp.M2].val = 1
	p.tiles[mjp.M3].val = 1
	p.tiles[mjp.M4].val = 1
	p.tiles[mjp.M5].val = 1
	p.tiles[mjp.M6].val = 1
	p.tiles[mjp.M7].val = 1
	p.tiles[mjp.M8].val = 1
	p.tiles[mjp.M9].val = 1
	p.tiles[mjp.P1].val = 1
	p.tiles[mjp.P2].val = 1
	p.tiles[mjp.P3].val = 1
	p.tiles[mjp.S2].val = 2

	fmt.Println(yakuCheck(p))
}

func yakuCheck(p Player) []string {
	fmt.Println(p)

	res := make([]string, 0)

	// TODO: 少牌判定

	// TODO: 多牌判定

	// 特殊役の判定

	// 国士無双判定
	if isKokushimusou(p.tiles) {
		// 確定
		res = append(res, "国士無双")
		return res
	}

	// 七対子判定
	if isNikoNiko(p.tiles) {
		// TODO: チンイツとかホンイツとかホンロウ等もありえる
		res = append(res, "七対子")
	}

	// 清一色判定
	if isChinniTsu(p.tiles) {
		res = append(res, "清一色")
	}

	// TODO: 七対子は面子判定不要

	// 手牌を雀頭と面子に分解する
	var jyanto mjp.MJP
	mentsu := make([][]mjp.MJP, 0)

	temp := make([]mjp.MJP, 0)
	nokori := make([]Tehai, 34)
	copy(nokori, p.tiles)

	for {
		isMentsu := false
		tiles := make([]Tehai, 34)
		copy(tiles, nokori)
		for _, tehai := range tiles {
			if tehai.val < 1 {
				continue
			}

			// 面子の判定（順子）
			if len(temp) == 0 {
				// 面子候補
				temp = append(temp, tehai.pai)
			} else {
				if temp[len(temp) - 1] == (tehai.pai - 1) {
					// 面子候補追加
					temp = append(temp, tehai.pai)
				} else if tehai.val >= 3 {
					// 面子候補リセット
					temp = make([]mjp.MJP, 0)
					for i := 0; i < 3; i++ {
						temp = append(temp, tehai.pai)
					}
				} else {
					// 面子候補リセット
					temp = make([]mjp.MJP, 0)
					temp = append(temp, tehai.pai)
				}
			}

			// 面子完成
			if len(temp) == 3 {

				for _, t := range temp {
					nokori[t].val -= 1
				}

				mentsu = append(mentsu, temp)
				temp = make([]mjp.MJP, 0)
				isMentsu = true
			}
		}

		if !isMentsu {
			break
		}
	}

	for _, n := range nokori {
		// 雀頭
		if n.val == 2 {
			jyanto = n.pai
			nokori[n.pai].val -= 2
			break
		}
	}
	fmt.Println("雀頭 ", jyanto)
	fmt.Println("面子 ", mentsu)

	fmt.Print("残り ")
	if !isNikoNiko(p.tiles) {
		for _, n := range nokori {
			if n.val >= 1 {
				fmt.Print(n)
			}
		}
	}
	fmt.Println()

	// TODO: 雀頭の判定

	// TODO: 面子の判定
	// TODO: 対子と順子どっち優先？

	return res
}

