package main

import (
	"errors"
	"math/rand"
	"time"
)

const invalidMove int = -1

var rpsMoves = [3]string{"ROCK", "PAPER", "SCISSORS"}
var source = rand.NewSource(time.Now().UnixNano())
var prg = rand.New(source)


type RpsMove int

func NewRpsMove(str string) (RpsMove, error) {
	for i, move := range rpsMoves {
		if move == str {
			return RpsMove(i), nil
		}
		i++
	}
	return -1, errors.New("INVALID")
}

func NewComputerRpsMove() RpsMove {
	return RpsMove(prg.Int31n(3))
}

func (r RpsMove) String() string {
	return rpsMoves[int(r)]
}
