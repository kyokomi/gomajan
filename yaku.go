package main

import "github.com/kyokomi/gomajan/mjp"

// 国士無双.
func isKokushimusou(t []Tehai) bool {
	if  t[mjp.M1].val >= 1 &&
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
		return true;
	}
	return false;
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

