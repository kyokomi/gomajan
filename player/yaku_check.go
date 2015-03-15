package player

import (
	"fmt"

	"github.com/kyokomi/gomajan/pai"
)

type 役 struct {
	Fan  int    // 翻数
	Name string // 名前
}

// Equal 役の名前で比較
func (y1 役) Equal(y2 役) bool {
	return y1.Name == y2.Name
}

var (
	国士無双 = 役{Fan: 13, Name: "国士無双"}
	大四喜  = 役{Fan: 13, Name: "大四喜"}
	小四喜  = 役{Fan: 13, Name: "小四喜"}

	大三元 = 役{Fan: 13, Name: "大三元"}
	字一色 = 役{Fan: 13, Name: "字一色"}
	四暗刻 = 役{Fan: 13, Name: "四暗刻"}
	清老頭 = 役{Fan: 13, Name: "清老頭"}
	緑一色 = 役{Fan: 13, Name: "緑一色"}
	四槓子 = 役{Fan: 13, Name: "四槓子"}

	七対子 = 役{Fan: 2, Name: "七対子"}

	平和  = 役{Fan: 1, Name: "平和"}
	断么九 = 役{Fan: 1, Name: "断么九"}
	一盃口 = 役{Fan: 1, Name: "一盃口"}

	混全帯么九 = 役{Fan: 2, Name: "混全帯么九"}
	対々和   = 役{Fan: 2, Name: "対々和"}
	三色同順  = 役{Fan: 2, Name: "三色同順"}
	一気通貫  = 役{Fan: 2, Name: "一気通貫"}
	三暗刻   = 役{Fan: 2, Name: "三暗刻"}
	小三元   = 役{Fan: 2, Name: "小三元"}
	三槓子   = 役{Fan: 2, Name: "三槓子"}

	混老頭   = 役{Fan: 3, Name: "混老頭"}
	混一色   = 役{Fan: 3, Name: "混一色"}
	二盃口   = 役{Fan: 3, Name: "二盃口"}
	純全帯么九 = 役{Fan: 3, Name: "純全帯么九"}

	清一色 = 役{Fan: 6, Name: "清一色"}
)

// YakuCheck 役判定結果
type YakuCheck struct {
	mentsuCheck MentuCheck // 面子判定結果
	yakus       []役        // 役
}

// Yakus getter yakus
func (y YakuCheck) Yakus() []役 {
	return y.yakus
}

// Map create Map
func (y YakuCheck) Map() map[string]int {
	s := make(map[string]int, 0)

	for _, ya := range y.yakus {
		s[ya.Name] = ya.Fan
	}

	return s
}

func (y YakuCheck) String() string {

	var yakus string
	for _, ya := range y.Yakus() {
		yakus += (" " + ya.Name)
		if ya.Name == 国士無双.Name || ya.Name == 七対子.Name {
			return fmt.Sprintf("役 %s", yakus)
		}
	}

	mc := y.mentsuCheck

	var mentsu string
	for _, m := range mc.mentsu {
		for _, p := range m {
			mentsu += (" " + p.String())
		}
		mentsu += " |"
	}

	for _, m := range mc.nakiMentsu {
		if len(m) == 0 {
			continue
		}
		mentsu += " ("
		for _, p := range m {
			mentsu += p.String()
		}
		mentsu += ") |"
	}

	var nokori string
	for _, n := range mc.nokori {
		if n.Val >= 1 {
			nokori += (" " + n.Pai.String())
		}
	}
	if nokori == "" {
		nokori = " なし"
	}

	return fmt.Sprintf("雀頭 %s 面子|%s 残り%s => 役 %s", mc.jyanto, mentsu, nokori, yakus)
}

// NewYakuCheck 役チェック
func (p Player) NewYakuCheck(agari pai.MJP) *YakuCheck {
	y := YakuCheck{}

	y.mentsuCheck = newMentuCheck(agari, p.tiles, p.foos)

	// TODO: こっちにもsetしないといけないのがイマイチ
	p.yaku = &y
	yakuman := p.yakuManCheck()
	if len(yakuman) != 0 {
		// 役満確定したらチェック終わり
		y.yakus = yakuman
	} else {
		/////////////////////
		// 通常役の判定
		y.yakus = p.yakuCheck()
	}

	return &y
}

func (p Player) yakuManCheck() []役 {
	var res []役

	// 国士無双判定
	if is国士無双(p) {
		res = append(res, 国士無双)
		// 国士無双と他の組み合わせはないので終わり
		return res
	}

	// 大四喜 or 小四喜
	if is大四喜(p) {
		res = append(res, 大四喜)
	} else if is小四喜(p) {
		res = append(res, 小四喜)
	}

	// その他
	if is大三元(p) {
		res = append(res, 大三元)
	}

	if is字一色(p) {
		res = append(res, 字一色)
	}

	if is四暗刻(p) {
		res = append(res, 四暗刻)
	}

	if is清老頭(p) {
		res = append(res, 清老頭)
	}

	if is緑一色(p) {
		res = append(res, 緑一色)
	}

	if is四槓子(p) {
		res = append(res, 四槓子)
	}

	return res
}

func (p Player) yakuCheck() []役 {
	var yakus []役

	if is七対子(p) {
		yakus = append(yakus, 七対子)
	}

	// --- 1翻 ---

	// TODO: リーチ
	// TODO: 一発
	// TODO: 門前清自摸

	if is平和(p) {
		yakus = append(yakus, 平和)
	}

	if is断么九(p) {
		yakus = append(yakus, 断么九)
	}

	if is一盃口(p) {
		yakus = append(yakus, 一盃口)
	}

	// TODO: 嶺上開花
	// TODO: 槍槓
	// TODO: 海底摸月
	// TODO: 河底撈魚

	// TODO: 風牌
	// TODO: 三元牌
	// TODO: 自風

	// --- 2 ---

	// TODO: ダブルリーチ

	if is三色同順(p) {
		yakus = append(yakus, 三色同順)
	}

	if is一気通貫(p) {
		// TODO: 食い下がりある
		yakus = append(yakus, 一気通貫)
	}

	if is対々和(p) {
		yakus = append(yakus, 対々和)
	}

	if is三暗刻(p) {
		// TODO: 食い下がりある
		yakus = append(yakus, 三暗刻)
	}

	if is三槓子(p) {
		yakus = append(yakus, 三槓子)
	}

	if is小三元(p) {
		yakus = append(yakus, 小三元)
	}

	// --- 3 ---

	if is混一色(p) {
		yakus = append(yakus, 混一色)
	}

	if is二盃口(p) {
		yakus = append(yakus, 二盃口)
	}

	if is純全帯么九(p) {
		yakus = append(yakus, 純全帯么九)
	} else if is混老頭(p) {
		yakus = append(yakus, 混老頭)
	} else if is混全帯么九(p) {
		yakus = append(yakus, 混全帯么九)
	}

	// --- 6 ---

	if is清一色(p) {
		yakus = append(yakus, 清一色)
	}

	return yakus
}

func count役牌(p Player) int {
	count := 0
	for _, t := range p.tiles {
		if t.Val < 3 {
			continue
		}

		// TODO: 風牌
		if t.Pai.Type() == pai.GType {
			count++
		}
		// TODO: 三元牌
		// TODO: 自風
	}

	return count
}
