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

func setupHandler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/assets/proposedMethod/100B_160+10KB_30+100KB_10(200)/", http.StripPrefix("/assets/proposedMethod/100B_160+10KB_30+100KB_10(200)/", http.FileServer(http.Dir("assets/proposedMethod/100B_160+10KB_30+100KB_10(200)/"))))
	mux.HandleFunc("/", IndexHandler)

	return mux
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	requestHeader := r.Header.Get("Accept")
	accept := strings.Split(requestHeader, ",")
	responsHeader := w.Header()
	var text = "text/html"
	if accept[0] == text {
		responsHeader.Set("Alt-Svc", "h2=':8080'; ma=2592000;")
	}

	data, _ := ioutil.ReadDir("assets/proposedMethod/100B_160+10KB_30+100KB_10(200)")
	var templates = template.Must(template.ParseFiles("templates/index.html"))
	if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
		log.Fatalln("Unable to execute template.")
	}
}

func main() {
	// 証明書
	var crtPath = "letsencrypt/live/localhost111919.ml/fullchain.pem"
	var keyPath = "letsencrypt/live/localhost111919.ml/privkey.pem"

	// HTTP/3サーバー起動
	handler := setupHandler()
	err := http3.ListenAndServe(":8080", crtPath, keyPath, handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	// HTTP/2サーバー起動
	// http.Handle("/assets/proposedMethod/100B_160+10KB_30+100KB_10(200)/", http.StripPrefix("/assets/proposedMethod/100B_160+10KB_30+100KB_10(200)/", http.FileServer(http.Dir("assets/proposedMethod/100B_160+10KB_30+100KB_10(200)"))))
	// http.HandleFunc("/", IndexHandler)
	// err := http.ListenAndServeTLS(":8080", crtPath, keyPath, nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
}
