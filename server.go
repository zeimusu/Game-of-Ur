package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func (g *Game) handleMove(w http.ResponseWriter, r *http.Request) {
	// move?from=a&spaces=b
	q := r.URL.Query()
	_ = q
	spacesList, ok := r.URL.Query()["spaces"]
	if !ok {
		fmt.Println("no spaces key")
		return
	}
	fromList, ok := r.URL.Query()["from"]
	if !ok {
		fmt.Println("no from key")
		return
	}
	if len(spacesList) == 0 {
		fmt.Println("Empty spaces list")
		return
	}
	if len(fromList) == 0 {
		fmt.Println("Empty from list")
		return
	}
	from, err := strconv.Atoi(fromList[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	spaces, err := strconv.Atoi(spacesList[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	colour := "black"
	g.Move(colour, from, spaces)
	if j, err := json.Marshal(g); err == nil {
		w.Write(j)
	}
}

func (g *Game) handleComputerMove(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Computer move")
	spacesList, ok := r.URL.Query()["spaces"]
	if !ok {
		fmt.Println("no spaces key")
		return
	}
	if len(spacesList) == 0 {
		fmt.Println("Empty spaces list")
		return
	}
	spaces, err := strconv.Atoi(spacesList[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	colour := "white"
	from := ComputerMove(g, colour, spaces)
	g.Move(colour, from, spaces)
	if j, err := json.Marshal(g); err == nil {
		w.Write(j)
	}
}

func handleUr(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
	}
	http.ServeFile(w, r, "jstest.html")
}

func handleJS(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "urscript.js")
}

func handleCSS(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "urstyle.css")
}
func handleBG(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "RoyalGameofUr.png")
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}

func (g *Game) server() {
	fmt.Println("Ur server")
	/*
		Create a webserver on port 8080
		the root web page will just have a text div and a button "roll"
		when the button is clicked, and ajax request will be made and
		a number? will be sent. The web page will present this as text
	*/
	http.HandleFunc("/", handleUr)
	http.HandleFunc("/move", g.handleMove)
	http.HandleFunc("/computer", g.handleComputerMove)
	http.HandleFunc("/urscript", handleJS)
	http.HandleFunc("/urstyle", handleCSS)
	http.HandleFunc("/RoyalGameofUr", handleBG)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
