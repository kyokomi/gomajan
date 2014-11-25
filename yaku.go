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

// 緑一色.
func isRyouiso(t []Tehai) bool {
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
