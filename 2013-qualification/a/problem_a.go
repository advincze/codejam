package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var fileName = flag.String("file", "", "4x4 tic tac toe file")

func main() {
	flag.Parse()
	f, err := os.Open(*fileName)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	c, _ := r.ReadString('\n')

	k := 1
	for count, _ := strconv.Atoi(c[:len(c)-1]); count > 0; count-- {
		var res GameState = draw
		var cross1 [4]byte
		var cross2 [4]byte
		var vert [4][4]byte
		for i := 0; i < 4; i++ {
			var horizontal [4]byte
			for j := 0; j < 4; j++ {
				b, _ := r.ReadByte()
				horizontal[j] = b
				vert[j][i] = b
				if i == j {
					cross1[i] = b
				} else if i+j == 3 {
					cross2[i] = b
				}
			}
			res = mergeresult(res, checkline(horizontal))
			r.ReadBytes('\n')
		}

		res = mergeresult(res, checkline(cross1))
		res = mergeresult(res, checkline(cross2))

		for i := 0; i < 4; i++ {
			res = mergeresult(res, checkline(vert[i]))
		}

		fmt.Printf("Case #%v: %v\n", k, res)
		k++
		r.ReadBytes('\n')
	}

}

func mergeresult(res, newres GameState) GameState {
	switch res {
	case xWins, oWins:
		return res
	case draw:
		switch newres {
		case draw:
			return draw
		default:
			return newres
		}
	case notFinished:
		switch newres {
		case xWins, oWins:
			return newres
		default:
			return notFinished
		}

	}
	panic("")
}

func checkline(line [4]byte) GameState {
	var gameState GameState
	for _, b := range line {

		if b == empty {
			return notFinished
		}
		switch gameState {
		case none:
			switch b {
			case x:
				gameState = xWins
			case o:
				gameState = oWins
			}
		case xWins:
			if b == o {
				gameState = draw
			}
		case oWins:
			if b == x {
				gameState = draw
			}
		}

	}

	return gameState

}

const (
	x     byte = 'X'
	o     byte = 'O'
	empty byte = '.'
	t     byte = 'T'
)

type GameState int

const (
	none GameState = iota
	xWins
	oWins
	draw
	notFinished
)

func (gs GameState) String() string {
	switch gs {
	case none:
		return "none"
	case xWins:
		return "X won"
	case oWins:
		return "O won"
	case draw:
		return "Draw"
	case notFinished:
		return "Game has not completed"
	}
	panic("")
}
