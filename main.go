package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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
	for {
		i := g.getMove()
		g.processMove(i)
		g.render()
	}
}

func (g *Game) processMove(m int) {
	side := g.Config.CellsPerSide
	p1Store := side
	p2Store := (side * 2) + 1
	speed := time.Millisecond * time.Duration(190)
	var oppStore, currentStore int

	if g.Turn == 1 {
		oppStore = p2Store
		currentStore = p1Store
	} else {
		oppStore = p1Store
		currentStore = p2Store
	}

	hand := g.Board[m]
	g.Board[m] = 0
	g.render()
	time.Sleep(speed)

	current := m

	for hand > 0 {
		current++

		if current >= len(g.Board) {
			current = 0
		}

		if current == oppStore {
			continue
		}

		g.Board[current]++
		hand--

		g.render()
		time.Sleep(speed)
	}

	if current != oppStore && current != currentStore {

		if g.Board[g.oppIndex(current)] != 0 && g.Board[current] == 1 {
			g.Board[currentStore] = g.Board[currentStore] + g.Board[g.oppIndex(current)] + g.Board[current]
			g.Board[g.oppIndex(current)] = 0
			g.Board[current] = 0
		}
	}

	if current != currentStore {
		if g.Turn == 1 {
			g.Turn = 2
		} else {
			g.Turn = 1
		}
	}
}

func (g *Game) initGame() {
	g.Config = &Config{}
	c := g.Config
	g.Turn = 1

	g.clearScreen()

	for {
		err, input := g.getInt("enter desired cells per side")
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

	g.clearScreen()
	for {
		err, input := g.getInt("enter desired stones per cell")
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

func (g *Game) oppIndex(n int) int {
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

func (g *Game) getInt(m string) (error, int) {
	for {
		i := g.getInput(m)

		if i == "q" {
			fmt.Println("terminated")
			os.Exit(0)
		}

		input, err := strconv.Atoi(i)

		if err != nil || input < 1 {
			fmt.Println("input error. try again")
			continue
		}

		return err, input
	}
}

func (g *Game) getInput(m string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s\n", m)
	fmt.Print("> ")
	i, _ := reader.ReadString('\n')
	i = strings.TrimSpace(i)

	return i

}

func (g *Game) clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (g *Game) getMove() int {
	var i int
	var err error
	side := g.Config.CellsPerSide
	for {
		err, i = g.getInt("enter move")
		if err != nil {
			fmt.Println(err)
			continue
		}

		if i >= side+1 {
			fmt.Println("move invalid")
			continue
		}
		break
	}
	if g.Turn == 1 {
		return i - 1
	} else {
		return (side * 2) - (i - 1)
	}
}
