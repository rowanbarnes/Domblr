//go:build !wasm

package comm

import (
	"Domblr/util"
	"bytes"
	"encoding/base64"
	"github.com/chai2010/webp"
	"image"
	_ "image/png"
	"os"
)

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

func LoadImage(path string) *util.Promise {
	promise := util.NewPromise()
	go func() {
		// Load image
		file, err := os.Open(path)
		promise.Reject(err)
		defer file.Close()
		img, _, err := image.Decode(file)
		promise.Reject(err)

		// Compress image to WebP
		var buf bytes.Buffer
		err = webp.Encode(&buf, img, &webp.Options{Lossless: false, Quality: 0})
		promise.Reject(err)

		// Resolve promise with base64 string
		encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
		promise.Resolve([]any{"data:image/jpeg;base64," + encoded})
	}()

	return promise
}
