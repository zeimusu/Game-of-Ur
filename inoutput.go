package main

import "fmt"

func (g *Game) printGame() {
	fmt.Println("White stones left", g.WhiteStones)
	fmt.Println(g.drawGame())
	fmt.Println("Black stones left", g.BlackStones)
	fmt.Println("Pot", g.Pot)
}

func (g *Game) drawGame() string {
	grid := `
+-v-+---         ---+---+---+-v-+
> %s < %s |       | %s | %s | %s > %s <
+-^-+---+---+---+-v-+---+---+-^-+
| %s | %s | %s | %s > %s < %s | %s | %s |
+-v-+---+---+---+-^-+---+---+-v-+
> %s < %s |       | %s | %s | %s > %s <
+-^-+---         ---+---+---+-^-+
`
	stones := []interface{}{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}

	for i, x := range g.Board.White {
		if x == 1 {
			switch true {
			case i < 4:
				stones[i+2] = "O"
			case i < 12:
				stones[17-i] = "O"
			case i >= 12:
				stones[i-12] = "O"
			}
		}
	}
	for i, x := range g.Board.Black {
		if x == 1 {
			switch true {
			case i < 4:
				stones[i+16] = "#"
			case i < 12:
				stones[17-i] = "#"
			case i >= 12:
				stones[i+2] = "#"
			}
		}
	}
	return fmt.Sprintf(grid, stones...)
}

func contains(x int, s []int) bool {
	for _, y := range s {
		if y == x {
			return true
		}
	}
	return false
}

func getMove(g *Game, colour string) (int, int) {
	r := roll()
	fmt.Println("Roll is", r)
	options := g.GetAllMoves(r, colour)
	if len(options) == 0 {
		return -2, -1
	}
	fmt.Println("You can move", options)
	m := 0
	for {
		fmt.Scan(&m)
		if contains(m, options) {
			return m, r
		}
		fmt.Println("Try again")
	}
	panic("No valid move")
}
