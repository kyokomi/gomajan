package yama

import (
	"testing"

	"github.com/kyokomi/gomajan/pai"
	"github.com/kyokomi/gomajan/tehai"
)

func TestYama(t *testing.T) {
	// test

	y := New()

	pais := tehai.NewTakuPai()
	for i := 0; i < len(y); i++ {
		for j := 0; j < len(y[i]); j++ {
			for k := 0; k < len(y[i][j]); k++ {
				pais[y[i][j][k]].Val--
			}
		}
	}

	// check

	for _, p := range pais {
		if p.Pai == pai.NonePai {
			continue
		}

		if p.Val != 0 {
			t.Errorf("4枚じゃないです %s = %d", p.Pai, p.Val)
		}
	}
}
