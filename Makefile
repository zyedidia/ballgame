ballgame: require
	go build
assetfs_vfsdata.go: require
	go run assets_generate.go
ballgame.wasm: require
	GOOS=js GOARCH=wasm go build -o ballgame.wasm
.PHONY: require
