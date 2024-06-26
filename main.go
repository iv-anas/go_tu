package main

import (
	// "github.com/iv-ana/go_tu/funct"
	// "github.com/iv-ana/go_tu/values"
	// "github.com/iv-ana/go_tu/hw"
	// "flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main(){
	// hw.Hello("")
	// funct.Func()
	// values.Values()
	// var dir string

	// flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	// flag.Parse()
	r := mux.NewRouter()

	// // This will serve files under http://localhost:8000/static/<filename>
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))/

	r.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>hello world</h1>"))
	})

	r.HandleFunc("/anas",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>hello anas</h1>"))
	})

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
	
}