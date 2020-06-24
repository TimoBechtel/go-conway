# go-conway

Conway's Game of Life experiment, written in Go rendered in the browser using WebAssembly.

Cells are not updated all at once, but instead update themselves asynchronously. Everytime they update, they notify their neighbors.

To start this process, the game first triggers a single cell.

It may not work completely reliably yet.


Part of my 111-Challenge. 
https://twitter.com/hashtag/111Challenge

## Dev
1. compile: `make build`
2. serve `./web` using a web server
