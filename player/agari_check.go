package player

import (
	"github.com/kyokomi/gomajan/agari"
	"github.com/kyokomi/gomajan/pai"
)

// CheckAgari 上がり系取得
func (m MentuCheck) CheckAgari() []agari.Agari {
	var agaris []agari.Agari

	for _, pais := range m.mentsu {

		// 上がり牌を含む面子か？
		isAgari := false
		// 上がり牌抜きの上がり形判定の残り牌
		var ac []pai.MJP
		if isAgari, ac = createAgariPais(m.agari, pais); !isAgari {
			continue
		}

		if len(ac) == 0 {
			agaris = append(agaris, agari.Agari{
				Pai:     m.agari,
				Syanten: [2]pai.MJP{m.agari, m.agari},
				Type:    agari.Shabo,
			})
			// 雀頭の対子も入れる
			agaris = append(agaris, agari.Agari{
				Pai:     m.jyanto,
				Syanten: [2]pai.MJP{m.jyanto, m.jyanto},
				Type:    agari.Shabo,
			})

			// シャボ待ちは確定
			return agaris
		}

		if ac[0] == (ac[1] - 1) {
			// 繋がってる（両面待ち、辺張待ち）
			if ac[0].Is19() || ac[1].Is19() {
				// 辺張待ち
				agaris = append(agaris, agari.Agari{
					Pai:     m.agari,
					Syanten: [2]pai.MJP{ac[0], ac[1]},
					Type:    agari.Penchan,
				})

			} else {
				// 両面待ち
				agaris = append(agaris, agari.Agari{
					Pai:     m.agari,
					Syanten: [2]pai.MJP{ac[0], ac[1]},
					Type:    agari.Ryanmen,
				})
			}
		} else {
			// つながってない（嵌張待ち）

			// 嵌張待ち
			agaris = append(agaris, agari.Agari{
				Pai:     m.agari,
				Syanten: [2]pai.MJP{ac[0], ac[1]},
				Type:    agari.Kanchan,
			})
		}
	}

	return agaris
}

func createAgariPais(agari pai.MJP, pais []pai.MJP) (bool, []pai.MJP) {
	// 上がり牌を含む面子か？
	isAgari := false
	// 上がり牌抜きの上がり形判定の残り牌
	var ac []pai.MJP

	for _, pa := range pais {
		if pa.Type() == pai.NoneType {
			continue
		}

		if agari == pa {
			isAgari = true
			continue
		}
		ac = append(ac, pa)
	}

	return isAgari, ac
}
