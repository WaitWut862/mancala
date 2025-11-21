package main

import (
	"fmt"
)


type Config struct {
	StonesPerCell int
	CellsPerSide int
}

type Game struct {
	Board []int
	Config Config
	Turn int
}


func main() {
	g := &Game{}
	g.initGame()
	fmt.Println(g.Board)
}


func (g *Game) initGame() {
	c := &Config{}
	fmt.Println("enter desired board size (default 6) ")
	fmt.Scanln(&c.CellsPerSide)
	totalCells := (c.CellsPerSide * 2) + 2

	g.Board = make([]int, totalCells)

	p1Pot := c.CellsPerSide
	p2Pot := totalCells - 1


	fmt.Println("enter desired stones per cell (default 4)")
	fmt.Scanln(&c.StonesPerCell)
	
	for i := 0; i < len(g.Board); i++ {
		if i != p1Pot && i != p2Pot {
			g.Board[i] = c.StonesPerCell
		}
	}
}

func oppIndex(n int) int {
	g := &Game{}
	o := len(g.Board) -2 - n
	return o
}
