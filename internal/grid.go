package internal

import (
	"math/rand"
	"time"
)

type Grid struct {
	Width  int
	Height int
	cells  [][]*Cell
}

func (g *Grid) Init(cellsAlive []*Position) {
	g.initializeEmptyCells()
	g.setInitialCells(cellsAlive)
	g.connectCells()
}

func (g *Grid) setInitialCells(positions []*Position) {
	for _, position := range positions {
		g.cells[position.X][position.Y] = &Cell{position: position, state: true}
	}
}

func (g *Grid) RandomPositions(count int) []*Position {
	var positions []*Position

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		x := rand.Intn(g.Width)
		y := rand.Intn(g.Height)
		positions = append(positions, &Position{x, y})
	}

	return positions
}

func (g *Grid) TriggerCell() {
	g.cells[0][0].Notify()
}

func (g *Grid) GetCells() [][]*Cell {
	return g.cells
}

func (g *Grid) initializeEmptyCells() {
	for x := 0; x < g.Width; x++ {
		col := []*Cell{}
		for y := 0; y < g.Height; y++ {
			col = append(col, &Cell{position: &Position{x, y}})
		}
		g.cells = append(g.cells, col)
	}
}

func (g *Grid) connectCells() {
	neighborPositions := []Position{{-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}}

	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			for _, position := range neighborPositions {
				cell := g.cells[x][y]
				if !g.withinBounds(&Position{x + position.X, y + position.Y}) {
					continue
				}
				neighbor := g.cells[x+position.X][y+position.Y]
				cell.RegisterNeighbour(neighbor)
			}
		}
	}
}

func (g *Grid) withinBounds(position *Position) bool {
	return position.X >= 0 && position.X < g.Width && position.Y >= 0 && position.Y < g.Height
}
