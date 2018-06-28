package main

import (
	"fmt"
	"math/rand"
)

func ComputerMove(g *Game, colour string) (int, int) {
	r := roll()
	options := g.GetAllMoves(r, colour)
	//fmt.Println("Options", options, "roll", r)
	if len(options) == 0 {
		return -2, 0
	}
	if len(options) == 1 {
		return options[0], r
	}
	moveScores := make([]float64, len(options))
	//use some ai to choose the best move.
	for i := range options {
		//fmt.Println("testing", options[i])
		myBoard, theirBoard, myStones, theirStones := trialMove(g, colour, options[i], r)
		moveScores[i] = score(myBoard, theirBoard, myStones, theirStones)
	}
	//fmt.Println("options", options)
	//fmt.Println("scores ", moveScores)
	bestMove := pickBest(options, moveScores)
	//fmt.Println("choice", bestMove)
	return bestMove, r
}

func copyRow(source Row, destination *Row) {
	for i := range source {
		(*destination)[i] = source[i]
	}
}

func trialMove(g *Game, colour string, from, spaces int) (Row, Row, int, int) {
	myBoard := Row{}
	theirBoard := Row{}
	var myStones, theirStones int
	if colour == "black" {
		copyRow(g.Board.Black, &myBoard)
		copyRow(g.Board.White, &theirBoard)
		myStones = g.BlackStones
		theirStones = g.WhiteStones
	} else {
		copyRow(g.Board.White, &myBoard)
		copyRow(g.Board.Black, &theirBoard)
		myStones = g.WhiteStones
		theirStones = g.BlackStones
	}
	//fmt.Println("myboard", myBoard)
	//fmt.Println("theirboard", theirBoard)
	//fmt.Println("my, their stones", myStones, theirStones)
	if from == -1 {
		if myBoard[from+spaces] == 1 {
			panic("can't move stone on")
		}
		myBoard[from+spaces] = 1
		myStones--
	}
	if from+spaces == 14 {
		if myBoard[from] == 0 {
			panic("can't remove stone")
		}
		myBoard[from] = 0
	}
	if from >= 0 && from+spaces < 14 {
		if myBoard[from] == 0 {
			panic(fmt.Sprintf("no stone at that position %v", from))
		}
		if myBoard[from+spaces] == 1 {
			panic("cant move to that position")
		}
		myBoard[from] = 0
		myBoard[from+spaces] = 1
	}
	if from+spaces < 14 && theirBoard[from+spaces] == 1 {
		theirBoard[from+spaces] = 0
		theirStones++
	}
	return myBoard, theirBoard, myStones, theirStones
}

func max(s []float64) (out float64) {
	for _, x := range s {
		if x > out {
			out = x
		}
	}
	return
}

func pickBest(options []int, scores []float64) int {
	if len(options) != len(scores) {
		panic(fmt.Sprintf("Lengths in pick best not equal, options %v/ scores %v", len(options), len(scores)))
	}
	m := max(scores)
	filteredOptions := make([]int, 0, len(options))
	for i := range options {
		if scores[i] == m {
			filteredOptions = append(filteredOptions, options[i])
		}
	}

	return filteredOptions[rand.Intn(len(filteredOptions))]
}

func score(myBoard, theirBoard Row, myStones, theirStones int) float64 {
	score := 1000.0                               // So the score stays positive
	score = score + float64(theirStones-myStones) //we don't want stones in hand

	dangerFactor := 0.25
	danger := make([]int, len(theirBoard))
	for i := range theirBoard {
		if theirBoard[i] == 1 {
			if i+1 < len(danger) {
				danger[i+1] += 4
			}
			if i+2 < len(danger) {
				danger[i+2] += 6
			}
			if i+3 < len(danger) {
				danger[i+3] += 4
			}
			if i+4 < len(danger) {
				danger[i+4] += 1
			} // 4Cr
		}
	}
	danger[0], danger[1], danger[2], danger[3], danger[12], danger[13] = 0, 0, 0, 0, 0, 0
	totalDanger := 0
	for i := range myBoard {
		totalDanger += myBoard[i] * danger[i]
	}
	score -= dangerFactor * float64(totalDanger)
	return score
}
