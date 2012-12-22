package main

import (
	"./backend"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	port        = flag.Int("port", 8090, "port server is on")
	templateDir = flag.String("templateDir", "templates", "folder where template files will be served from")
	cssDir      = flag.String("css", "css", "folder where css is stored")
	jsDir       = flag.String("js", "js", "folder where css is stored")
)

func FileHandler(fileName, mimeType string, writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", mimeType)
	fileObj, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Could not open %s, %v", fileName, err)
		return
	}
	defer fileObj.Close()
	_, err = io.Copy(writer, fileObj)
	if err != nil {
		log.Printf("Could not read %s, %v", fileName, err)
		io.WriteString(writer, "Error reading file.")
		return
	}

}

func ServeFile(path, fileName, mimeType string) {
	http.Handle(path, http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		FileHandler(fileName, mimeType, writer, request)
	}))
}

func serve404(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, "Not Found")
}

func AddHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		serve404(writer)
		return
	}
}

func main() {
	flag.Parse()
	//connect to mongoDB
	mongoConn := backend.NewMongoDBConn()
	s := mongoConn.Connect("localhost")
	defer s.Close()
	ServeFile("/", *templateDir+"/index.html", "text/html")
	http.Handle("add/", http.HandlerFunc(AddHandler))
	ServeFile("/css/bootstrap.css", *cssDir+"/bootstrap.css", "text/css")
	ServeFile("/js/bootstrap.js", *jsDir+"/bootstrap.js", "application/javascript")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatalf("Could not start web server: %v", err)
		return
	}

}
