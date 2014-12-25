package pai

// MJPType 麻雀牌の種類
type MJPType int
// MJP 麻雀牌
type MJP int

const (
	// NoneType 初期値
	NoneType MJPType = (0 + iota)
	// PType 筒子
	PType
	// MType 萬子
	MType
	// SType 索子
	SType
	// KType 風牌
	KType
	// GType 字牌
	GType
)
const (
	// NonePai 初期値
	NonePai MJP = (0 + iota)

	// P1 筒子の1
	P1
	// P2 筒子の2
	P2
	// P3 筒子の3
	P3
	// P4 筒子の4
	P4
	// P5 筒子の5
	P5
	// P6 筒子の6
	P6
	// P7 筒子の7
	P7
	// P8 筒子の8
	P8
	// P9 筒子の9
	P9

	// M1 萬子の1
	M1
	// M2 萬子の2
	M2
	// M3 萬子の3
	M3
	// M4 萬子の4
	M4
	// M5 萬子の5
	M5
	// M6 萬子の6
	M6
	// M7 萬子の7
	M7
	// M8 萬子の8
	M8
	// M9 萬子の9
	M9

	// S1 索子の1
	S1
	// S2 索子の2
	S2
	// S3 索子の3
	S3
	// S4 索子の4
	S4
	// S5 索子の5
	S5
	// S6 索子の6
	S6
	// S7 索子の7
	S7
	// S8 索子の8
	S8
	// S9 索子の9
	S9

	// TON 東
	TON
	// NAN 南
	NAN
	// SHA 西
	SHA
	// PEI 北
	PEI

	// HAK 白
	HAK
	// HAT 発
	HAT
	// CHN 中
	CHN

	paiSize
)

// PaiSize 麻雀牌の数
func PaiSize() int {
	return int(paiSize)
}

// Type 麻雀牌の種類
func (i MJP) Type() MJPType {
	if i >= P1 && i <= P9 {
		return PType
	} else if i >= M1 && i <= M9 {
		return MType
	} else if i >= S1 && i <= S9 {
		return SType
	} else if i >= TON && i <= PEI {
		return KType
	} else if i >= HAK && i <= CHN {
		return GType
	}
	return NoneType
}

// Is19 1,9牌かを返す
func (i MJP) Is19() bool {
	if i == P1 || i == P9 ||
		i == M1 || i == M9 ||
		i == S1 || i == S9 {
		return true
	}
	return false
}

// IsJipai 字牌かを返す
func (i MJP) IsJipai() bool {
	if i.Type() == GType || i.Type() == KType {
		return true
	}
	return false
}

// IsJipai19 1,9字牌かを返す
func (i MJP) IsJipai19() bool {
	if i.IsJipai() || i.Is19() {
		return true
	}
	return false
}

func (i MJP) String() string {
	switch i {
	default:
		return "None"
	case P1:
		return "P1"
	case P2:
		return "P2"
	case P3:
		return "P3"
	case P4:
		return "P4"
	case P5:
		return "P5"
	case P6:
		return "P6"
	case P7:
		return "P7"
	case P8:
		return "P8"
	case P9:
		return "P9"

	case M1:
		return "M1"
	case M2:
		return "M2"
	case M3:
		return "M3"
	case M4:
		return "M4"
	case M5:
		return "M5"
	case M6:
		return "M6"
	case M7:
		return "M7"
	case M8:
		return "M8"
	case M9:
		return "M9"

	case S1:
		return "S1"
	case S2:
		return "S2"
	case S3:
		return "S3"
	case S4:
		return "S4"
	case S5:
		return "S5"
	case S6:
		return "S6"
	case S7:
		return "S7"
	case S8:
		return "S8"
	case S9:
		return "S9"

		// 風牌
	case TON:
		return "東"
	case NAN:
		return "南"
	case SHA:
		return "西"
	case PEI:
		return "北"

		// 三元牌
	case HAK:
		return "白"
	case HAT:
		return "發"
	case CHN:
		return "中"
	}
}
