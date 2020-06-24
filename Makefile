build: cmd/main.go
	GOOS=js GOARCH=wasm go build -o web/goconway.wasm cmd/main.go