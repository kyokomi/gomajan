package yaku

// Yaku 麻雀の役
type Yaku struct {
	Fan  int    // 翻数
	Name string // 名前
}

// Equal 役の名前で比較
func (y1 Yaku) Equal(y2 Yaku) bool {
	return y1.Name == y2.Name
}
