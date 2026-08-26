all: clapclap.ndp

clapclap.ndp: plugin.wasm
	zip -j clapclap.ndp manifest.json plugin.wasm

plugin.wasm: *.go *.json *.mod
	tinygo build -o plugin.wasm -target wasip1 -gc=leaking -buildmode=c-shared .

clean:
	rm -r plugin.wasm clapclap.ndp