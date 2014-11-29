package main

import "github.com/kyokomi/gomajan/mjp"

// 国士無双.
func isKokushimusou(t []Tehai) bool {
	if t[mjp.M1].val >= 1 &&
		t[mjp.M9].val >= 1 &&
		t[mjp.S1].val >= 1 &&
		t[mjp.S9].val >= 1 &&
		t[mjp.P1].val >= 1 &&
		t[mjp.P9].val >= 1 &&
		t[mjp.TON].val >= 1 &&
		t[mjp.NAN].val >= 1 &&
		t[mjp.SHA].val >= 1 &&
		t[mjp.PEI].val >= 1 &&
		t[mjp.HAK].val >= 1 &&
		t[mjp.HAT].val >= 1 &&
		t[mjp.CHN].val >= 1 {
		return true
	}
	return false
}

// 大三元.
func isDaisangen(t []Tehai) bool {

	// TODO: 鳴き面子OK

	hak := 0
	hat := 0
	chn := 0
	for _, tehai := range t {
		if tehai.val < 1 {
			continue
		}

		if tehai.pai == mjp.HAK {
			hak += tehai.val
		}
		if tehai.pai == mjp.HAT {
			hat += tehai.val
		}
		if tehai.pai == mjp.CHN {
			chn += tehai.val
		}
	}
	if hak >= 3 && hat >= 3 && chn >= 3 {
		return true
	}

	return false
}

// 字一色.
func isTuiso(t []Tehai) bool {

	// TODO: 鳴き面子OK

	for _, tehai := range t {
		if tehai.val < 1 {
			continue
		}

		if tehai.pai.Type() != mjp.G_TYPE && tehai.pai.Type() != mjp.K_TYPE {
			return false
		}
	}
	return true
}

// 大四喜.
func isDaisushi(t []Tehai) bool {

	// TODO: 鳴き面子OK

	ton := 0
	nan := 0
	sha := 0
	pei := 0
	for _, tehai := range t {
		if tehai.val < 1 {
			continue
		}

		if tehai.pai == mjp.TON {
			ton += tehai.val
		}
		if tehai.pai == mjp.NAN {
			nan += tehai.val
		}
		if tehai.pai == mjp.SHA {
			sha += tehai.val
		}
		if tehai.pai == mjp.PEI {
			pei += tehai.val
		}
	}
	if ton >= 3 && nan >= 3 && sha >= 3 && pei >= 3 {
		return true
	}

	return false
}

// 小四喜
func isSyosushi(t []Tehai) bool {

	// TODO: 鳴き面子OK

	ton := 0
	nan := 0
	sha := 0
	pei := 0
	for _, tehai := range t {
		if tehai.val < 1 {
			continue
		}

		if tehai.pai == mjp.TON {
			ton += tehai.val
		}
		if tehai.pai == mjp.NAN {
			nan += tehai.val
		}
		if tehai.pai == mjp.SHA {
			sha += tehai.val
		}
		if tehai.pai == mjp.PEI {
			pei += tehai.val
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
func isSuanko(t []Tehai) bool {

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

	return c == 4
}

// TODO: 四槓子.

// 清老頭.
func isChinrotou(t []Tehai) bool {

	// TODO: 鳴き面子OK

	for _, tehai := range t {
		if tehai.val < 1 {
			continue
		}

		if !tehai.pai.Is19() {
			return false
		}
	}
	return true
}

// 緑一色.
func isRyouiso(t []Tehai) bool {

	// TODO: 鳴き面子OK

	for _, tehai := range t {
		if tehai.val < 1 {
			continue
		}

		// S2, S3, S4, S6, S8, 発 以外はNG
		if tehai.pai != mjp.S2 &&
			tehai.pai != mjp.S3 &&
			tehai.pai != mjp.S4 &&
			tehai.pai != mjp.S6 &&
			tehai.pai != mjp.S8 &&
			mjp.HAT != tehai.pai {
			return false
		}
	}
	return true
}

// 清一色.
func isChinniTsu(t []Tehai) bool {
	// TODO: 鳴き牌OK
	// TODO: 食い下がり

	mjpType := mjp.NONE_TYPE
	for _, tehai := range t {
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

// 混一色.
func isHonniTsu(t []Tehai) bool {
	jihai := false
	mjpType := mjp.NONE_TYPE
	for _, tehai := range t {
		if tehai.val < 1 {
			continue
		}

		if tehai.pai.Type() == mjp.G_TYPE || tehai.pai.Type() == mjp.K_TYPE {
			jihai = true
			continue
		}

		if mjpType == mjp.NONE_TYPE {
			mjpType = tehai.pai.Type()
		} else if mjpType != tehai.pai.Type() {
			return false
		}
	}
	return jihai
}

// 純全帯.
func isJyunchan(t []Tehai) bool {

	// TODO: 食い下がり

	// TODO: 面子判定が必要なので一旦保留

	return false
}

// 一気通貫.
func isIkkitsukan(t []Tehai) bool {

	// TODO: 食い下がり

	// TODO: 面子判定が必要なので一旦保留

	return false
}

// 小三元.
func isSyosangen(t []Tehai) bool {

	// TODO: 鳴き面子OK

	hak := 0
	hat := 0
	chn := 0
	for _, tehai := range t {
		if tehai.val < 1 {
			continue
		}

		if tehai.pai == mjp.HAK {
			hak += tehai.val
		}
		if tehai.pai == mjp.HAT {
			hat += tehai.val
		}
		if tehai.pai == mjp.CHN {
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
func isTanyao(t []Tehai) bool {
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
func isNikoNiko(t []Tehai) bool {
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
func isSananko(t []Tehai) bool {

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
