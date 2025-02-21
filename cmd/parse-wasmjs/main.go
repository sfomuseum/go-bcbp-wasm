package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"syscall/js"

	"github.com/sfomuseum/go-bcbp"
)

func ParseFunc() js.Func {

	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		bcbp_str := args[0].String()

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			resolve := args[0]
			reject := args[1]

			b, err := bcbp.Parse(bcbp_str)
			
			if err != nil {
				reject.Invoke(fmt.Printf("Failed to parse '%s', %v\n", bcbp_str, err))
				return nil
			}

			enc, err := json.Marshal(b)

			if err != nil {
				reject.Invoke(fmt.Printf("Failed to marshal result for '%s', %v\n", bcbp_str, err))
				return nil
			}

			resolve.Invoke(string(enc))
			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}

func main() {

	parse_func := ParseFunc()
	defer parse_func.Release()

	js.Global().Set("parse_bcbp", parse_func)

	c := make(chan struct{}, 0)

	slog.Info("WASM parse_bcbp function initialized")
	<-c
	
}
