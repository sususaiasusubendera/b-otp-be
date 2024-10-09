package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// STRUCT
type Page struct {
	Title string
	Body  []byte
}

// FUNCTION & METHOD
// save method
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

// loadPage function
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
	return &Page{Title: title, Body: body}, nil
}

// viewHandler function
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

// MAIN
func main() {
    http.HandleFunc("/view/", viewHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
