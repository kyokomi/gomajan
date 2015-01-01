package mjp

import (
	"github.com/kyokomi/gomajan/mjp/agari"
	"github.com/kyokomi/gomajan/mjp/pai"
	"github.com/kyokomi/gomajan/mjp/tehai"
)

// MentuCheck 面子判定結果
type MentuCheck struct {
	agari      pai.MJP       // あがり牌
	mentsu     [][]pai.MJP   // 面子
	nakiMentsu [][]pai.MJP   // 鳴き面子
	jyanto     pai.MJP       // 雀頭
	nokori     []tehai.Tehai // 面子外残り
}

func (p Player) newMentuCheck(agari pai.MJP) MentuCheck {
	var mc MentuCheck

	mc.agari = agari

	// 面子
	mc.mentsu = make([][]pai.MJP, 0)

	// 鳴き面子
	mc.nakiMentsu = make([][]pai.MJP, 0)

	// 残り牌（テンパイ判定用）
	mc.nokori = make([]tehai.Tehai, pai.PaiSize())
	copy(mc.nokori, p.tiles)

	// 面子がひとつも出来ない場合、判定終わり
	for {
		men := tehai.CheckTehaiMentsu(mc.nokori)
		if len(men) == 0 {
			break
		}

		for _, m := range men {
			// 完成した面子を更新
			mc.mentsu = append(mc.mentsu, m)

			// 残り牌を更新
			for _, p := range m {
				mc.nokori[p].Val--
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
		if n.Val == 2 {
			mc.jyanto = n.Pai
			mc.nokori[n.Pai].Val -= 2
			break
		}
	}

	return mc
}

// CheckAgari 上がり系取得
func (m MentuCheck) CheckAgari() []agari.Agari {
	var agaris []agari.Agari

	for _, pais := range m.mentsu {

		// 上がり牌を含む面子か？
		isAgari := false
		// 上がり牌抜きの上がり形判定の残り牌
		var ac []pai.MJP

		for _, pa := range pais {
			if pa.Type() == pai.NoneType {
				continue
			}

			if m.agari == pa {
				isAgari = true
				continue
			}
			ac = append(ac, pa)
		}

		if !isAgari {
			continue
		}

		if len(ac) >= 3 {
			agaris = append(agaris, agari.Agari{
				Agari:     m.agari,
				Syanten:   [2]pai.MJP{ac[0], ac[1]},
				AgariType: agari.Shabo,
			})
			// 雀頭の対子も入れる
			agaris = append(agaris, agari.Agari{
				Agari:     m.jyanto,
				Syanten:   [2]pai.MJP{m.jyanto, m.jyanto},
				AgariType: agari.Shabo,
			})

			// シャボ待ちは確定
			return agaris
		}

		if ac[0] == (ac[1] - 1) {
			// 繋がってる（両面待ち、辺張待ち）
			if ac[0].Is19() || ac[1].Is19() {
				// 辺張待ち
				agaris = append(agaris, agari.Agari{
					Agari:     m.agari,
					Syanten:   [2]pai.MJP{ac[0], ac[1]},
					AgariType: agari.Penchan,
				})

			} else {
				// 両面待ち
				agaris = append(agaris, agari.Agari{
					Agari:     m.agari,
					Syanten:   [2]pai.MJP{ac[0], ac[1]},
					AgariType: agari.Ryanmen,
				})
			}
		} else {
			// つながってない（嵌張待ち）

			// 嵌張待ち
			agaris = append(agaris, agari.Agari{
				Agari:     m.agari,
				Syanten:   [2]pai.MJP{ac[0], ac[1]},
				AgariType: agari.Kanchan,
			})
		}
	}

	return agaris
}
