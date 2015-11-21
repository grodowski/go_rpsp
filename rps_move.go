package main

import (
	"math/rand"
	"time"
)

var rpsMoves = [3]string{"ROCK", "PAPER", "SCISSORS"}
var source = rand.NewSource(time.Now().UnixNano())
var prg = rand.New(source)

type RpsMove int

func NewRpsMove(str string) RpsMove {
	for i, move := range rpsMoves {
		if move == str {
			return RpsMove(i)
		}
		i++
	}
	return InvalidMove()
}

func InvalidMove() RpsMove {
	return RpsMove(-1)
}

func NewComputerRpsMove() RpsMove {
	return RpsMove(prg.Int31n(3))
}

func (r RpsMove) String() string {
	i := int(r)
	if i > -1 && i < 3 {
		return rpsMoves[i]
	}
	return "INVALID"
}
