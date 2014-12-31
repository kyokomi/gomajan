package mjp

import (
	"fmt"

	"github.com/kyokomi/gomajan/mjp/foo"
	"github.com/kyokomi/gomajan/mjp/mentsu"
	"github.com/kyokomi/gomajan/mjp/pai"
	"github.com/kyokomi/gomajan/mjp/yaku"
)

var (
	国士無双 = yaku.Yaku{Fan: 13, Name: "国士無双"}
	大四喜  = yaku.Yaku{Fan: 13, Name: "大四喜"}
	小四喜  = yaku.Yaku{Fan: 13, Name: "小四喜"}

	大三元 = yaku.Yaku{Fan: 13, Name: "大三元"}
	字一色 = yaku.Yaku{Fan: 13, Name: "字一色"}
	四暗刻 = yaku.Yaku{Fan: 13, Name: "四暗刻"}
	清老頭 = yaku.Yaku{Fan: 13, Name: "清老頭"}
	緑一色 = yaku.Yaku{Fan: 13, Name: "緑一色"}
	四槓子 = yaku.Yaku{Fan: 13, Name: "四槓子"}

	七対子 = yaku.Yaku{Fan: 2, Name: "七対子"}

	断么九 = yaku.Yaku{Fan: 1, Name: "断么九"}
	一盃口 = yaku.Yaku{Fan: 1, Name: "一盃口"}

	対々和  = yaku.Yaku{Fan: 2, Name: "対々和"}
	一気通貫 = yaku.Yaku{Fan: 2, Name: "一気通貫"}
	三暗刻  = yaku.Yaku{Fan: 2, Name: "三暗刻"}
	小三元  = yaku.Yaku{Fan: 2, Name: "小三元"}
	三槓子  = yaku.Yaku{Fan: 2, Name: "三槓子"}

	混老頭 = yaku.Yaku{Fan: 3, Name: "混老頭"}
	混一色 = yaku.Yaku{Fan: 3, Name: "混一色"}
	二盃口 = yaku.Yaku{Fan: 3, Name: "二盃口"}
	純全帯 = yaku.Yaku{Fan: 3, Name: "純全帯"}

	清一色 = yaku.Yaku{Fan: 6, Name: "清一色"}
)

// YakuCheck 役判定結果
type YakuCheck struct {
	mentsuCheck MentuCheck  // 面子判定結果
	yakus       []yaku.Yaku // 役
}

// Yakus getter yakus
func (y YakuCheck) Yakus() []yaku.Yaku {
	return y.yakus
}

