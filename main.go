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
	Config *Config
	Turn int
}


func main() {
	g := &Game{}
	g.initGame()
	g.render()
}


func (g *Game) initGame() {
	g.Config = &Config{}
	c := g.Config
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


func (g *Game) render() {
	side := g.Config.CellsPerSide
	p1Store := side
	p2Store := (side * 2) +1

	fmt.Print("⌈  ⌉")
	for i := (side * 2); i > side; i-- {
		fmt.Printf("[%2d]", g.Board[i])
	}
	fmt.Print("⌈  ⌉\n")

	fmt.Printf(" %2d ", g.Board[p2Store])

	for i := 0; i < side; i++ {
		fmt.Print("----")
	}

	fmt.Printf(" %2d \n", g.Board[p1Store])

	fmt.Print("⌊  ⌋")
	for i := 0; i < side; i++ {
		fmt.Printf("[%2d]", g.Board[i])
	}
	fmt.Print("⌊  ⌋\n")
}


