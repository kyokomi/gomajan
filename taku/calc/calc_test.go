package calc

import (
	"fmt"
	"testing"

	"github.com/kyokomi/gomajan/taku/hora"
	"github.com/kyokomi/gomajan/taku/oyako"
)

func TestCalcPoint(t *testing.T) {
	c := CalcPoint{
		Fu:  30,
		Fan: 2,
	}

	fmt.Println(c.Point(oyako.Oya, hora.Ron))

	c = CalcPoint{
		Fu:  50,
		Fan: 3,
	}

	fmt.Println(c.Point(oyako.Ko, hora.Tsumo))
}
