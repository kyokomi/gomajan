package player

import (
	"github.com/kyokomi/gomajan/agari"
	"github.com/kyokomi/gomajan/foo"
	"github.com/kyokomi/gomajan/mentsu"
	"github.com/kyokomi/gomajan/pai"
)

func is国士無双(p Player) bool {
	t := p.tiles
	if t[pai.M1].Val >= 1 &&
		t[pai.M9].Val >= 1 &&
		t[pai.S1].Val >= 1 &&
		t[pai.S9].Val >= 1 &&
		t[pai.P1].Val >= 1 &&
		t[pai.P9].Val >= 1 &&
		t[pai.TON].Val >= 1 &&
		t[pai.NAN].Val >= 1 &&
		t[pai.SHA].Val >= 1 &&
		t[pai.PEI].Val >= 1 &&
		t[pai.HAK].Val >= 1 &&
		t[pai.HAT].Val >= 1 &&
		t[pai.CHN].Val >= 1 {
		return true
	}
	return false
}

func is大三元(p Player) bool {

	hak := 0
	hat := 0
	chn := 0
	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if t.Pai == pai.HAK {
			hak += t.Val
		}
		if t.Pai == pai.HAT {
			hat += t.Val
		}
		if t.Pai == pai.CHN {
			chn += t.Val
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}

		if f.FooType() == foo.Che {
			continue
		}

		if f.NakiPai() == pai.HAK {
			hak = len(f.Mentsu())
		}
		if f.NakiPai() == pai.HAT {
			hat = len(f.Mentsu())
		}
		if f.NakiPai() == pai.CHN {
			chn = len(f.Mentsu())
		}
	}

	if hak >= 3 && hat >= 3 && chn >= 3 {
		return true
	}

	return false
}

func is字一色(p Player) bool {

	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if !t.Pai.Is字牌() {
			return false
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}

		if f.FooType() == foo.Che {
			return false
		}

		if !f.NakiPai().Is字牌() {
			return false
		}
	}

	return true
}

func is大四喜(p Player) bool {

	ton := 0
	nan := 0
	sha := 0
	pei := 0
	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if t.Pai == pai.TON {
			ton += t.Val
		}
		if t.Pai == pai.NAN {
			nan += t.Val
		}
		if t.Pai == pai.SHA {
			sha += t.Val
		}
		if t.Pai == pai.PEI {
			pei += t.Val
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.Che {
			continue
		}

		if f.NakiPai() == pai.TON {
			ton += len(f.Mentsu())
		}
		if f.NakiPai() == pai.NAN {
			nan += len(f.Mentsu())
		}
		if f.NakiPai() == pai.SHA {
			sha += len(f.Mentsu())
		}
		if f.NakiPai() == pai.PEI {
			pei += len(f.Mentsu())
		}
	}

	if ton >= 3 && nan >= 3 && sha >= 3 && pei >= 3 {
		return true
	}
	return false
}

func is小四喜(p Player) bool {

	ton := 0
	nan := 0
	sha := 0
	pei := 0
	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if t.Pai == pai.TON {
			ton += t.Val
		}
		if t.Pai == pai.NAN {
			nan += t.Val
		}
		if t.Pai == pai.SHA {
			sha += t.Val
		}
		if t.Pai == pai.PEI {
			pei += t.Val
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}

		if f.FooType() == foo.Che {
			continue
		}

		if f.NakiPai() == pai.TON {
			ton += len(f.Mentsu())
		}
		if f.NakiPai() == pai.NAN {
			nan += len(f.Mentsu())
		}
		if f.NakiPai() == pai.SHA {
			sha += len(f.Mentsu())
		}
		if f.NakiPai() == pai.PEI {
			pei += len(f.Mentsu())
		}
	}

	kaze := []int{
		ton, nan, sha, pei,
	}

	c2 := 0
	c3 := 0
	for _, pai := range kaze {
		if pai >= 3 {
			c3++
		} else if pai == 2 {
			c2++
		}
	}
	// 4種が3枚づつ揃ってる or 3種3枚づつ揃ってて、1種だけ2枚
	return c3 == 4 || (c3 == 3 && c2 == 1)
}

func is四暗刻(p Player) bool {

	c := 0
	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if t.Val >= 3 {
			c++
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}

		if f.FooType() != foo.AnnKan {
			continue
		}
		c++
	}

	return c == 4
}

func is四槓子(p Player) bool {

	c := 0
	for _, f := range p.foos {
		if f.FooType() != foo.AnnKan && f.FooType() != foo.MinKan {
			continue
		}
		c++
	}

	return c == 4
}

func is清老頭(p Player) bool {

	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if !t.Pai.Is19() {
			return false
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}

		if f.FooType() == foo.Che {
			return false
		}

		if !f.NakiPai().Is19() {
			return false
		}
	}
	return true
}

