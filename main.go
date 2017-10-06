package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index handler")

	err := indexTemplate.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path[1:]
	log.Println(path)
	// data, err := ioutil.ReadFile(string(path))
	// Reading from bindata.go
	data, err := Asset(path)

	if err == nil {
		var contentType string

		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".svg") {
			contentType = "image/svg+xml"
		} else {
			log.Println("text plain")
			contentType = "text/plain"
		}

		w.Header().Add("Content-Type", contentType)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 Error - " + http.StatusText(404)))
	}
}

var indexTemplate *template.Template

func init() {
	fmt.Println("init")

	indexData, err := Asset("ui/index.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	indexTemplate = template.Must(template.New("index.html").Parse(string(indexData)))
}

func main() {

	http.HandleFunc("/ui/", serveResource)

	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":8080", nil)
}
