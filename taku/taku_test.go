package taku

import (
	"testing"

	"fmt"
)

func TestDoCalcPoint(t *testing.T) {
	taku := Taku{}
	c := taku.DoCalcPoint(1)
	fmt.Println(c.String())
}
