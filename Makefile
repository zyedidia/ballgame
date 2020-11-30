ballgame: assetfs_vfsdata.go
	go build
assetfs_vfsdata.go:
	go run assets_generate.go
ballgame.wasm: assetfs_vfsdata.go
	GOOS=js GOARCH=wasm go build -o ballgame.wasm