func (yc YakuCheck) String() string {

	var yakus string
	for _, y := range yc.Yakus() {
		yakus += (" " + y.Name)
		if y.Name == 国士無双.Name || y.Name == 七対子.Name {
			return fmt.Sprintf("役 %s", yakus)
		}
	}

	mc := yc.mentsuCheck

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
func (p Player) NewYakuCheck() *YakuCheck {
	y := YakuCheck{}

	y.mentsuCheck = p.newMentuCheck()
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

func (p Player) yakuManCheck() []yaku.Yaku {
	var res []yaku.Yaku

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

func (p Player) yakuCheck() []yaku.Yaku {
	var yakus []yaku.Yaku

	if is七対子(p) {
		yakus = append(yakus, 七対子)
	}

	// --- 1翻 ---

	// TODO: リーチ
	// TODO: 一発
	// TODO: 門前清自摸

	// TODO: 平和

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

	// TODO: 混全帯

	// TODO: 三色同順
	// 食い下がり

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

	// TODO: 食い下がり
	if is混老頭(p) {
		yakus = append(yakus, 混老頭)
	}

	if is混一色(p) {
		yakus = append(yakus, 混一色)
	}

	if is二盃口(p) {
		yakus = append(yakus, 二盃口)
	}

	if is純全帯(p) {
		yakus = append(yakus, 純全帯)
	}

	// --- 6 ---

	if is清一色(p) {
		yakus = append(yakus, 清一色)
	}

	return yakus
}

func is国士無双(p Player) bool {
	t := p.tiles
	if t[pai.M1].Val >= 1 &&
		t[pai.M9].Val >= 1 &&
		t[pai.S1].Val >= 1 &&
		t[pai.S9].Val >= 1 &&
		t[pai.P1].Val >= 1 &&
		t[pai.P9].Val >= 1 &&
		t[pai.TON].Val >= 1 &&
		t[pai.NAN].Val >= 1 &&
		t[pai.SHA].Val >= 1 &&
		t[pai.PEI].Val >= 1 &&
		t[pai.HAK].Val >= 1 &&
		t[pai.HAT].Val >= 1 &&
		t[pai.CHN].Val >= 1 {
		return true
	}
	return false
}

func is大三元(p Player) bool {

	hak := 0
	hat := 0
	chn := 0
	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if t.Pai == pai.HAK {
			hak += t.Val
		}
		if t.Pai == pai.HAT {
			hat += t.Val
		}
		if t.Pai == pai.CHN {
			chn += t.Val
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}

		if f.FooType() == foo.Che {
			continue
		}

		if f.NakiPai() == pai.HAK {
			hak = len(f.Mentsu())
		}
		if f.NakiPai() == pai.HAT {
			hat = len(f.Mentsu())
		}
		if f.NakiPai() == pai.CHN {
			chn = len(f.Mentsu())
		}
	}

	if hak >= 3 && hat >= 3 && chn >= 3 {
		return true
	}

	return false
}

func is字一色(p Player) bool {

	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if !t.Pai.IsJipai() {
			return false
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}

		if f.FooType() == foo.Che {
			return false
		}

		if !f.NakiPai().IsJipai() {
			return false
		}
	}

	return true
}

func is大四喜(p Player) bool {

	ton := 0
	nan := 0
	sha := 0
	pei := 0
	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if t.Pai == pai.TON {
			ton += t.Val
		}
		if t.Pai == pai.NAN {
			nan += t.Val
		}
		if t.Pai == pai.SHA {
			sha += t.Val
		}
		if t.Pai == pai.PEI {
			pei += t.Val
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.Che {
			continue
		}

		if f.NakiPai() == pai.TON {
			ton += len(f.Mentsu())
		}
		if f.NakiPai() == pai.NAN {
			nan += len(f.Mentsu())
		}
		if f.NakiPai() == pai.SHA {
			sha += len(f.Mentsu())
		}
		if f.NakiPai() == pai.PEI {
			pei += len(f.Mentsu())
		}
	}

	if ton >= 3 && nan >= 3 && sha >= 3 && pei >= 3 {
		return true
	}
	return false
}

func is小四喜(p Player) bool {

	ton := 0
	nan := 0
	sha := 0
	pei := 0
	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if t.Pai == pai.TON {
			ton += t.Val
		}
		if t.Pai == pai.NAN {
			nan += t.Val
		}
		if t.Pai == pai.SHA {
			sha += t.Val
		}
		if t.Pai == pai.PEI {
			pei += t.Val
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}

		if f.FooType() == foo.Che {
			continue
		}

		if f.NakiPai() == pai.TON {
			ton += len(f.Mentsu())
		}
		if f.NakiPai() == pai.NAN {
			nan += len(f.Mentsu())
		}
		if f.NakiPai() == pai.SHA {
			sha += len(f.Mentsu())
		}
		if f.NakiPai() == pai.PEI {
			pei += len(f.Mentsu())
		}
	}

	kaze := []int{
		ton, nan, sha, pei,
	}

	c2 := 0
	c3 := 0
	for _, pai := range kaze {
		if pai >= 3 {
			c3++
		} else if pai == 2 {
			c2++
		}
	}
	// 4種が3枚づつ揃ってる or 3種3枚づつ揃ってて、1種だけ2枚
	return c3 == 4 || (c3 == 3 && c2 == 1)
}

func is四暗刻(p Player) bool {

	c := 0
	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if t.Val >= 3 {
			c++
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}

		if f.FooType() != foo.AnnKan {
			continue
		}
		c++
	}

	return c == 4
}

