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

type YakuCheck struct {
	mentsu     [][]mjp.MJP
	nakiMentsu [][]mjp.MJP
	jyanto     mjp.MJP
	nokori     []Tehai
	//	yaku []string
}

func NewYakuCheck(p Player) *YakuCheck {
	y := YakuCheck{}

	// 面子
	y.mentsu = make([][]mjp.MJP, 0)

	// TODO: 鳴き面子
	y.nakiMentsu = make([][]mjp.MJP, 0)

	// 残り牌（テンパイ判定用）
	y.nokori = make([]Tehai, 34)
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
				y.nokori[p].val -= 1
			}
		}
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

	return &y
}

func yakuCheck(p Player) []string {
	fmt.Println(p)

	res := make([]string, 0)

	// TODO: 少牌判定

	// TODO: 多牌判定

	/////////////////////
	// 役満の判定
	yakuman := yakuManCheck(p)
	if len(yakuman) != 0 {
		// 役満確定したらチェック終わり
		res = yakuman
	} else {
		/////////////////////
		// 通常役の判定

		if isNikoNiko(p.tiles) {
			// TODO: チンイツとかホンイツとかホンロウ等もありえる
			res = append(res, "七対子")
		}

		// --- 1翻 ---

		// TODO: リーチ
		// TODO: 一発
		// TODO: 門前清自摸

		// TODO: 平和

		if isTanyao(p.tiles) {
			res = append(res, "断么九")
		}

		// TODO: 一盃口

		// TODO: 嶺上開花
		// TODO: 槍槓
		// TODO: 海底摸月
		// TODO: 河底撈魚

		// TODO: 風牌
		// TODO: 三元牌
		// TODO: 自風

		// --- 2 ---

		// TODO: ダブルリーチ

		// TODO: 混全帯

		// TODO: 三色同順
		// 食い下がり

		if isIkkitsukan(p.tiles) {
			res = append(res, "一気通貫")
		}

		// TODO: 対々和

		if isSananko(p.tiles) {
			res = append(res, "三暗刻")
		}

		// TODO: 三槓子

		if isSyosangen(p.tiles) {
			res = append(res, "小三元")
		}

		// --- 3 ---

		// TODO: 混老頭
		// 食い下がり

		if isHonniTsu(p.tiles) {
			res = append(res, "混一色")
		}

		// TODO: 二盃口

		if isJyunchan(p.tiles) {
			res = append(res, "純全帯")
		}

		// --- 6 ---

		if isChinniTsu(p.tiles) {
			res = append(res, "清一色")
		}
	}

	// TODO: 七対子と国士無双は面子判定不要
	y := NewYakuCheck(p)

	// TODO: test用
	if !isNikoNiko(p.tiles) && !isKokushimusou(p.tiles) {
		fmt.Println("雀頭 ", y.jyanto)
		fmt.Println("面子 ", y.mentsu)

		fmt.Print("残り ")
		for _, n := range y.nokori {
			if n.val >= 1 {
				fmt.Print(n)
			}
		}
		fmt.Println()
	}

	// TODO: 雀頭の判定

	// TODO: 面子の判定
	// TODO: 対子と順子どっち優先？

	return res
}

func yakuManCheck(p Player) []string {
	res := make([]string, 0)

	// 国士無双判定
	if isKokushimusou(p.tiles) {
		res = append(res, "国士無双")
		// 国士無双と他の組み合わせはないので終わり
		return res
	}

	// 大三元
	if isDaisangen(p.tiles) {
		res = append(res, "大三元")
	}

	// 字一色
	if isTuiso(p.tiles) {
		res = append(res, "字一色")
	}

	// 大四喜 or 小四喜
	if isDaisushi(p.tiles) {
		res = append(res, "大四喜")
	} else if isSyosushi(p.tiles) {
		res = append(res, "小四喜")
	}

	// 四暗刻
	if isSuanko(p.tiles) {
		res = append(res, "四暗刻")
	}

	// TODO: 四槓子

	// 清老頭
	if isChinrotou(p.tiles) {
		res = append(res, "清老頭")
	}

	// 緑一色
	if isRyouiso(p.tiles) {
		res = append(res, "緑一色")
	}

	return res
}
