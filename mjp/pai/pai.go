package pai

type MJPType int
type MJP int

const (
	// NONE_TYPE 初期値
	NONE_TYPE MJPType = (0 + iota)
	P_TYPE
	M_TYPE
	S_TYPE
	K_TYPE
	G_TYPE
)
const (
	// NonePai 初期値
	NonePai MJP = (0 + iota)
	// ピンズ
	P1
	P2
	P3
	P4
	P5
	P6
	P7
	P8
	P9
	// マンズ
	M1
	M2
	M3
	M4
	M5
	M6
	M7
	M8
	M9
	// ソーズ
	S1
	S2
	S3
	S4
	S5
	S6
	S7
	S8
	S9
	// 風牌
	TON
	NAN
	SHA
	PEI
	// 三元牌
	HAK
	HAT
	CHN
	// Size
	paiSize
)

func PaiSize() int {
	return int(paiSize)
}

func (i MJP) Type() MJPType {
	if i >= P1 && i <= P9 {
		return P_TYPE
	} else if i >= M1 && i <= M9 {
		return M_TYPE
	} else if i >= S1 && i <= S9 {
		return S_TYPE
	} else if i >= TON && i <= PEI {
		return K_TYPE
	} else if i >= HAK && i <= CHN {
		return G_TYPE
	}
	return NONE_TYPE
}

func (i MJP) Is19() bool {
	if i == P1 || i == P9 ||
		i == M1 || i == M9 ||
		i == S1 || i == S9 {
		return true
	}
	return false
}

func (i MJP) IsJipai() bool {
	if i.Type() == G_TYPE || i.Type() == K_TYPE {
		return true
	}
	return false
}

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
