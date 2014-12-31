package mjp

import (
	"github.com/kyokomi/gomajan/mjp/pai"
)

func is国士無双(p Player) bool {
	t := p.tiles
	if t[pai.M1].val >= 1 &&
		t[pai.M9].val >= 1 &&
		t[pai.S1].val >= 1 &&
		t[pai.S9].val >= 1 &&
		t[pai.P1].val >= 1 &&
		t[pai.P9].val >= 1 &&
		t[pai.TON].val >= 1 &&
		t[pai.NAN].val >= 1 &&
		t[pai.SHA].val >= 1 &&
		t[pai.PEI].val >= 1 &&
		t[pai.HAK].val >= 1 &&
		t[pai.HAT].val >= 1 &&
		t[pai.CHN].val >= 1 {
		return true
	}
	return false
}

func is大三元(p Player) bool {

	hak := 0
	hat := 0
	chn := 0
	for _, tehai := range p.tiles {
		if tehai.val < 1 {
			continue
		}

		if tehai.pai == pai.HAK {
			hak += tehai.val
		}
		if tehai.pai == pai.HAT {
			hat += tehai.val
		}
		if tehai.pai == pai.CHN {
			chn += tehai.val
		}
	}

	for _, f := range p.foos {
		if f.fooType == NoneFoo {
			continue
		}

		if f.FooType() == Che {
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

	for _, tehai := range p.tiles {
		if tehai.val < 1 {
			continue
		}

		if !tehai.pai.IsJipai() {
			return false
		}
	}

	for _, f := range p.foos {
		if f.fooType == NoneFoo {
			continue
		}

		if f.FooType() == Che {
			return false
		}

		if !f.NakiPai().IsJipai() {
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
	for _, tehai := range p.tiles {
		if tehai.val < 1 {
			continue
		}

		if tehai.pai == pai.TON {
			ton += tehai.val
		}
		if tehai.pai == pai.NAN {
			nan += tehai.val
		}
		if tehai.pai == pai.SHA {
			sha += tehai.val
		}
		if tehai.pai == pai.PEI {
			pei += tehai.val
		}
	}

	for _, f := range p.foos {
		if f.FooType() == Che {
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
	for _, tehai := range p.tiles {
		if tehai.val < 1 {
			continue
		}

		if tehai.pai == pai.TON {
			ton += tehai.val
		}
		if tehai.pai == pai.NAN {
			nan += tehai.val
		}
		if tehai.pai == pai.SHA {
			sha += tehai.val
		}
		if tehai.pai == pai.PEI {
			pei += tehai.val
		}
	}

	for _, f := range p.foos {
		if f.fooType == NoneFoo {
			continue
		}

		if f.FooType() == Che {
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
	for _, tehai := range p.tiles {
		if tehai.val < 1 {
			continue
		}

		if tehai.val >= 3 {
			c++
		}
	}

	for _, f := range p.foos {
		if f.fooType == NoneFoo {
			continue
		}

		if f.FooType() != AnnKan {
			continue
		}
		c++
	}

	return c == 4
}

func is四槓子(p Player) bool {

	c := 0
	for _, f := range p.foos {
		if f.FooType() != AnnKan && f.FooType() != MinKan {
			continue
		}
		c++
	}

	return c == 4
}

func is清老頭(p Player) bool {

	for _, tehai := range p.tiles {
		if tehai.val < 1 {
			continue
		}

		if !tehai.pai.Is19() {
			return false
		}
	}

	for _, f := range p.foos {
		if f.fooType == NoneFoo {
			continue
		}

		if f.FooType() == Che {
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

	for _, tehai := range p.tiles {
		if tehai.val < 1 {
			continue
		}

		if !checkFunc(tehai.pai) {
			return false
		}
	}

	for _, f := range p.foos {
		if f.fooType == NoneFoo {
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

	for _, tehai := range p.tiles {
		if tehai.val < 1 {
			continue
		}

		if !checkFunc(tehai.pai) {
			return false
		}
	}

	for _, f := range p.foos {
		if f.fooType == NoneFoo {
			continue
		}
		if !checkFunc(f.NakiPai()) {
			return false
		}
	}

	return true
}

func is混老頭(p Player) bool {

	for _, tehai := range p.tiles {
		if tehai.val < 1 {
			continue
		}

		if !tehai.pai.IsJipai19() {
			return false
		}
	}

	for _, f := range p.foos {
		if f.fooType == NoneFoo {
			continue
		}

		if f.FooType() == Che {
			return false
		}

		if !f.NakiPai().IsJipai19() {
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

	for _, tehai := range p.tiles {
		if tehai.val < 1 {
			continue
		}

		if !checkFunc(tehai.pai) {
			return false
		}
	}

	for _, f := range p.foos {
		if f.fooType == NoneFoo {
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

func is純全帯(p Player) bool {

	// TODO: 面子判定が必要なので一旦保留

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
	m1 := pai.Mentsu{pai.M1, pai.M2, pai.M3}
	m2 := pai.Mentsu{pai.M4, pai.M5, pai.M6}
	m3 := pai.Mentsu{pai.M7, pai.M8, pai.M9}

	s1 := pai.Mentsu{pai.S1, pai.S2, pai.S3}
	s2 := pai.Mentsu{pai.S4, pai.S5, pai.S6}
	s3 := pai.Mentsu{pai.S7, pai.S8, pai.S9}

	p1 := pai.Mentsu{pai.P1, pai.P2, pai.P3}
	p2 := pai.Mentsu{pai.P4, pai.P5, pai.P6}
	p3 := pai.Mentsu{pai.P7, pai.P8, pai.P9}

	checkFunc := func(pais pai.Mentsu) {
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

		checkFunc(pai.Mentsu{pais[0], pais[1], pais[2]})
	}

	for _, f := range p.foos {
		if f.fooType != Che {
			continue
		}
		checkFunc(pai.Mentsu{f.mentsu[0], f.mentsu[1], f.mentsu[2]})
	}

	if (m123 && m456 && m789) || (s123 && s456 && s789) || (p123 && p456 && p789) {
		return true
	}

	return false
}

func is対々和(p Player) bool {

	for _, tehai := range p.tiles {
		if tehai.val < 1 {
			continue
		}

		if p.yaku.mentsuCheck.jyanto == tehai.pai {
			continue
		}

		// 雀頭以外は3枚以上ないとダメ
		if tehai.val < 3 {
			return false
		}
	}

	for _, f := range p.foos {
		if f.fooType == NoneFoo {
			continue
		}

		if f.FooType() == Che {
			return false
		}
	}

	return true
}

func is三槓子(p Player) bool {

	c := 0
	for _, f := range p.foos {
		if f.FooType() != AnnKan && f.FooType() != MinKan {
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
	for _, tehai := range p.tiles {
		if tehai.val < 1 {
			continue
		}

		if tehai.pai == pai.HAK {
			hak += tehai.val
		} else if tehai.pai == pai.HAT {
			hat += tehai.val
		} else if tehai.pai == pai.CHN {
			chn += tehai.val
		}
	}

	for _, f := range p.foos {
		if f.fooType == NoneFoo {
			continue
		}

		if f.FooType() == Che {
			continue
		}

		if f.nakiPai == pai.HAK {
			hak += len(f.mentsu)
		} else if f.nakiPai == pai.HAT {
			hat += len(f.mentsu)
		} else if f.nakiPai == pai.CHN {
			chn += len(f.mentsu)
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

func is断么九(p Player) bool {
	for _, tehai := range p.tiles {
		if tehai.val < 1 {
			continue
		}

		// 三元牌、風牌
		if tehai.pai.IsJipai19() {
			return false
		}
	}

	for _, f := range p.foos {
		if f.fooType == NoneFoo {
			continue
		}

		for _, m := range f.mentsu {
			// 三元牌、風牌
			if m.IsJipai19() {
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
	for _, tehai := range p.tiles {
		if tehai.val != 2 {
			continue
		}
		count++
	}
	return count == 7
}

func is三暗刻(p Player) bool {

	c := 0
	for _, tehai := range p.tiles {
		if tehai.val < 1 {
			continue
		}

		if tehai.val >= 3 {
			c++
		}
	}

	for _, f := range p.foos {
		if f.fooType == NoneFoo {
			continue
		}

		if f.fooType == AnnKan {
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
	var hitMentsu *pai.Mentsu = nil
	for idx, pais1 := range p.yaku.mentsuCheck.mentsu {
		m1 := pai.NewMentsu(pais1)
		if m1 == nil {
			continue
		}

		// 重複チェック
		if hitMentsu != nil && m1.Equal(*hitMentsu) {
			continue
		}

		for i := idx+1; i < len(p.yaku.mentsuCheck.mentsu); i++ {
			m2 := pai.NewMentsu(p.yaku.mentsuCheck.mentsu[i])
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