func is緑一色(p Player) bool {

	checkFunc := func(mjpPai pai.MJP) bool {
		// S2, S3, S4, S6, S8, 発 以外はNG
		if mjpPai != pai.S2 &&
			mjpPai != pai.S3 &&
			mjpPai != pai.S4 &&
			mjpPai != pai.S6 &&
			mjpPai != pai.S8 &&
			mjpPai != pai.HAT {
			return false
		}
		return true
	}

	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if !checkFunc(t.Pai) {
			return false
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}

		for _, mjpPai := range f.Mentsu() {
			if !checkFunc(mjpPai) {
				return false
			}
		}
	}

	return true
}

func is清一色(p Player) bool {
	mjpType := pai.NoneType
	checkFunc := func(mjpPai pai.MJP) bool {
		if mjpType == pai.NoneType {
			mjpType = mjpPai.Type()
		} else if mjpType != mjpPai.Type() {
			return false
		}
		return true
	}

	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if !checkFunc(t.Pai) {
			return false
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}
		if !checkFunc(f.NakiPai()) {
			return false
		}
	}

	return true
}

func is混老頭(p Player) bool {

	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if !t.Pai.Is19字牌() {
			return false
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}

		if f.FooType() == foo.Che {
			return false
		}

		if !f.NakiPai().Is19字牌() {
			return false
		}
	}

	return true
}

func is混一色(p Player) bool {

	// 字牌有無
	jihai := false
	// 染めた種類
	mjpType := pai.NoneType

	checkFunc := func(mjpPai pai.MJP) bool {
		if mjpPai.Type() == pai.GType || mjpPai.Type() == pai.KType {
			jihai = true
			return true
		}

		if mjpType == pai.NoneType {
			mjpType = mjpPai.Type()
		} else if mjpType != mjpPai.Type() {
			return false
		}
		return true
	}

	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if !checkFunc(t.Pai) {
			return false
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}
		if !checkFunc(f.NakiPai()) {
			return false
		}
	}

	return jihai
}

func is二盃口(p Player) bool {
	return count一盃口(p) == 2
}

func is混全帯么九(p Player) bool {

	for _, pais := range p.yaku.mentsuCheck.mentsu {

		isJipai19 := false
		for _, m := range pais {
			if m.Is19字牌() {
				isJipai19 = true
				break
			}
		}

		if !isJipai19 {
			return false
		}
	}

	for _, pais := range p.yaku.mentsuCheck.nakiMentsu {

		isJipai19 := false
		for _, m := range pais {
			if m.Is19字牌() {
				isJipai19 = true
				break
			}
		}

		if !isJipai19 {
			return false
		}
	}

	if !p.yaku.mentsuCheck.jyanto.Is19字牌() {
		return false
	}

	return true
}

func is純全帯么九(p Player) bool {

	for _, pais := range p.yaku.mentsuCheck.mentsu {

		is19 := false
		for _, m := range pais {
			if m.Is19() {
				is19 = true
				break
			}
		}

		if !is19 {
			return false
		}
	}

	for _, pais := range p.yaku.mentsuCheck.nakiMentsu {

		is19 := false
		for _, m := range pais {
			if m.Is19() {
				is19 = true
				break
			}
		}

		if !is19 {
			return false
		}
	}

	if !p.yaku.mentsuCheck.jyanto.Is19() {
		return false
	}

	return true
}
func is三色同順(p Player) bool {
	sansyoku := [9]int{0}

	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if t.Pai.Num() == 0 {
			continue
		}

		if t.Val >= 3 {
			sansyoku[t.Pai.Num()-1]++
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}

		if f.FooType() == foo.Che {
			continue
		}

		if f.NakiPai().Num() == 0 {
			continue
		}
		sansyoku[f.NakiPai().Num()-1]++
	}

	for _, count := range sansyoku {
		if count == 3 {
			return true
		}
	}

	return false
}

