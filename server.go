package main

import (
	"./backend"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"log"
	"net/http"
)

const templateDir = "templates"

var (
	port = flag.Int("port", 8090, "port server is on")
	mongoConn *backend.MongoDBConn
)


func serve404(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	writer.WriteHeader(http.StatusNotFound)
	io.WriteString(writer, "Not Found")
}

func AddHandler(writer http.ResponseWriter, request *http.Request) {
	log.Printf("serving %v %v", request.Method, request.URL.Path[1:])
	if request.Method != "POST" {
		serve404(writer)
		return
	}
	title := request.FormValue("title")
	description := request.FormValue("description")
	log.Printf(" title description %v %v", title, description)
	err := mongoConn.AddToDo(title, description)
	if err != nil {
		panic(err)
		fmt.Fprintln(writer, "fail")
		return
	}
	fmt.Fprintln(writer, "success")
}
func add(x, y int) int {
	return x + y
}

func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	results := mongoConn.ListToDo()
	funcs := template.FuncMap{"add": add} 
  temp := template.Must(template.New("index.html").Funcs(funcs).ParseFiles(templateDir + "/index.html"))
	temp.Execute(writer, results)
}

func main() {
	flag.Parse()
	//connect to mongoDB
	mongoConn = backend.NewMongoDBConn()
	_ = mongoConn.Connect("localhost")
	defer mongoConn.Stop()
	log.Printf("Starting server on %v", *port)
  r := mux.NewRouter()
  r.HandleFunc("/", IndexHandler)
  r.HandleFunc("/add/", AddHandler)
  r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", 
          http.FileServer(http.Dir("css/"))))
  r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", 
                http.FileServer(http.Dir("js/"))))
  err := http.ListenAndServe(":8090", r)
  if err != nil {
      log.Fatalf("Could not start web server: %v", err)
  }
}
