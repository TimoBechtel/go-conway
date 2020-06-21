package internal

import (
	"sync"
	"testing"
)

func TestCellHasPosition(t *testing.T) {
	cell := Cell{position: &Position{1, 5}}
	if cell.GetPosition().X != 1 || cell.GetPosition().Y != 5 {
		t.Error("Cell has wrong position")
	}
}

func TestCellCanRegisterNeighbors(t *testing.T) {
	cell := Cell{}
	cell2 := Cell{}

	cell.RegisterNeighbour(&cell2)

	if len(cell.neighbors) < 1 {
		t.Error("Cell didn't register neighbor")
	}

	t.Log(cell.neighbors)
}

func TestCellUpdatesNextGenerationWhenUpdated(t *testing.T) {
	cell := Cell{}
	cell.update()

	if cell.GetNextgeneration() != 1 {
		t.Error("Cell did not udpate generation when updating")
		t.Log(cell.GetNextgeneration())
	}
}

func TestCellUpdatesCurrentGenerationWhenUpdateIsApplied(t *testing.T) {
	cell := Cell{}
	cell.update()

	if cell.GetGeneration() > 0 {
		t.Error("Cell has wrong generation")
		t.Log(cell.GetNextgeneration())
	}

	cell.applyUpdate()
	if cell.GetGeneration() != 1 {
		t.Error("Cell did not udpate generation when udpate was applied")
		t.Log(cell.GetGeneration())
	}
}

func TestCellKnowsWhenItNeedsAnUpdate(t *testing.T) {
	cell := Cell{}
	cell.update()

	if cell.isUpdated() == false {
		t.Error("Reports false update need")
	}

	cell.applyUpdate()
	if cell.isUpdated() == true {
		t.Error("Reports false update need after applying update")
	}
}

func TestCellDoesNotUpdateBeforeUpdateIsApplied(t *testing.T) {
	cell := Cell{}

	cell.Notify()
	cell.Notify()
	cell.Notify()
	cell.Notify()

	if cell.nextGeneration-cell.generation > 1 {
		t.Error("Cell updated more than once")
		t.Log(cell)
	}

	cell.applyUpdate()

	if cell.generation != cell.nextGeneration+1 {
		t.Error("Cell did not apply update")
		t.Log(cell)
	}
}

func TestCellOnlyUpdatesWhenAllNeighborsAreOnSameGeneration(t *testing.T) {
	cell := Cell{generation: 0}
	neighbor := Cell{generation: 0}
	neighbor2 := Cell{generation: 1}
	cell.RegisterNeighbour(&neighbor)
	cell.RegisterNeighbour(&neighbor2)

	cell.Notify()

	if cell.nextGeneration > 0 {
		t.Error("Cell updated when it shouldn't")
		t.Log(cell.nextGeneration)
	}

	neighbor2.generation = 0

	cell.Notify()

	if cell.nextGeneration != 1 {
		t.Error("Cell did't update when it should")
		t.Log(cell.nextGeneration)
	}
	t.Log(neighbor)
}

func TestCellOnlyAppliesUpdateWhenAllNeighborsAreUpdated(t *testing.T) {
	cell := &Cell{}
	neighbor := &Cell{}
	neighbor2 := &Cell{}

	cell.RegisterNeighbour(neighbor)
	cell.RegisterNeighbour(neighbor2)

	neighbor.update()

	t.Log("cell", cell)
	t.Log("neighbor", neighbor)
	t.Log("neighbor2", neighbor2)

	cell.Notify()
	if cell.generation > 0 {
		t.Error("Cell applied update when it shouldn't")
	}

	neighbor2.update()
	cell.Notify()

	if cell.generation != 1 {
		t.Error("Cell didn't apply update when it should have")
	}
}

type cellMock struct {
	wg             *sync.WaitGroup
	cell           *Cell
	generation     int
	nextGeneration int
}

func (c cellMock) Notify() {
	c.cell.Notify()
	c.wg.Done()
}

func (c cellMock) GetGeneration() int {
	return c.generation
}

func (c cellMock) GetNextgeneration() int {
	return c.nextGeneration
}

func (c cellMock) IsAlive() bool {
	return true
}

func TestCellNotifiesNeighborsWhenUpdated(t *testing.T) {
	cell := Cell{}
	neighbor := Cell{}

	wg := sync.WaitGroup{}
	wg.Add(1)

	mock := cellMock{wg: &wg, cell: &neighbor}

	cell.RegisterNeighbour(&mock)

	cell.update()

	wg.Wait()

	if neighbor.nextGeneration != cell.nextGeneration {
		t.Error("Cell didn't update neighbor")
		t.Log(neighbor)
		t.Log(cell)
	}
}

func TestCellNotifiesNeighborsWhenUpdateIsApplied(t *testing.T) {
	cell := Cell{generation: 0, nextGeneration: 1}
	neighbor := Cell{generation: 0, nextGeneration: 1}

	wg := sync.WaitGroup{}
	wg.Add(1)

	mock := cellMock{wg: &wg, cell: &neighbor}

	cell.RegisterNeighbour(mock)

	cell.applyUpdate()

	wg.Wait()

	if neighbor.generation != 1 {
		t.Error("Cell didn't update neighbor")
		t.Log(neighbor)
		t.Log(cell)
	}
}

func TestCellIsResurrectedWhenHavingThreeLivingNeighbors(t *testing.T) {
	cell := &Cell{state: false}
	neighbor := &Cell{state: true}
	neighbor2 := &Cell{state: true}
	neighbor3 := &Cell{state: true}

	cell.RegisterNeighbour(neighbor)
	cell.RegisterNeighbour(neighbor2)
	cell.RegisterNeighbour(neighbor3)

	cell.update()

	if cell.nextState != true {
		t.Error("Cell was not resurrected")
	}
}

func TestCellDiesWithMoreThanThreeLivingNeighbors(t *testing.T) {
	cell := &Cell{state: true, nextState: true}
	neighbor := &Cell{state: true}
	neighbor2 := &Cell{state: true}
	neighbor3 := &Cell{state: true}
	neighbor4 := &Cell{state: true}

	cell.RegisterNeighbour(neighbor)
	cell.RegisterNeighbour(neighbor2)
	cell.RegisterNeighbour(neighbor3)
	cell.RegisterNeighbour(neighbor4)

	cell.update()

	if cell.nextState != false {
		t.Error("Cell did not die")
		t.Log(cell)
	}
}

func TestCellDiesWithLessThanTwoLivingNeighbors(t *testing.T) {
	cell := &Cell{state: true, nextState: true}
	neighbor := &Cell{state: true}
	neighbor2 := &Cell{state: false}
	neighbor3 := &Cell{state: false}
	neighbor4 := &Cell{state: false}

	cell.RegisterNeighbour(neighbor)
	cell.RegisterNeighbour(neighbor2)
	cell.RegisterNeighbour(neighbor3)
	cell.RegisterNeighbour(neighbor4)

	cell.update()

	if cell.nextState != false {
		t.Error("Cell did not die")
		t.Log(cell)
	}
}
