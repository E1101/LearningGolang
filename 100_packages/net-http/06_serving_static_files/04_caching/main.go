package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

// Data structure to store
// a file in memory
type cacheFile struct {
	content io.ReadSeeker
	modTime time.Time
}

// Map to store files in memory
var cache map[string]*cacheFile
var mutex = new(sync.RWMutex)

func main() {
	cache = make(map[string]*cacheFile)

	http.HandleFunc("/", serveFiles)
	http.ListenAndServe(":8080", nil)
}

func serveFiles(res http.ResponseWriter, req *http.Request) {
	// Loads from the cache if
	// it’s already populated
	mutex.RLock()
	v, found := cache[req.URL.Path]
	mutex.RUnlock()

	if !found {
		// Maps can’t be written to concurrently
		// or be read while being written to.
		// Using a mutex prevents this from happening
		mutex.Lock()
		defer mutex.Unlock()

		fileName := "./files" + req.URL.Path
		f, err := os.Open(fileName)
		defer f.Close()

		// Handles an error when
		// a file can’t be opened
		if err != nil {
			http.NotFound(res, req)
			return
		}

		var b bytes.Buffer
		_, err = io.Copy(&b, f)
		if err != nil {
			http.NotFound(res, req)
			return
		}

		r := bytes.NewReader(b.Bytes())
		info, _ := f.Stat()
		// Populates the cache object and stores it for later
		v := &cacheFile{
			content: r,
			modTime: info.ModTime(),
		}

		cache[req.URL.Path] = v
	}

	http.ServeContent(res, req, req.URL.Path, v.modTime, v.content)
}
