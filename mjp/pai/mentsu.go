package pai

// Mentsu 面子
type Mentsu [3]MJP

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
