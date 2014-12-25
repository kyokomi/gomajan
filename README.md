gomajan
=======

[![Build Status](https://drone.io/github.com/kyokomi/gomajan/status.png)](https://drone.io/github.com/kyokomi/gomajan/latest)
[![Coverage Status](https://img.shields.io/coveralls/kyokomi/gomajan.svg)](https://coveralls.io/r/kyokomi/gomajan?branch=master)
[![GoDoc](https://godoc.org/github.com/kyokomi/gomajan?status.svg)](https://godoc.org/github.com/kyokomi/gomajan)

golang 麻雀

```
$ go test -v
=== RUN TestYakuCheck
 =>  役  国士無双
 =>  役  七対子
 =>  雀頭 S9 面子| S1 S2 S3 | S4 S5 S6 | S7 S7 S7 | S1 S2 S3 | 残り なし => 役  清一色
 =>  雀頭 M9 面子| M1 M2 M3 | M4 M5 M6 | M7 M7 M7 | M1 M2 M3 | 残り なし => 役  清一色
 =>  雀頭 P9 面子| P1 P2 P3 | P4 P5 P6 | P7 P7 P7 | P1 P2 P3 | 残り なし => 役  清一色
 =>  雀頭 白 面子| S1 S2 S3 | S4 S5 S6 | S7 S7 S7 | S1 S2 S3 | 残り なし => 役  混一色
 =>  雀頭 東 面子| M1 M2 M3 | M4 M5 M6 | M7 M7 M7 | M1 M2 M3 | 残り なし => 役  混一色
 =>  雀頭 中 面子| P1 P2 P3 | P4 P5 P6 | P7 P7 P7 | P1 P2 P3 | 残り なし => 役  混一色
 =>  雀頭 P8 面子| P2 P3 P4 | M7 M7 M7 | S4 S5 S6 | S7 S7 S7 | 残り なし => 役  断么九
 =>  雀頭 S8 面子| S2 S3 S4 | S6 S6 S6 | 發 發 發 | S2 S3 S4 | 残り なし => 役  緑一色
 =>  雀頭 P5 面子| S2 S3 S4 | 發 發 發 | 中 中 中 | (白白白) | 残り なし => 役  大三元
 =>  雀頭 中 面子| 東 東 東 | 南 南 南 | 發 發 發 | (白白白) | 残り なし => 役  字一色
 =>  雀頭 P1 面子| 東 東 東 | 南 南 南 | 北 北 北 | (西西西) | 残り なし => 役  大四喜
 =>  雀頭 北 面子| P1 P2 P3 | 東 東 東 | 南 南 南 | (西西西) | 残り なし => 役  小四喜
 =>  雀頭 中 面子| P9 P9 P9 | S1 S1 S1 | S5 S5 S5 | 東 東 東 | 残り なし => 役  四暗刻
 =>  雀頭 中 面子| P7 P8 P9 | S1 S1 S1 | S5 S5 S5 | 東 東 東 | 残り なし => 役  三暗刻
 =>  雀頭 S1 面子| P1 P1 P1 | P9 P9 P9 | S9 S9 S9 | (M1M1M1) | 残り なし => 役  清老頭
 =>  雀頭 中 面子| P5 P6 P7 | S2 S3 S4 | 白 白 白 | 發 發 發 | 残り なし => 役  小三元
--- PASS: TestYakuCheck (0.00s)
PASS
ok  	github.com/kyokomi/gomajan	0.008s
```


