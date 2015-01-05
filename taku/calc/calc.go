package calc

import (
	"bytes"
	"text/template"

	"github.com/kyokomi/gomajan/taku/hora"
	"github.com/kyokomi/gomajan/taku/oyako"
)

// CalcPoint 和了点数計算結果
// 上がった相手
// 対局情報(東場or南場、海底摸月or河底撈魚or槍槓or嶺上開花、天和or地和or人和、ドラ、裏ドラ）
type CalcPoint struct {
	// Oyako 親子区分
	Oyako oyako.OyakoType
	// Hora 和了区分
	Hora hora.HoraType
	// Yakus 役名:翻数
	Yakus map[string]int
	// Fu 符
	Fu int
	// TokutenRon ロン点数
	TokutenRon int
	// TokutenTsumo ツモ点数
	TokutenTsumo [2]int
}

const templateText = `
# 役
{{range $key, $val := .Yakus}}- {{$key}} {{$val}}翻
{{end}}
# 点数
{{.Oyako}} {{.Fu}}符 {{.TotalFan}}翻
{{if .IsRon}}
{{.TokutenRon}}点
{{else}}
{{.TokutenTsumo}}点
{{end}}
`

func (c CalcPoint) String() string {
	var w bytes.Buffer
	t := template.Must(template.New("main").Parse(templateText))
	if err := t.Execute(&w, c); err != nil {
		return err.Error()
	}

	return w.String()
}

// TotalFan 合計翻数を取得
func (c CalcPoint) TotalFan() int {
	fun := 0
	for _, f := range c.Yakus {
		fun += f
	}
	return fun
}

// IsRon ロンか判定
func (c CalcPoint) IsRon() bool {
	return c.Hora == hora.Ron
}

// IsTsumo ツモか判定
func (c CalcPoint) IsTsumo() bool {
	return c.Hora == hora.Tsumo
}
