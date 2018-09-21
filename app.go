package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

const (
	portCMDDoc = "port to be used (Example: ':8080')"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "port", ":8080", portCMDDoc)
	flag.StringVar(&port, "p", ":8080", portCMDDoc)
}

func main() {
	flag.Parse()

	indexPage, err := ioutil.ReadFile(filepath.Join(".", "templates", "personal_website.html"))
	if err != nil {
		log.Fatal(err)
	}

	pokemonPage, err := ioutil.ReadFile(filepath.Join(".", "templates", "pokemon_site.html"))
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

	mux.HandleFunc("/pokemon", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New("pokemonPage").Parse(string(pokemonPage))
		if err != nil {
			log.Fatal(err)
		}
		err = t.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	})

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(".", "static")))))

	server := http.Server{Addr: port, Handler: mux}
	log.Fatal(server.ListenAndServe())
}
