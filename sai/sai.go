package sai

import (
	"math/rand"
	"time"
)

type Sai [2]int

func init() {
	rand.Seed(time.Now().UnixNano())
}

func DoubleDiceRoll() Sai {
	return Sai{diceRoll(), diceRoll()}
}

func diceRoll() int {
	return rand.Intn(6) + 1
}

func (s Sai) Sum() int {
	return s[0] + s[1]
}
