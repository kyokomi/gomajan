package main

import (
	"fmt"
	"github.com/kyokomi/gomajan/mjp"
)

type Tehai struct {
	pai mjp.MJP
	val int
}

type Player struct {
	// 手牌
	tiles []Tehai
	// フーロ
	foo [][]mjp.MJP
}

func NewPlayer() Player {
	p := Player{}
	// 33種類
	p.tiles = NewTehai(nil)
	// 最大4フーロ
	p.foo = make([][]mjp.MJP, 4)

	// TODO: ランダムな牌を設定する

	return p
}

func NewTehai(tehai map[mjp.MJP]int) []Tehai {
	tiles := make([]Tehai, 34)
	for i := 0; i < 34; i++ {
		tiles[i].pai = mjp.MJP(i)

		if tehai != nil && tehai[tiles[i].pai] > 0 {
			tiles[i].val = tehai[tiles[i].pai]
		} else {
			tiles[i].val = 0
		}
	}

	return tiles
}

// 国士無双.
func (p Player) isKokushimusou() bool {
	if p.tiles[mjp.M1].val >= 1 &&
		p.tiles[mjp.M9].val >= 1 &&
		p.tiles[mjp.S1].val >= 1 &&
		p.tiles[mjp.S9].val >= 1 &&
		p.tiles[mjp.P1].val >= 1 &&
		p.tiles[mjp.P9].val >= 1 &&
		p.tiles[mjp.TON].val >= 1 &&
		p.tiles[mjp.NAN].val >= 1 &&
		p.tiles[mjp.SHA].val >= 1 &&
		p.tiles[mjp.PEI].val >= 1 &&
		p.tiles[mjp.HAK].val >= 1 &&
		p.tiles[mjp.HAT].val >= 1 &&
		p.tiles[mjp.CHN].val >= 1 {
		return true;
	}
	return false;
}

// 清一色.
func (p Player) isChinniTsu() bool {
	mjpType := mjp.NONE_TYPE
	for _, tehai := range p.tiles {
		if tehai.val < 1 {
			continue
		}

		if mjpType == mjp.NONE_TYPE {
			mjpType = tehai.pai.Type()
		} else if mjpType != tehai.pai.Type() {
			return false
		}
	}
	return true
}

// 七対子.
func (p Player) isNikoNiko() bool {
	count := 0
	for _, tehai := range p.tiles {
		if tehai.val != 2 {
			continue
		}
		count++
	}
	return count == 7
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
	if p.isKokushimusou() {
		// 確定
		res = append(res, "国士無双")
		return res
	}

	// 七対子判定
	if p.isNikoNiko() {
		// TODO: チンイツとかホンイツとかホンロウ等もありえる
		res = append(res, "七対子")
	}

	// 清一色判定
	if p.isChinniTsu() {
		res = append(res, "清一色")
	}

	// 手牌を雀頭と面子に分解する
	var jyanto mjp.MJP
	mentsu := make([][]mjp.MJP, 0)

	temp := make([]mjp.MJP, 0)

	// TODO: tempつくって面子できる毎にtempを更新して余ったやつを雀頭にする
	for _, tehai := range p.tiles {
		if tehai.val < 1 {
			continue
		}
		// 雀頭
		if tehai.val == 2 {
			jyanto = tehai.pai
		}

		// 面子の判定
		if tehai.val == 1 && len(temp) == 0{
			// 面子候補
			temp = append(temp, tehai.pai)
		} else if tehai.val == 1 {
			if temp[len(temp) - 1] == (tehai.pai - 1) {
				// 面子候補追加
				temp = append(temp, tehai.pai)
			} else {
				// 面子候補リセット
				temp = make([]mjp.MJP, 0)
				temp = append(temp, tehai.pai)
			}
		} else {
			// TODO: 暗刻または4枚
			temp = make([]mjp.MJP, 0)
		}

		// 面子完成
		if len(temp) == 3 {
			mentsu = append(mentsu, temp)
			temp = make([]mjp.MJP, 0)
		}
	}
	fmt.Println("雀頭 ", jyanto)
	fmt.Println("面子 ", mentsu)

	// TODO: 雀頭の判定

	// TODO: 面子の判定
	// TODO: 対子と順子どっち優先？

	return res
}

