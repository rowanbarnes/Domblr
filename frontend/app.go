package frontend

import (
	"Domblr/shared/widget"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Request struct {
	Api    string `json:"api"`
	Params []any  `json:"params"`
}

type App struct {
	Addr string
	Page *widget.Page
	URL  *url.URL
}

func (app *App) Get(api string, params ...any) error {
	paramSlice := make([]any, len(params))
	copy(paramSlice, params)
	requestData := Request{
		Api:    api,
		Params: paramSlice,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return fmt.Errorf("failed to serialize request data: %v", err)
	}
	println("Sending request: ", requestBody)

	// Send HTTP POST request
	resp, err := http.Post(app.Addr, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()
	return nil
}
