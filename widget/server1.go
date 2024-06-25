//go:build !wasm

package widget

import (
	"Domblr/comm"
	"Domblr/util"
	"encoding/json"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
)

type Server struct {
	Addr      string
	ResPath   string
	ApiRouter comm.ApiRouter
}

func (server *Server) RunApp(page *App) {
	buildCtx := &BuildContext{
		Addr:      server.Addr,
		ResPath:   server.ResPath,
		Variables: make(map[int]string),
		ID:        0,
	}
	page.Build(buildCtx)

	// Serve all the files in the res folder on the root
	err := filepath.Walk(server.ResPath, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if err != nil {
			return err
		}

		// Get the relative path to serve the file
		relPath, err := filepath.Rel(server.ResPath, path)
		if err != nil {
			return err
		}

		// Define the HTTP handler for the file
		// TODO: Use Content-Type header for JS/WASM?
		http.HandleFunc("/"+relPath, func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, path)
		})
		return nil
	})
	util.Panic(err)

	// Serve the generated application content
	http.HandleFunc("/api/", server.InvokeRequest)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderCtx := NewRenderContext()
		page.Render(renderCtx)
		_, err := w.Write(renderCtx.Buffer.Bytes())
		if err != nil {
			return
		}
	})
	err = http.ListenAndServe(server.Addr, nil)
	util.Panic(err)
}

func (server *Server) Setup() {
	comm.Api = server.ApiRouter
}

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

	println("Request name: ", requestName)

	// Parse request body to get parameters
	var params []any
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	// Invoke the handler function with parameters
	response, err := handler(params...)
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
