package mjp

import (
	"github.com/kyokomi/gomajan/mjp/pai"
)

// 国士無双.
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

// 大三元.
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

// 字一色.
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
			continue
		}

		if !f.NakiPai().IsJipai() {
			return false
		}
	}

	return true
}

// 大四喜.
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

// 小四喜
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

// 四暗刻
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

// 四槓子
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

// 清老頭.
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
			continue
		}

		if !f.NakiPai().Is19() {
			return false
		}
	}
	return true
}

// 緑一色.
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

// 清一色.
func is清一色(p Player) bool {
	// TODO: 食い下がり

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

// 混一色.
func is混一色(t []Tehai) bool {
	jihai := false
	mjpType := pai.NoneType
	for _, tehai := range t {
		if tehai.val < 1 {
			continue
		}

		if tehai.pai.Type() == pai.GType || tehai.pai.Type() == pai.KType {
			jihai = true
			continue
		}

		if mjpType == pai.NoneType {
			mjpType = tehai.pai.Type()
		} else if mjpType != tehai.pai.Type() {
			return false
		}
	}
	return jihai
}

// 純全帯.
func is純全帯(t []Tehai) bool {

	// TODO: 食い下がり

	// TODO: 面子判定が必要なので一旦保留

	return false
}

// 一気通貫.
func is一気通貫(t []Tehai) bool {

	// TODO: 食い下がり

	// TODO: 面子判定が必要なので一旦保留

	return false
}

// 小三元.
func is小三元(t []Tehai) bool {

	// TODO: 鳴き面子OK

	hak := 0
	hat := 0
	chn := 0
	for _, tehai := range t {
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

// 断么九.
func is断么九(t []Tehai) bool {
	for _, tehai := range t {
		if tehai.val < 1 {
			continue
		}

		// 三元牌、風牌
		if tehai.pai.IsJipai19() {
			return false
		}
	}
	return true
}

// 七対子.
func is七対子(t []Tehai) bool {
	count := 0
	for _, tehai := range t {
		if tehai.val != 2 {
			continue
		}
		count++
	}
	return count == 7
}

// 三暗刻
func is三暗刻(t []Tehai) bool {

	// TODO: 鳴き面子NG, 槓子OK

	c := 0
	for _, tehai := range t {
		if tehai.val < 1 {
			continue
		}

		if tehai.val >= 3 {
			c++
		}
	}

	return c == 3
}
