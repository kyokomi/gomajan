package yama

import (
	"math/rand"

	"github.com/kyokomi/gomajan/pai"
	"github.com/kyokomi/gomajan/tehai"
	"fmt"
	"time"
)

type Yama [4][2][17]pai.MJP

var takuName = map[int]string{
	1: "東",
	2: "南",
	3: "西",
	4: "北",
}

func New() Yama {
	pais := tehai.NewTehai(nil)

	rand.Seed(time.Now().UnixNano())

	var y Yama
	for i := 0; i < len(y); i++ {
		for j := 0; j < len(y[i]); j++ {
			for k := 0; k < len(y[i][j]); k++ {
			RAND:
				p := pai.MJP(rand.Intn(pai.PaiSize() -1)) + 1
				if pais[p].Val == 4 {
					goto RAND
				}
				y[i][j][k] = p
				pais[p].Val++
			}
		}
	}

	// debug log

	for idx, yy := range y {
		for i, yyy := range yy {
			if i == 0 {
				fmt.Print(takuName[idx+1] + " ")
			} else {
				fmt.Print("   ")
			}

			for _, yyyy := range yyy {
				fmt.Print(yyyy)
				fmt.Print(" ")
			}
			fmt.Println()
		}
		fmt.Println()
	}

	return y
}
