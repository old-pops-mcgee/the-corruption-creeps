build:
	GOOS=js GOARCH=wasm go build -tags web -o ./Raylib-Go-Wasm/index/main.wasm .