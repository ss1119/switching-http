package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	_ "net/http/pprof"

	"github.com/lucas-clemente/quic-go/http3"
)

type binds []string

func (b binds) String() string {
	return strings.Join(b, ",")
}

func (b *binds) Set(v string) error {
	*b = strings.Split(v, ",")
	return nil
}

// Size is needed by the /demo/upload handler to determine the size of the uploaded file
type Size interface {
	Size() int64
}

func setupHandler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	mux.HandleFunc("/", IndexHandler)

	return mux
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadDir("assets")
	var templates = template.Must(template.ParseFiles("templates/index.html"))
	if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
		log.Fatalln("Unable to execute template.")
	}
}

func main() {
	var crtPath = "../live/localhost111919.ml/fullchain.pem"
	var keyPath = "../live/localhost111919.ml/privkey.pem"
	handler := setupHandler()
	http3.ListenAndServe(":8080", crtPath, keyPath, handler)
}
