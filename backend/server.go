package backend

import (
	"Domblr/frontend"
	"Domblr/shared/communication"
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

// InvokeRequest invokes the request with the given name and parameters
func (server *Server) InvokeRequest(w http.ResponseWriter, r *http.Request) {
	println("InvokeRequest")

	// Parse URL to get the request name
	urlParts := strings.Split(r.URL.Path, "/")
	requestName := urlParts[len(urlParts)-1]

	// Check if request name exists in apiRouter
	handler, ok := server.ApiRouter[requestName]
	if !ok {
		http.Error(w, "Request not found", http.StatusNotFound)
		return
	}

	println("requestName: ", requestName)

	// Parse request body to get parameters
	var params []any
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	// Invoke the handler function with parameters
	response, err := handler() //params...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert response to JSON and send it
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
	http.StatusText(http.StatusOK)
	println("jsonResponse: ", jsonResponse)
}

type Server struct {
	Addr      string
	ApiRouter communication.ApiRouter
}

func (server *Server) Setup() {
	communication.Api = server.ApiRouter
}

// Serve begins serving the given application
func (server *Server) Serve(app *frontend.App) {
	http.HandleFunc("/wasm_exec.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		http.ServeFile(w, r, "./res/wasm_exec.js")
	})
	http.HandleFunc("/main.wasm", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/wasm")
		http.ServeFile(w, r, "./res/main.wasm")
	})
	http.HandleFunc("/api/", server.InvokeRequest)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var buffer bytes.Buffer
		app.Page.Render(&buffer)
		_, err := w.Write(buffer.Bytes())
		if err != nil {
			return
		}
	})
	err := http.ListenAndServe(server.Addr, nil)
	if err != nil {
		return
	}
}
