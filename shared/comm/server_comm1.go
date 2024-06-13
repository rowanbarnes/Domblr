//go:build !wasm

package comm

import "Domblr/util"

type ApiCallback func([]any, error)
type ApiRouter map[string]func(...any) ([]any, error)

var Api ApiRouter

func Call(api string, params ...any) *util.Promise {
	promise := util.NewPromise()
	resp, err := Api[api](params)
	if err != nil {
		promise.Reject(err)
	} else {
		promise.Resolve(resp)
	}
	return promise
}

func Fetch(url string, body string, then func(body string, err error)) {
}
