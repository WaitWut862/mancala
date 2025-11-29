package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	StonesPerCell int
	CellsPerSide  int
}

type Game struct {
	Board  []int
	Config *Config
	Turn   int
}

func main() {
	g := &Game{}
	g.initGame()
	g.run()
}

func (g *Game) run() {
	//c := g.Config
	for {
		var i int
		var err error
		for {
			err, i = g.getInput("enter move")
			if err != nil {
				fmt.Println(err)
				continue
			}

			if i >= (len(g.Board) / 2) {
				fmt.Println("move invalid")
				continue
			}
			break
		}
		g.render()
	}
}

func (g *Game) initGame() {
	g.Config = &Config{}
	c := g.Config
	g.Turn = 1

	for {
		err, input := g.getInput("enter desired cells per side")
		if err != nil {
			fmt.Println(err)
			continue
		}
		c.CellsPerSide = input
		break
	}

	totalCells := (c.CellsPerSide * 2) + 2

	g.Board = make([]int, totalCells)

	p1Pot := c.CellsPerSide
	p2Pot := totalCells - 1

	for {
		err, input := g.getInput("enter desired stones per cell")
		if err != nil {
			fmt.Println(err)
			continue
		}
		c.StonesPerCell = input
		break
	}

	for i := 0; i < len(g.Board); i++ {
		if i != p1Pot && i != p2Pot {
			g.Board[i] = c.StonesPerCell
		}
	}
	g.render()
}

func oppIndex(n int) int {
	g := &Game{}
	o := len(g.Board) - 2 - n
	return o
}

func (g *Game) render() {
	g.clearScreen()
	side := g.Config.CellsPerSide
	p1Store := side
	p2Store := (side * 2) + 1

	fmt.Printf("player %d turn\n\n", g.Turn)

	fmt.Print("⌈  ⌉")
	for i := (side * 2); i > side; i-- {
		fmt.Printf("[%2d]", g.Board[i])
	}
	fmt.Print("⌈  ⌉\n")

	fmt.Printf(" %2d ", g.Board[p2Store])

	for i := range side {
		i = i
		fmt.Print("----")
	}

	fmt.Printf(" %2d \n", g.Board[p1Store])

	fmt.Print("⌊  ⌋")
	for i := range side {
		fmt.Printf("[%2d]", g.Board[i])
	}
	fmt.Print("⌊  ⌋\n\n")
}

func (g *Game) getInput(m string) (error, int) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s\n", m)
		i, _ := reader.ReadString('\n')

		i = strings.TrimSpace(i)

		input, err := strconv.Atoi(i)

		if err != nil || input < 0 {
			fmt.Println("input error. try again")
			continue
		}

		return err, input
	}
}

func (g *Game) clearScreen() {
	fmt.Print("\033[H\033[2J")
}
