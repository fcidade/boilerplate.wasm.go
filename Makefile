up:
	go run ./cmd/server

wasm:
	cd ./cmd/wasm/; GOOS=js GOARCH=wasm go build -o ../../assets/json.wasm

out=./assets/
setup:
	mkdir -p $(out)
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" $(out)