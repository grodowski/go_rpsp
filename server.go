package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

const (
	QuitMessage string = "QUIT"
	StatsMessage string = "STATS"
	Draw string = "DRAW"
	Win string = "WIN"
	Lose string = "LOSE"
)

type RockPaperScissorsHandler struct {
	// Result matrix for m0 and m1
	ResultMatrix [3][3]string

	// Stats map
	Stats map[string]int
}

func NewRockPaperScissorsHandler() *RockPaperScissorsHandler {
	r := new(RockPaperScissorsHandler)
	//   R P S
	// R D L W
	// P W D L
	// S L W D
	r.ResultMatrix = [3][3]string{
		{Draw, Lose, Win},
		{Win, Draw, Lose},
		{Lose, Win, Draw},
	}
	r.Stats = make(map[string]int)
	r.Stats[Draw] = 0
	r.Stats[Win] = 0
	r.Stats[Lose] = 0
	return r
}

func (r *RockPaperScissorsHandler) Handle(c net.Conn) {
	br := bufio.NewReader(c)
	bw := bufio.NewWriter(c)
	for {
		msg, err := br.ReadString('\r')
		msg = strings.Trim(msg, "\r\f\n")
		if msg == QuitMessage || err == io.EOF {
			c.Close()
			return
		}
		if err != nil {
			fmt.Printf("%#v", err)
		} else {
			bw.WriteString(r.Move(msg))
			bw.WriteString("\n")
			bw.Flush()
		}
	}
}

func (r *RockPaperScissorsHandler) Move(msg string) string {
	if msg == StatsMessage {
		return r.formatStats()
	}
	userMove := NewRpsMove(msg)
	if userMove == InvalidMove() {
		return InvalidMove().String()
	}
	computerMove := NewComputerRpsMove()
	return fmt.Sprintf("%s %s", r.Result(userMove, computerMove), computerMove.String())
}

func (r *RockPaperScissorsHandler) Result(m0 RpsMove, m1 RpsMove) string {
	res := r.ResultMatrix[int(m0)][int(m1)]
	r.Stats[res]++
	return res
}

func (r *RockPaperScissorsHandler) formatStats() string {
	return fmt.Sprintf("W: %d D: %d L: %d", r.Stats[Win], r.Stats[Draw], r.Stats[Lose])
}

func main() {
	ln, err := net.Listen("tcp", ":1983")
	if err != nil {
		fmt.Println(err)
		return
	}
	rps := NewRockPaperScissorsHandler()
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rps.Handle(c)
	}
}