func is一気通貫(p Player) bool {

	m123 := false
	m456 := false
	m789 := false
	s123 := false
	s456 := false
	s789 := false
	p123 := false
	p456 := false
	p789 := false
	m1 := mentsu.Mentsu{pai.M1, pai.M2, pai.M3}
	m2 := mentsu.Mentsu{pai.M4, pai.M5, pai.M6}
	m3 := mentsu.Mentsu{pai.M7, pai.M8, pai.M9}

	s1 := mentsu.Mentsu{pai.S1, pai.S2, pai.S3}
	s2 := mentsu.Mentsu{pai.S4, pai.S5, pai.S6}
	s3 := mentsu.Mentsu{pai.S7, pai.S8, pai.S9}

	p1 := mentsu.Mentsu{pai.P1, pai.P2, pai.P3}
	p2 := mentsu.Mentsu{pai.P4, pai.P5, pai.P6}
	p3 := mentsu.Mentsu{pai.P7, pai.P8, pai.P9}

	checkFunc := func(pais mentsu.Mentsu) {
		if m1.Equal(pais) {
			m123 = true
		} else if m2.Equal(pais) {
			m456 = true
		} else if m3.Equal(pais) {
			m789 = true
		} else if s1.Equal(pais) {
			s123 = true
		} else if s2.Equal(pais) {
			s456 = true
		} else if s3.Equal(pais) {
			s789 = true
		} else if p1.Equal(pais) {
			p123 = true
		} else if p2.Equal(pais) {
			p456 = true
		} else if p3.Equal(pais) {
			p789 = true
		}
	}

	for _, pais := range p.yaku.mentsuCheck.mentsu {
		if len(pais) != 3 {
			continue
		}

		checkFunc(*(mentsu.NewMentsu(pais)))
	}

	for _, f := range p.foos {
		if f.FooType() != foo.Che {
			continue
		}
		checkFunc(*(mentsu.NewMentsu(f.Mentsu())))
	}

	if (m123 && m456 && m789) || (s123 && s456 && s789) || (p123 && p456 && p789) {
		return true
	}

	return false
}

func is対々和(p Player) bool {

	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if p.yaku.mentsuCheck.jyanto == t.Pai {
			continue
		}

		// 雀頭以外は3枚以上ないとダメ
		if t.Val < 3 {
			return false
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}

		if f.FooType() == foo.Che {
			return false
		}
	}

	return true
}

func is三槓子(p Player) bool {

	c := 0
	for _, f := range p.foos {
		if f.FooType() != foo.AnnKan && f.FooType() != foo.MinKan {
			continue
		}
		c++
	}

	return c == 3
}

func is小三元(p Player) bool {

	hak := 0
	hat := 0
	chn := 0
	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if t.Pai == pai.HAK {
			hak += t.Val
		} else if t.Pai == pai.HAT {
			hat += t.Val
		} else if t.Pai == pai.CHN {
			chn += t.Val
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}

		if f.FooType() == foo.Che {
			continue
		}

		if f.NakiPai() == pai.HAK {
			hak += len(f.Mentsu())
		} else if f.NakiPai() == pai.HAT {
			hat += len(f.Mentsu())
		} else if f.NakiPai() == pai.CHN {
			chn += len(f.Mentsu())
		}
	}

	sangen := []int{
		hak, hat, chn,
	}

	c2 := 0
	c3 := 0
	for _, pai := range sangen {
		if pai >= 3 {
			c3++
		} else if pai == 2 {
			c2++
		}
	}
	// 2種3枚づつ揃ってて、1種だけ2枚
	return c3 == 2 && c2 == 1
}

func is平和(p Player) bool {

	// 鳴いてたらダメ
	for _, f := range p.foos {
		if f.FooType() != foo.NoneFoo {
			return false
		}
	}

	// 暗刻もだめ
	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if t.Val >= 3 {
			return false
		}
	}

	if len(p.yaku.mentsuCheck.mentsu) != 4 {
		return false
	}

	// 両面待ちのみ
	agaris := p.yaku.mentsuCheck.CheckAgari()
	for _, a := range agaris {
		if a.Pai != p.yaku.mentsuCheck.agari {
			continue
		}

		if a.Type != agari.Ryanmen {
			return false
		}
	}

	return true
}

func is断么九(p Player) bool {
	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		// 三元牌、風牌
		if t.Pai.Is19字牌() {
			return false
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}

		for _, m := range f.Mentsu() {
			// 三元牌、風牌
			if m.Is19字牌() {
				return false
			}
		}
	}

	return true
}

func is七対子(p Player) bool {
	if len(p.yaku.mentsuCheck.mentsu) != 0 {
		return false
	}

	count := 0
	for _, t := range p.tiles {
		if t.Val != 2 {
			continue
		}
		count++
	}
	return count == 7
}

func is三暗刻(p Player) bool {

	c := 0
	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if t.Val >= 3 {
			c++
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}

		if f.FooType() == foo.AnnKan {
			c++
		}
	}

	return c == 3
}

func is一盃口(p Player) bool {
	return count一盃口(p) == 1
}

// ============================

func count一盃口(p Player) int {
	count := 0
	var hitMentsu *mentsu.Mentsu
	for idx, pais1 := range p.yaku.mentsuCheck.mentsu {
		m1 := mentsu.NewMentsu(pais1)
		if m1 == nil {
			continue
		}

		// 重複チェック
		if hitMentsu != nil && m1.Equal(*hitMentsu) {
			continue
		}

		for i := idx + 1; i < len(p.yaku.mentsuCheck.mentsu); i++ {
			m2 := mentsu.NewMentsu(p.yaku.mentsuCheck.mentsu[i])
			if m2 == nil {
				continue
			}

			if m1.Equal(*m2) {
				count++
				hitMentsu = m2
			}
		}
	}
	return count
}
