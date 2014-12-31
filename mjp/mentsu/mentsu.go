package mentsu

import "github.com/kyokomi/gomajan/mjp/pai"

// MentsuLen 面子の長さ
const MentsuLen = 3

// Mentsu 面子
type Mentsu [MentsuLen]pai.MJP

// NewMentsu 面子作成
func NewMentsu(pais []pai.MJP) *Mentsu {
	if len(pais) != MentsuLen {
		return nil
	}
	return &Mentsu{pais[0], pais[1], pais[2]}
}

// Equal 面子の一致判定
func (m1 Mentsu) Equal(m2 Mentsu) bool {
	if len(m1) != len(m2) {
		return false
	}

	for idx, m := range m1 {
		if m != m2[idx] {
			return false
		}
	}
	return true
}
