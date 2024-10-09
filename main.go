package main

import (
	"fmt"
	"os"
)


// Struct
type Page struct {
	Title string
	Body  []byte
}

// Function & Method
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
	return &Page{Title: title, Body: body}, nil
}

// Main
func main() {
    p1:= &Page{Title: "test", Body: []byte("A test page.")}
    p1.save()
    p2, _ := loadPage("test")
    fmt.Println(string(p2.Body))
}
