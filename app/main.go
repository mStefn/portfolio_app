package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// Struktura dla danych do template (na razie pusta, do rozbudowy)
type PageData struct {
	Title string
}

func main() {
	// Ścieżka do template
	tmplPath := filepath.Join("templates", "index.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))

	// Obsługa strony głównej
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title: "Moje Portfolio",
		}
		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Obsługa statycznych plików (CSS itd.)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Serwer działa na http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Błąd serwera:", err)
	}
}
