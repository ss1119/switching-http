package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	_ "net/http/pprof"

	"github.com/lucas-clemente/quic-go/http3"
)

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
	// 証明書
	var crtPath = "./localhost19.ml/localhost19.ml.pem"
	var keyPath = "./localhost19.ml/localhost19.ml-key.pem"

	// HTTP/3サーバー起動
	handler := setupHandler()
	err := http3.ListenAndServe(":8080", crtPath, keyPath, handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	// HTTP/2サーバー起動
	// http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	// http.HandleFunc("/", IndexHandler)
	// err := http.ListenAndServeTLS(":8080", crtPath, keyPath, nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
}
