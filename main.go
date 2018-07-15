package main

/*

for the webserver version, there are only a few things that the browser
can do:
   "Please roll the dice and tell me the result"
   "Here is the move for 'black'; show me the board"
or "Tell me the move for 'black', and show me the board"

The response should be in some form that the browser can easily convert
to a displayed board

for learning, how about a simple roll the dice app first.

func main() {
	g := New()
	var from, spaces int
	g.printGame()
	for {
		from, spaces = getMove(g, "black")
		if spaces > 0 {
			g.Move("black", from, spaces)
		} else {
			fmt.Println("Roll zero, press enter")
			fmt.Scanln()
		}
		g.printGame()
		r := roll()
		from = ComputerMove(g, "white", r)
		fmt.Println("Computer rolls", r)
		if spaces > 0 {
			g.Move("white", from, r)
		}
		g.printGame()
		fmt.Println("Computer moves ", from, "Press enter")
		fmt.Scanln()
	}

}
*/
func main() {
	g := New()
	g.server()
}
