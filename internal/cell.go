package internal

import (
	"sync"
	"time"
)

type Neighbor interface {
	Notify()
	GetGeneration() int
	GetNextgeneration() int
	IsAlive() bool
}

type Cell struct {
	position        *Position
	state           bool
	nextState       bool
	generation      int
	nextGeneration  int
	neighbors       []Neighbor
	wg              *sync.WaitGroup
	queuedProcesses int
}

func (c *Cell) RegisterNeighbour(n Neighbor) {
	c.neighbors = append(c.neighbors, n)
}

func (c *Cell) GetPosition() *Position {
	return c.position
}

func (c *Cell) Notify() {

	// make sure only one notification is waiting to be processed
	// prevents memory leakage
	if c.queuedProcesses > 1 {
		return
	}
	c.queuedProcesses++

	currentProcess := c.wg

	thisProcess := &sync.WaitGroup{}
	thisProcess.Add(1)
	c.wg = thisProcess

	if currentProcess != nil {
		currentProcess.Wait()
	}

	// artificial delay, so you actually see something
	time.Sleep(time.Millisecond * 170)

	if c.canUpdate() && !c.isUpdated() {
		c.update()
	}
	if c.canApplyUpdate() && c.isUpdated() {
		c.applyUpdate()
	}

	thisProcess.Done()
	c.queuedProcesses--
}

func (c *Cell) GetGeneration() int {
	return c.generation
}

func (c *Cell) GetNextgeneration() int {
	return c.nextGeneration
}

func (c *Cell) IsAlive() bool {
	return c.state
}

func (c *Cell) update() {
	c.nextGeneration++
	var livingNeighbors int

	for _, n := range c.neighbors {
		if n.IsAlive() {
			livingNeighbors++
		}
	}

	if livingNeighbors < 2 || livingNeighbors > 3 {
		c.nextState = false
	} else if livingNeighbors == 3 {
		c.nextState = true
	}

	c.notifyNeighbours()
}

func (c *Cell) applyUpdate() {
	c.state = c.nextState
	c.generation++
	c.notifyNeighbours()
}

func (c *Cell) isUpdated() bool {
	return c.nextGeneration-c.generation == 1
}

func (c *Cell) canUpdate() bool {
	for _, neighbor := range c.neighbors {
		if neighbor.GetGeneration() != c.GetGeneration() {
			return false
		}
	}
	return true
}

func (c *Cell) canApplyUpdate() bool {
	for _, neighbor := range c.neighbors {
		if neighbor.GetNextgeneration() != c.GetNextgeneration() {
			return false
		}
	}
	return true
}

func (c *Cell) notifyNeighbours() {
	for _, neighbor := range c.neighbors {
		go neighbor.Notify()
	}
}
