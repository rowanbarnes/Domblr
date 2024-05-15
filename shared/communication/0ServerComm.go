//go:build wasm

package communication

import (
	"encoding/json"
	"errors"
	"syscall/js"
)

type ApiCallback func([]any, error)
type ApiRouter map[string]func(...any) ([]any, error)

var Api ApiRouter

func Call(api string, params ...any) *Promise {
	println("Call()")

	promise := NewPromise()
	//defer close(promise.resultChan)
	// Convert params to JSON
	jsonParams := []byte("[]")
	if len(params) > 0 {
		var err error
		jsonParams, err = json.Marshal(params)
		if err != nil {
			println("promise.Reject()")
			promise.Reject(err)
			return promise
		}
	}

	// Create HTTP request
	Fetch("http://localhost:8080/api/"+api, string(jsonParams), func(body string, err error) {
		println("Fetch callback")
		if err != nil {
			println("promise.Reject()")
			promise.Reject(err)
			return
		}

		// Unmarshal response JSON
		var responseData []any
		err = json.Unmarshal([]byte(body), &responseData)
		if err != nil {
			println("promise.Reject")
			promise.Reject(err)
			return
		}

		// Resolve Promise with response data
		println("promise.Resolve, resp len: ", string(rune(len(responseData))))
		promise.Resolve(responseData)
	})
	return promise
}

func Fetch(url string, body string, then func(body string, err error)) {
	println("Fetch()")

	// Create options object for fetch
	options := js.Global().Get("Object").New()
	options.Set("method", "POST")
	options.Set("body", body)

	// Perform fetch
	awaitable := js.Global().Get("fetch").Invoke(url, options)

	awaitable.Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// Handle success
		if len(args) > 0 {
			rsp := args[0]

			// Check for HTTP error code
			if !rsp.Get("ok").Bool() {
				then("", errors.New("fetch: server responded not OK"))
				return nil
			}

			// Parse JSON
			rsp.Call("text").Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
				// Handle JSON parsing success
				if len(args) > 0 {
					then(args[0].String(), nil)
				} else {
					then("", errors.New("fetch: response.text() failed"))
				}
				return nil
			}))
		} else {
			then("", errors.New("fetch: error"))
		}
		return nil
	}))
}
