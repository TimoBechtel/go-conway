package main

import (
	"syscall/js"

	"github.com/timobechtel/go-conway/internal"
)

var canvas js.Value
var grid internal.Grid

const cellSize = 20

func registerCanvas(args []js.Value) {
	canvas = js.Global().Get("document").Call("getElementById", args[0].String())
}

func registerJSFunctions() {
	js.Global().Set("registerCanvas", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		registerCanvas(args)
		return nil
	}))
	js.Global().Set("run", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		run()
		return nil
	}))
	js.Global().Set("render", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		render()
		return nil
	}))
}

func run() {
	width := canvas.Get("width").Int() / cellSize
	height := canvas.Get("height").Int() / cellSize

	grid = internal.Grid{Height: height, Width: width}

	grid.Init(grid.RandomPositions(height * width / 3))

	grid.TriggerCell()
}

func render() {
	ctx := canvas.Call("getContext", "2d")

	for _, col := range grid.GetCells() {
		for _, cell := range col {
			if cell.IsAlive() {
				ctx.Call("fillRect", cell.GetPosition().X*cellSize, cell.GetPosition().Y*cellSize, cellSize, cellSize)
			} else {
				ctx.Call("clearRect", cell.GetPosition().X*cellSize, cell.GetPosition().Y*cellSize, cellSize, cellSize)
			}
		}
	}
}

func main() {
	c := make(chan bool)
	registerJSFunctions()
	<-c
}
