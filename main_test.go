package main

import (
	"testing"
	"github.com/kyokomi/gomajan/mjp"
	"fmt"
)


func TestKokushimusou(t *testing.T) {

	p := NewPlayer()
	// 国士無双
	p.tiles[mjp.S1] = 2
	p.tiles[mjp.S9] = 1
	p.tiles[mjp.P1] = 1
	p.tiles[mjp.P9] = 1
	p.tiles[mjp.M1] = 1
	p.tiles[mjp.M9] = 1
	p.tiles[mjp.TON] = 1
	p.tiles[mjp.NAN] = 1
	p.tiles[mjp.SHA] = 1
	p.tiles[mjp.PEI] = 1
	p.tiles[mjp.HAK] = 1
	p.tiles[mjp.HAT] = 1
	p.tiles[mjp.CHN] = 1

	if yaku := yakuCheck(p); yaku != "国士無双 役満" {
		t.Error("国士無双error 手牌 => ", p)
	} else {
		fmt.Println(yaku)
	}
}