func is四槓子(p Player) bool {

	c := 0
	for _, f := range p.foos {
		if f.FooType() != foo.AnnKan && f.FooType() != foo.MinKan {
			continue
		}
		c++
	}

	return c == 4
}

func is清老頭(p Player) bool {

	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if !t.Pai.Is19() {
			return false
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}

		if f.FooType() == foo.Che {
			return false
		}

		if !f.NakiPai().Is19() {
			return false
		}
	}
	return true
}

func is緑一色(p Player) bool {

	checkFunc := func(mjpPai pai.MJP) bool {
		// S2, S3, S4, S6, S8, 発 以外はNG
		if mjpPai != pai.S2 &&
			mjpPai != pai.S3 &&
			mjpPai != pai.S4 &&
			mjpPai != pai.S6 &&
			mjpPai != pai.S8 &&
			mjpPai != pai.HAT {
			return false
		}
		return true
	}

	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if !checkFunc(t.Pai) {
			return false
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}

		for _, mjpPai := range f.Mentsu() {
			if !checkFunc(mjpPai) {
				return false
			}
		}
	}

	return true
}

func is清一色(p Player) bool {
	mjpType := pai.NoneType
	checkFunc := func(mjpPai pai.MJP) bool {
		if mjpType == pai.NoneType {
			mjpType = mjpPai.Type()
		} else if mjpType != mjpPai.Type() {
			return false
		}
		return true
	}

	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if !checkFunc(t.Pai) {
			return false
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}
		if !checkFunc(f.NakiPai()) {
			return false
		}
	}

	return true
}

func is混老頭(p Player) bool {

	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if !t.Pai.IsJipai19() {
			return false
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}

		if f.FooType() == foo.Che {
			return false
		}

		if !f.NakiPai().IsJipai19() {
			return false
		}
	}

	return true
}

func is混一色(p Player) bool {

	// 字牌有無
	jihai := false
	// 染めた種類
	mjpType := pai.NoneType

	checkFunc := func(mjpPai pai.MJP) bool {
		if mjpPai.Type() == pai.GType || mjpPai.Type() == pai.KType {
			jihai = true
			return true
		}

		if mjpType == pai.NoneType {
			mjpType = mjpPai.Type()
		} else if mjpType != mjpPai.Type() {
			return false
		}
		return true
	}

	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if !checkFunc(t.Pai) {
			return false
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}
		if !checkFunc(f.NakiPai()) {
			return false
		}
	}

	return jihai
}

func is二盃口(p Player) bool {
	return count一盃口(p) == 2
}

func is純全帯(_ Player) bool {

	// TODO: 面子判定が必要なので一旦保留

	return false
}

func is一気通貫(p Player) bool {

	m123 := false
	m456 := false
	m789 := false
	s123 := false
	s456 := false
	s789 := false
	p123 := false
	p456 := false
	p789 := false
	m1 := mentsu.Mentsu{pai.M1, pai.M2, pai.M3}
	m2 := mentsu.Mentsu{pai.M4, pai.M5, pai.M6}
	m3 := mentsu.Mentsu{pai.M7, pai.M8, pai.M9}

	s1 := mentsu.Mentsu{pai.S1, pai.S2, pai.S3}
	s2 := mentsu.Mentsu{pai.S4, pai.S5, pai.S6}
	s3 := mentsu.Mentsu{pai.S7, pai.S8, pai.S9}

	p1 := mentsu.Mentsu{pai.P1, pai.P2, pai.P3}
	p2 := mentsu.Mentsu{pai.P4, pai.P5, pai.P6}
	p3 := mentsu.Mentsu{pai.P7, pai.P8, pai.P9}

	checkFunc := func(pais mentsu.Mentsu) {
		if m1.Equal(pais) {
			m123 = true
		} else if m2.Equal(pais) {
			m456 = true
		} else if m3.Equal(pais) {
			m789 = true
		} else if s1.Equal(pais) {
			s123 = true
		} else if s2.Equal(pais) {
			s456 = true
		} else if s3.Equal(pais) {
			s789 = true
		} else if p1.Equal(pais) {
			p123 = true
		} else if p2.Equal(pais) {
			p456 = true
		} else if p3.Equal(pais) {
			p789 = true
		}
	}

	for _, pais := range p.yaku.mentsuCheck.mentsu {
		if len(pais) != 3 {
			continue
		}

		checkFunc(*(mentsu.NewMentsu(pais)))
	}

	for _, f := range p.foos {
		if f.FooType() != foo.Che {
			continue
		}
		checkFunc(*(mentsu.NewMentsu(f.Mentsu())))
	}

	if (m123 && m456 && m789) || (s123 && s456 && s789) || (p123 && p456 && p789) {
		return true
	}

	return false
}

