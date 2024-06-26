package main

import (
	"github.com/iv-ana/go_tu/array"
	"github.com/iv-ana/go_tu/ch_dirc"
	"github.com/iv-ana/go_tu/ch_sync"
	"github.com/iv-ana/go_tu/channels"
	"github.com/iv-ana/go_tu/clo"
	"github.com/iv-ana/go_tu/close_ch"
	"github.com/iv-ana/go_tu/cons_t"
	"github.com/iv-ana/go_tu/funct"
	"github.com/iv-ana/go_tu/goroutine"
	"github.com/iv-ana/go_tu/hw"
	"github.com/iv-ana/go_tu/loops"
	"github.com/iv-ana/go_tu/if_else"
	"github.com/iv-ana/go_tu/inter_face"
	// "github.com/iv-ana/go_tu/loops"
	"github.com/iv-ana/go_tu/maps"
	"github.com/iv-ana/go_tu/methods"
	"github.com/iv-ana/go_tu/rec"
	"github.com/iv-ana/go_tu/sel"
	"github.com/iv-ana/go_tu/st_emb"
	"github.com/iv-ana/go_tu/str"
	"github.com/iv-ana/go_tu/swt"
	"github.com/iv-ana/go_tu/tck"
	"github.com/iv-ana/go_tu/tim"
	"github.com/iv-ana/go_tu/timeout"
	"github.com/iv-ana/go_tu/vair_fun"
	"github.com/iv-ana/go_tu/values"
	"github.com/iv-ana/go_tu/vari"

	

	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main(){
	hw.Hello("")
	funct.Func()
	values.Values()
	array.Array()
	ch_dirc.Ch_dirc()
	ch_sync.Ch_sync()
	channels.Channels()
	clo.Clo()
	close_ch.Close_ch()
	cons_t.Cons_t()
	goroutine.Goroutine()
	if_else.If_else()
	inter_face.Interface()
	loops.Loops()
	maps.Maps()
	methods.Methods()
	rec.Rec()
	sel.Select()
	st_emb.St_emb()
	str.Struct()
	swt.Switch()
	tck.Tck()
	tim.Tim()
	timeout.Timeout()
	vari_fun.Vari_fun()
	vari.Vari()





	// // var dir string

	// // flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	// // flag.Parse()
	// r := mux.NewRouter()

	// // // This will serve files under http://localhost:8000/static/<filename>
	// // r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))/

	// r.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("<h1>hello world</h1>"))
	// })

	// r.HandleFunc("/anas",func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("<h1>hello anas</h1>"))
	// })

	// srv := &http.Server{
	// 	Handler:      r,
	// 	Addr:         "127.0.0.1:8000",
	// 	// Good practice: enforce timeouts for servers you create!
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }

	// log.Fatal(srv.ListenAndServe())
	
}