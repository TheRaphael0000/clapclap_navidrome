all: clap_dht.ndp

clap_dht.ndp: plugin.wasm
	zip -j clap_dht.ndp manifest.json plugin.wasm

plugin.wasm: *.go *.json *.mod
	tinygo build -o plugin.wasm -target wasip1 -gc=leaking -buildmode=c-shared .

clean:
	rm -r plugin.wasm clap_dht.ndp