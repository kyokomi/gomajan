package pai

import "testing"

func TestNum(t *testing.T) {

	type TestCase struct {
		in  MJP
		out int
	}

	testCases := []TestCase{
		TestCase{S1, 1},
		TestCase{S9, 9},
		TestCase{M1, 1},
		TestCase{M9, 9},
		TestCase{P1, 1},
		TestCase{P9, 9},
		TestCase{TON, 0},
		TestCase{PEI, 0},
		TestCase{HAK, 0},
		TestCase{CHN, 0},
	}

	for _, testCase := range testCases {
		if testCase.in.Num() != testCase.out {
			t.Errorf("%s Num 不一致 %d != %d", testCase.in, testCase.in.Num(), testCase.out)
		}
	}
}
