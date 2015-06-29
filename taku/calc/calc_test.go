package calc

import (
	"fmt"
	"testing"

	"github.com/kyokomi/gomajan/taku/hora"
	"github.com/kyokomi/gomajan/taku/oyako"
)

func TestCalcPoint(t *testing.T) {
	c := CalcPoint{
		Oyako: oyako.Oya,
		Hora:  hora.Ron,
		Fu:    30,
		Fan:   2,
	}

	fmt.Println(c.Point())

	c = CalcPoint{
		Oyako: oyako.Ko,
		Hora:  hora.Tsumo,
		Fu:    50,
		Fan:   3,
	}

	fmt.Println(c.Point())
}
