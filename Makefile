all: main.wasm serve

main.wasm: ${wildcard ./**/*.go} main.go
	GOOS=js GOARCH=wasm go generate
	GOOS=js GOARCH=wasm go build -o main.wasm main.go
serve:
	go run ./server/main.go

clean:
	rm -f *.wasm