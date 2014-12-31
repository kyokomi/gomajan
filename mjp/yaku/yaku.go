package yaku

// Yaku 麻雀の役
type Yaku struct {
	// 翻数
	Fan int
	// 名前
	Name string
}

// Equal 役の名前で比較
func (y1 Yaku) Equal(y2 Yaku) bool {
	return y1.Name == y2.Name
}
