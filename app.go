package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	indexPage, err := ioutil.ReadFile(filepath.Join(".", "templates", "personal_website.html"))
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New("indexPage").Parse(string(indexPage))
		if err != nil {
			log.Fatal(err)
		}
		err = t.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	})
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(".", "static")))))

	server := http.Server{Addr: ":8081", Handler: mux}
	log.Fatal(server.ListenAndServe())
}
