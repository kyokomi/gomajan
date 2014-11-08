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

	// 面子
	mentsu := make([][]mjp.MJP, 0)

	// 残り牌（テンパイ判定用）
	nokori := make([]Tehai, 34)
	copy(nokori, p.tiles)

	// 面子がひとつも出来ない場合、判定終わり
	for {
		men := checkMentsu(nokori)
		if len(men) == 0 {
			break
		}

		for _, m := range men {
			// 完成した面子を更新
			mentsu = append(mentsu, m)

			// 残り牌を更新
			for _, p := range m {
				nokori[p].val -= 1
			}
		}
	}

	// 手牌を雀頭と面子に分解する
	var jyanto mjp.MJP
	// 面子作成後の残り牌から雀頭を作成
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
