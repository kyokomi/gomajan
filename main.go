package main

import (
	"fmt"
	"github.com/kyokomi/gomajan/mjp"
)

// 手牌
var tiles map[mjp.MJP]int

func main() {
	tiles = make(map[mjp.MJP]int, 33)

	// ピンフ、一気通貫
	tiles[mjp.M1] = 1
	tiles[mjp.M2] = 1
	tiles[mjp.M3] = 1
	tiles[mjp.M4] = 1
	tiles[mjp.M5] = 1
	tiles[mjp.M6] = 1
	tiles[mjp.M7] = 1
	tiles[mjp.M8] = 1
	tiles[mjp.M9] = 1
	tiles[mjp.P1] = 1
	tiles[mjp.P2] = 1
	tiles[mjp.P3] = 1
	tiles[mjp.S2] = 2

	// 手牌を表示
	for key, val := range tiles {
		if val <= 0 {
			continue
		}
		for i := 0; i < val; i++ {
			fmt.Print(key.String() + " ")
		}
	}
	fmt.Println()
}

