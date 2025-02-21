package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"syscall/js"

	"github.com/sfomuseum/go-bcbp"
)

type ParseResponse struct {
	Raw    string     `json:"raw"`
	Fields *bcbp.BCBP `json:"fields"`
	Month  int        `json:"month"`
	Day    int        `json:"day"`
}

func ParseFunc() js.Func {

	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		bcbp_str := args[0].String()

		logger := slog.Default()
		logger = logger.With("raw", bcbp_str)

		logger.Info("Parse BCBP")

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			resolve := args[0]
			reject := args[1]

			b, err := bcbp.Parse(bcbp_str)

			if err != nil {
				logger.Error("Failed to parse BCBP", "error", err)
				reject.Invoke(fmt.Printf("Failed to parse '%s', %v\n", bcbp_str, err))
				return nil
			}

			rsp := ParseResponse{
				Raw:    bcbp_str,
				Fields: b,
			}

			m, d, err := b.MonthDay()

			if err != nil {
				logger.Error("Failed to derive month/day from date of flight", "error", err)
			} else {
				rsp.Month = m
				rsp.Day = d
			}

			enc, err := json.Marshal(rsp)

			if err != nil {
				logger.Error("Failed to marshal BCBP", "error", err)
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
