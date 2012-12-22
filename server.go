package main

import (
	"flag"
	"fmt"
	"io"
	"labix.org/v2/mgo"
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

type ToDo struct {
	Title       string
	Description string
}

type MongoDBConn struct {
	session *mgo.Session
}

func NewMongoDBConn() *MongoDBConn {
	return &MongoDBConn{}
}

func (m *MongoDBConn) Connect(url string) (*mgo.Session){
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	m.session = session
	return m.session
}

func (m *MongoDBConn) AddToDo(title, description string) (err error) {
	c := m.session.DB("test").C("people")
	err = c.Insert(&ToDo{title, description})
	if err != nil {
		panic(err)
	}
	return nil
}

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

func main() {
	flag.Parse()
	//connect to mongoDB
	mongoConn := NewMongoDBConn()
	s := mongoConn.Connect("localhost")
	defer s.Close()
	
	

	ServeFile("/", *templateDir+"/index.html", "text/html")
	ServeFile("/css/bootstrap.css", *cssDir+"/bootstrap.css", "text/css")
	ServeFile("/js/bootstrap.js", *jsDir+"/bootstrap.js", "application/javascript")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatalf("Could not start web server: %v", err)
		return
	}
}