func is対々和(p Player) bool {

	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if p.yaku.mentsuCheck.jyanto == t.Pai {
			continue
		}

		// 雀頭以外は3枚以上ないとダメ
		if t.Val < 3 {
			return false
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}

		if f.FooType() == foo.Che {
			return false
		}
	}

	return true
}

func is三槓子(p Player) bool {

	c := 0
	for _, f := range p.foos {
		if f.FooType() != foo.AnnKan && f.FooType() != foo.MinKan {
			continue
		}
		c++
	}

	return c == 3
}

func is小三元(p Player) bool {

	hak := 0
	hat := 0
	chn := 0
	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if t.Pai == pai.HAK {
			hak += t.Val
		} else if t.Pai == pai.HAT {
			hat += t.Val
		} else if t.Pai == pai.CHN {
			chn += t.Val
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}

		if f.FooType() == foo.Che {
			continue
		}

		if f.NakiPai() == pai.HAK {
			hak += len(f.Mentsu())
		} else if f.NakiPai() == pai.HAT {
			hat += len(f.Mentsu())
		} else if f.NakiPai() == pai.CHN {
			chn += len(f.Mentsu())
		}
	}

	sangen := []int{
		hak, hat, chn,
	}

	c2 := 0
	c3 := 0
	for _, pai := range sangen {
		if pai >= 3 {
			c3++
		} else if pai == 2 {
			c2++
		}
	}
	// 2種3枚づつ揃ってて、1種だけ2枚
	return c3 == 2 && c2 == 1
}

func is断么九(p Player) bool {
	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		// 三元牌、風牌
		if t.Pai.IsJipai19() {
			return false
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}

		for _, m := range f.Mentsu() {
			// 三元牌、風牌
			if m.IsJipai19() {
				return false
			}
		}
	}

	return true
}

func is七対子(p Player) bool {
	if len(p.yaku.mentsuCheck.mentsu) != 0 {
		return false
	}

	count := 0
	for _, t := range p.tiles {
		if t.Val != 2 {
			continue
		}
		count++
	}
	return count == 7
}

func is三暗刻(p Player) bool {

	c := 0
	for _, t := range p.tiles {
		if t.Val < 1 {
			continue
		}

		if t.Val >= 3 {
			c++
		}
	}

	for _, f := range p.foos {
		if f.FooType() == foo.NoneFoo {
			continue
		}

		if f.FooType() == foo.AnnKan {
			c++
		}
	}

	return c == 3
}

func is一盃口(p Player) bool {
	return count一盃口(p) == 1
}

// ============================

func count一盃口(p Player) int {
	count := 0
	var hitMentsu *mentsu.Mentsu = nil
	for idx, pais1 := range p.yaku.mentsuCheck.mentsu {
		m1 := mentsu.NewMentsu(pais1)
		if m1 == nil {
			continue
		}

		// 重複チェック
		if hitMentsu != nil && m1.Equal(*hitMentsu) {
			continue
		}

		for i := idx + 1; i < len(p.yaku.mentsuCheck.mentsu); i++ {
			m2 := mentsu.NewMentsu(p.yaku.mentsuCheck.mentsu[i])
			if m2 == nil {
				continue
			}

			if m1.Equal(*m2) {
				count++
				hitMentsu = m2
			}
		}
	}
	return count
}
