package taku

import (
	"testing"

	"github.com/k0kubun/pp"
)

func TestDoCalcPoint(t *testing.T) {
	taku := Taku{}
	c := taku.DoCalcPoint(1)
	pp.Println(c.String())
}
