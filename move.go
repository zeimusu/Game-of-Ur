package main

import (
	"fmt"
	"math/rand"
)

//Roll simulates the flip of 4 coins, counting heads
func roll() int {
	a := rand.Intn(16)
	switch a {
	case 13:
		return 0
	case 1, 5, 7, 11:
		return 1
	case 0, 2, 4, 8, 10, 14:
		return 2
	case 3, 6, 9, 12:
		return 3
	case 15:
		return 4
	}
	panic("Rolls On Floor, Laughing")
}

func (g *Game) capture(colour string) {
	for i := 4; i < 12; i++ {
		if g.Board.Black[i]+g.Board.White[i] == 2 {
			if colour == "black" {
				g.Board.White[i] = 0
				g.WhiteStones++
			} else {
				g.Board.Black[i] = 0
				g.BlackStones++
			}
		}
	}
}

func (r *Row) checkMove(from, spaces int) error {
	if spaces <= 0 {
		return fmt.Errorf("Illegal move, null or reverse move")
	}
	if from >= 0 && (*r)[from] == 0 {
		return fmt.Errorf("Illegal move, no stone at %v", from)
	}
	if from+spaces == 14 {
		return nil //bearing off is always unblocked
	}
	if from+spaces > 14 || (*r)[from+spaces] > 0 {
		return fmt.Errorf("Illegal move, occupied space at %v", from+spaces)
	}
	return nil
}

func (g *Game) Move(colour string, from, spaces int) error {
	var r *Row
	if colour == "black" {
		r = &g.Board.Black
	} else {
		r = &g.Board.White
	}

	if err := r.checkMove(from, spaces); err != nil {
		return err
	}
	if from+spaces == 14 {
		//bear off
		(*r)[from] = 0
		return nil
	}
	if from == -1 {
		//enter
		(*r)[from+spaces] = 1
		if colour == "black" {
			g.BlackStones--
		} else {
			g.WhiteStones--
		}
		fmt.Println(g.BlackStones)
		return nil
	}
	(*r)[from] = 0
	(*r)[from+spaces] = 1
	return nil
}

func (g *Game) GetAllMoves(roll int, colour string) []int {
	moves := []int{}
	var stones int
	var r Row
	if colour == "black" {
		r = g.Board.Black
		stones = g.BlackStones
	} else {
		r = g.Board.White
		stones = g.WhiteStones
	}

	if stones > 0 && r.checkMove(-1, roll) == nil {
		moves = append(moves, -1)
	}
	for i := range r {
		if r.checkMove(i, roll) == nil {
			moves = append(moves, i)
		}
	}
	return moves
}
