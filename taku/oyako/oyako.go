package oyako

// OyakoType 親子区分
type OyakoType int

const (
	// Oya 親
	Oya OyakoType = (iota + 0)
	// Ko 子
	Ko
)

func (o OyakoType) String() string {
	switch o {
	default:
		return ""
	case Oya:
		return "親"
	case Ko:
		return "子"
	}
}
