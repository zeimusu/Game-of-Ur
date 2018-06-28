package main

import (
	"math/rand"
	"time"
)

type Row [14]int

type Board struct {
	Black Row
	White Row
}

type Game struct {
	WhiteStones int
	BlackStones int
	Board       Board
	Pot         int
}

func New() *Game {
	rand.Seed(time.Now().UnixNano())
	return &Game{
		WhiteStones: 7,
		BlackStones: 7,
		Pot:         0,
	}
}
