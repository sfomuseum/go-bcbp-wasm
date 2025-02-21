GOROOT=$(shell go env GOROOT)
GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")

GOROOT=$(shell go env GOROOT)

SERVER_URI=http://localhost:8080

wasmexecjs:
	cp "$(GOROOT)/lib/wasm/wasm_exec.js" www/javascript/

wasmjs:
	GOOS=js GOARCH=wasm \
		go build -mod $(GOMOD) -ldflags="-s -w" \
		-o www/wasm/parse_bcbp.wasm \
		cmd/parse-wasmjs/main.go

# As in: https://github.com/aaronland/go-http-fileserver

debug:
	fileserver \
		-root ./www \
		-server-uri $(SERVER_URI) \
		-mimetype js=text/javascript \
		-mimetype wasm=application/wasm \
		-enable-cors

