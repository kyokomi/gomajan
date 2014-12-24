package mjp

type Yaku struct {
	// 翻数
	Fan int
	// 名前
	Name string
	// 判定ロジック
	CheckFunc func() bool
}

var (
	M_国士無双 = Yaku{Fan: 13, Name: "国士無双"}
	M_大三元  = Yaku{Fan: 13, Name: "大三元"}
	M_字一色  = Yaku{Fan: 13, Name: "字一色"}
	M_大四喜  = Yaku{Fan: 13, Name: "大四喜"}
	M_小四喜  = Yaku{Fan: 13, Name: "小四喜"}
	M_四暗刻  = Yaku{Fan: 13, Name: "四暗刻"}
	M_清老頭  = Yaku{Fan: 13, Name: "清老頭"}
	M_緑一色  = Yaku{Fan: 13, Name: "緑一色"}
	M_七対子  = Yaku{Fan: 2, Name: "七対子"}
	M_断么九  = Yaku{Fan: 1, Name: "断么九"}
	M_一気通貫 = Yaku{Fan: 2, Name: "一気通貫"}
	M_三暗刻  = Yaku{Fan: 2, Name: "三暗刻"}
	M_小三元  = Yaku{Fan: 2, Name: "小三元"}
	M_混一色  = Yaku{Fan: 3, Name: "混一色"}
	M_純全帯  = Yaku{Fan: 3, Name: "純全帯"}
	M_清一色  = Yaku{Fan: 3, Name: "清一色"}
)

func (p Player) yakuManCheck() []Yaku {
	res := make([]Yaku, 0)

	// 国士無双判定
	if is国士無双(p.tiles) {
		res = append(res, M_国士無双)
		// 国士無双と他の組み合わせはないので終わり
		return res
	}

	// 大三元
	if p.is大三元() {
		res = append(res, M_大三元)
	}

	// 字一色
	if p.is字一色() {
		res = append(res, M_字一色)
	}

	// 大四喜 or 小四喜
	if p.is大四喜() {
		res = append(res, M_大四喜)
	} else if p.is小四喜() {
		res = append(res, M_小四喜)
	}

	// 四暗刻
	if is四暗刻(p.tiles) {
		res = append(res, M_四暗刻)
	}

	// TODO: 四槓子

	// 清老頭
	if p.is清老頭() {
		res = append(res, M_清老頭)
	}

	// 緑一色
	if is緑一色(p.tiles) {
		res = append(res, M_緑一色)
	}

	return res
}

func (p Player) yakuCheck() []Yaku {
	yakus := make([]Yaku, 0)

	if is七対子(p.tiles) {
		yakus = append(yakus, M_七対子)
	}

	// --- 1翻 ---

	// TODO: リーチ
	// TODO: 一発
	// TODO: 門前清自摸

	// TODO: 平和

	if is断么九(p.tiles) {
		yakus = append(yakus, M_断么九)
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
		yakus = append(yakus, M_一気通貫)
	}

	// TODO: 対々和

	if is三暗刻(p.tiles) {
		// TODO: 食い下がりある
		yakus = append(yakus, M_三暗刻)
	}

	// TODO: 三槓子

	if is小三元(p.tiles) {
		yakus = append(yakus, M_小三元)
	}

	// --- 3 ---

	// TODO: 混老頭
	// 食い下がり

	if is混一色(p.tiles) {
		yakus = append(yakus, M_混一色)
	}

	// TODO: 二盃口

	if is純全帯(p.tiles) {
		yakus = append(yakus, M_純全帯)
	}

	// --- 6 ---

	if is清一色(p.tiles) {
		yakus = append(yakus, M_清一色)
	}

	return yakus
}
