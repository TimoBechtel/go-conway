package internal

import "testing"

func TestGridInitializesWithEnoughCells(t *testing.T) {
	grid := Grid{Height: 2, Width: 6}

	grid.Init([]*Position{})

	if len(grid.cells) != 6 || len(grid.cells[0]) != 2 {
		t.Error("Grid did not create expected number of cells")
		t.Log(len(grid.cells))
		t.Log(len(grid.cells[0]))
	}
}

func TestGridCanCreateRandomPositionsWithinGrid(t *testing.T) {
	grid := Grid{Height: 5, Width: 5}

	positions := grid.RandomPositions(5)

	if len(positions) != 5 {
		t.Error("Not the exptected number of positions")
	}

	for _, pos := range positions {
		if pos.X > grid.Width || pos.X < 0 || pos.Y > grid.Height || pos.Y < 0 {
			t.Error("position is out of grid")
			t.Log(pos)
		}
	}
}

func TestInitializesCellsAtPositions(t *testing.T) {
	grid := Grid{Height: 5, Width: 6}
	positions := []*Position{{X: 0, Y: 3}}
	grid.Init(positions)

	if grid.cells[0][3].state != true {
		t.Error("Grid did not initialize expected cell")
	}
}

// not tested, as this triggers the loop
// func TestGridConnectsAllCells(t *testing.T) {
// 	grid := Grid{Height: 1, Width: 2}
// 	grid.Init([]*Position{})
// 	grid.TriggerCell()

// 	for _, col := range grid.cells {
// 		for _, cell := range col {
// 			t.Log(cell)
// 			if cell.nextGeneration != 1 {
// 				t.Error("Cell did not update when it should")
// 			}
// 		}
// 	}
// }
