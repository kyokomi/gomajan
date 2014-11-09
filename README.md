gomajan
=======

golang 麻雀

```
$ go test -v .
=== RUN TestYakuCheck
P1 P9 M1 M9 S1 S1 S9 東 南 西 北 白 發 中
 =>  [国士無双]
P5 P5 P6 P6 M3 M3 M4 M4 S1 S1 S2 S2 東 東
雀頭  P5
面子  []
残り
 =>  [七対子]
S1 S1 S2 S2 S3 S3 S4 S5 S6 S7 S7 S7 S9 S9
雀頭  S9
面子  [[S1 S2 S3] [S4 S5 S6] [S7 S7 S7] [S1 S2 S3]]
残り
 =>  [清一色]
M1 M1 M2 M2 M3 M3 M4 M5 M6 M7 M7 M7 M9 M9
雀頭  M9
面子  [[M1 M2 M3] [M4 M5 M6] [M7 M7 M7] [M1 M2 M3]]
残り
 =>  [清一色]
P1 P1 P2 P2 P3 P3 P4 P5 P6 P7 P7 P7 P9 P9
雀頭  P9
面子  [[P1 P2 P3] [P4 P5 P6] [P7 P7 P7] [P1 P2 P3]]
残り
 =>  [清一色]
S1 S1 S2 S2 S3 S3 S4 S5 S6 S7 S7 S7 白 白
雀頭  白
面子  [[S1 S2 S3] [S4 S5 S6] [S7 S7 S7] [S1 S2 S3]]
残り
 =>  [混一色]
M1 M1 M2 M2 M3 M3 M4 M5 M6 M7 M7 M7 東 東
雀頭  東
面子  [[M1 M2 M3] [M4 M5 M6] [M7 M7 M7] [M1 M2 M3]]
残り
 =>  [混一色]
P1 P1 P2 P2 P3 P3 P4 P5 P6 P7 P7 P7 中 中
雀頭  中
面子  [[P1 P2 P3] [P4 P5 P6] [P7 P7 P7] [P1 P2 P3]]
残り
 =>  [混一色]
--- PASS: TestYakuCheck (0.00 seconds)
PASS
ok  	github.com/kyokomi/gomajan	0.009s
```


