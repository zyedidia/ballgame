ballgame:
	go build
ballgame.wasm:
	GOOS=js GOARCH=wasm go build -o ballgame.wasm
