package sai

import (
	"fmt"
	"testing"
)

func TestSai(t *testing.T) {
	var result [12]int

	testCount := 10000
	for i := 0; i < testCount; i++ {
		d := DoubleDiceRoll()
		result[d.Sum()-1]++
	}

	for idx, r := range result {
		if idx == 0 {
			continue
		}
		fmt.Printf("%2d: %6d (%8.2f%%)\n", idx+1, r, float32(r)/float32(testCount)*100.0)
	}
}
