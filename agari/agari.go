package agari

import (
	"fmt"

	"github.com/kyokomi/gomajan/pai"
)

// Agari 上がり形
type Agari struct {
	Pai     pai.MJP
	Syanten [2]pai.MJP
	Type    Type
}

// Type 上がり形区分
// go:generate stringer -type=AgariType
type Type int

const (
	// NoneAgari 初期値
	NoneAgari Type = (0 + iota)
	// Shabo シャボ待ち
	Shabo
	// Penchan 辺張待ち
	Penchan
	// Ryanmen 両面待ち
	Ryanmen
	// Kanchan 嵌張待ち
	Kanchan
)

func (a Agari) String() string {
	return fmt.Sprintf("上がり形区分: %s 上がり形: %s ==> %s ", a.Type, a.Syanten, a.Pai)
}
