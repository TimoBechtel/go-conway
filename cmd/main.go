package main

import (
	"time"

	tm "github.com/buger/goterm"

	"github.com/timobechtel/go-conway/internal"
)

func main() {

	width := 70
	height := 30

	grid := internal.Grid{Height: height, Width: width}

	grid.Init(grid.RandomPositions(height * width / 3))

	grid.TriggerCell()

	for {
		tm.Clear()
		tm.MoveCursor(1, 1)

		for _, col := range grid.GetCells() {
			for _, cell := range col {
				if cell.IsAlive() {
					tm.Print(tm.MoveTo(tm.Background(tm.Color(" ", tm.BLACK), tm.WHITE), cell.GetPosition().X, cell.GetPosition().Y))
				}
			}
		}

		tm.Flush()
		time.Sleep(time.Millisecond * 150)
	}
}
