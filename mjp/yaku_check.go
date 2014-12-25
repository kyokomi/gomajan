package mjp

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

var (
	// M国士無双 国士無双
	M国士無双 = Yaku{Fan: 13, Name: "国士無双"}
	// M大四喜 大四喜
	M大四喜 = Yaku{Fan: 13, Name: "大四喜"}
	// M小四喜 小四喜
	M小四喜 = Yaku{Fan: 13, Name: "小四喜"}

	// M大三元 大三元
	M大三元 = Yaku{Fan: 13, Name: "大三元"}
	// M字一色 字一色
	M字一色 = Yaku{Fan: 13, Name: "字一色"}
	// M四暗刻 四暗刻
	M四暗刻 = Yaku{Fan: 13, Name: "四暗刻"}
	// M清老頭 清老頭
	M清老頭 = Yaku{Fan: 13, Name: "清老頭"}
	// M緑一色 緑一色
	M緑一色 = Yaku{Fan: 13, Name: "緑一色"}
	// M四槓子 四槓子
	M四槓子 = Yaku{Fan: 13, Name: "四槓子"}

	// M七対子 七対子
	M七対子 = Yaku{Fan: 2, Name: "七対子"}

	// M断么九 断么九
	M断么九 = Yaku{Fan: 1, Name: "断么九"}

	// M一気通貫 一気通貫
	M一気通貫 = Yaku{Fan: 2, Name: "一気通貫"}
	// M三暗刻 三暗刻
	M三暗刻 = Yaku{Fan: 2, Name: "三暗刻"}
	// M小三元 小三元
	M小三元 = Yaku{Fan: 2, Name: "小三元"}

	// M混一色 混一色
	M混一色 = Yaku{Fan: 3, Name: "混一色"}
	// M純全帯 純全帯
	M純全帯 = Yaku{Fan: 3, Name: "純全帯"}

	// M清一色 清一色
	M清一色 = Yaku{Fan: 6, Name: "清一色"}
)

func (p Player) yakuManCheck() []Yaku {
	var res []Yaku

	// 国士無双判定
	if is国士無双(p) {
		res = append(res, M国士無双)
		// 国士無双と他の組み合わせはないので終わり
		return res
	}

	// 大四喜 or 小四喜
	if is大四喜(p) {
		res = append(res, M大四喜)
	} else if is小四喜(p) {
		res = append(res, M小四喜)
	}

	// その他
	if is大三元(p) {
		res = append(res, M大三元)
	}

	if is字一色(p) {
		res = append(res, M字一色)
	}

	if is四暗刻(p) {
		res = append(res, M四暗刻)
	}

	if is清老頭(p) {
		res = append(res, M清老頭)
	}

	if is緑一色(p) {
		res = append(res, M緑一色)
	}

	if is四槓子(p) {
		res = append(res, M四槓子)
	}

	return res
}

func (p Player) yakuCheck() []Yaku {
	var yakus []Yaku

	if is七対子(p.tiles) {
		yakus = append(yakus, M七対子)
	}

	// --- 1翻 ---

	// TODO: リーチ
	// TODO: 一発
	// TODO: 門前清自摸

	// TODO: 平和

	if is断么九(p.tiles) {
		yakus = append(yakus, M断么九)
	}

	// TODO: 一盃口

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

	if is一気通貫(p.tiles) {
		// TODO: 食い下がりある
		yakus = append(yakus, M一気通貫)
	}

	// TODO: 対々和

	if is三暗刻(p.tiles) {
		// TODO: 食い下がりある
		yakus = append(yakus, M三暗刻)
	}

	// TODO: 三槓子

	if is小三元(p.tiles) {
		yakus = append(yakus, M小三元)
	}

	// --- 3 ---

	// TODO: 混老頭
	// 食い下がり

	if is混一色(p.tiles) {
		yakus = append(yakus, M混一色)
	}

	// TODO: 二盃口

	if is純全帯(p.tiles) {
		yakus = append(yakus, M純全帯)
	}

	// --- 6 ---

	if is清一色(p) {
		yakus = append(yakus, M清一色)
	}

	return yakus
}
