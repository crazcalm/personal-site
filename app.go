package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	portCMDDoc = "port to be used (Example: ':8080')"
)

var (
	port   string
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Llongfile)
)

func init() {
	flag.StringVar(&port, "port", ":8080", portCMDDoc)
	flag.StringVar(&port, "p", ":8080", portCMDDoc)
}

type page struct {
	name         string
	templatePage string
	log          *log.Logger
}

func (p page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t, err := template.New(p.name).Parse(string(p.templatePage))
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}

	p.log.Printf("Serving %s", p.name)
}

func main() {
	flag.Parse()

	indexPagePath, err := ioutil.ReadFile(filepath.Join(".", "templates", "personal_website.html"))
	if err != nil {
		log.Fatal(err)
	}
	indexPage := page{name: "indexPage", templatePage: string(indexPagePath), log: logger}

	pokemonPagePath, err := ioutil.ReadFile(filepath.Join(".", "templates", "pokemon_site.html"))
	if err != nil {
		log.Fatal(err)
	}
	pokemonPage := page{name: "pokemonPage", templatePage: string(pokemonPagePath), log: logger}

	mux := http.NewServeMux()
	mux.Handle("/", indexPage)
	mux.Handle("/pokemon", pokemonPage)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(".", "static")))))

	server := http.Server{Addr: port, Handler: mux}
	log.Fatal(server.ListenAndServe())
}
