package main

import "fmt"

func main() {
	g := New()
	var from, spaces int
	for {
		g.printGame()
		from, spaces = getMove(g, "black")
		if spaces > 0 {
			g.Move("black", from, spaces)
			g.capture("black")
		}
		g.printGame()
		from, spaces = ComputerMove(g, "white")
		fmt.Println("Computer rolls", spaces)
		if spaces > 0 {
			g.Move("white", from, spaces)
			g.capture("white")
		}
	}

}
