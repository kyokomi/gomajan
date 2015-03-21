package yama

import (
	"math/rand"

	"fmt"
	"time"

	"github.com/kyokomi/gomajan/pai"
	"github.com/kyokomi/gomajan/tehai"
)

type Yama [4][2][17]pai.MJP
type YamaMask [4][2][17]int

var takuName = map[int]string{
	1: "東",
	2: "南",
	3: "西",
	4: "北",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func New() Yama {
	pais := tehai.NewTehai(nil)

	var y Yama
	for i := 0; i < len(y); i++ {
		for j := 0; j < len(y[i]); j++ {
			for k := 0; k < len(y[i][j]); k++ {
			RAND:
				p := pai.MJP(rand.Intn(pai.PaiSize()-1)) + 1
				// すでに4枚あるやつは再抽選
				if pais[p].Val == 4 {
					goto RAND
				}
				y[i][j][k] = p
				pais[p].Val++
			}
		}
	}

	return y
}

func DebugMaskLog(y YamaMask) {
	// debug log
	fmt.Println("     00 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16")
	fmt.Println("-------------------------------------------------------")

	for idx, yy := range y {
		for i, yyy := range yy {
			if i == 0 {
				fmt.Print(idx, " "+takuName[idx+1]+" ")
			} else {
				fmt.Print("     ")
			}

			for _, yyyy := range yyy {
				fmt.Print(" ")
				fmt.Print(yyyy)
				fmt.Print(" ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func DebugLog(y Yama) {
	// debug log
	fmt.Println("     00 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16")
	fmt.Println("-------------------------------------------------------")

	for idx, yy := range y {
		for i, yyy := range yy {
			if i == 0 {
				fmt.Print(idx, " "+takuName[idx+1]+" ")
			} else {
				fmt.Print("     ")
			}

			for _, yyyy := range yyy {
				fmt.Print(yyyy)
				fmt.Print(" ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
