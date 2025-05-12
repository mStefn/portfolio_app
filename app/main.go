package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

// Struktura danych dla strony
type PageData struct {
	Title  string
	Visits int
}

// Funkcja do pobierania liczby wizyt
func getVisitCount() (int, error) {
	dbURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT visits FROM visit_counter WHERE id = 1").Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Funkcja do inkrementacji liczby wizyt
func incrementVisitCount() error {
	dbURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE visit_counter SET visits = visits + 1 WHERE id = 1")
	if err != nil {
		return err
	}

	return nil
}

// Funkcja do ustawiania nagłówka CSP
func setCSP(w http.ResponseWriter) {
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' https://formspree.io; style-src 'self';")
}

// Funkcja obsługująca stronę główną
func serveHome(w http.ResponseWriter, r *http.Request) {
	err := incrementVisitCount()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	visits, err := getVisitCount()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:  "Moje Portfolio",
		Visits: visits,
	}

	tmplPath := "templates/index.html"
	tmpl := template.Must(template.ParseFiles(tmplPath))

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Funkcja obsługująca stronę "About"
func serveAbout(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:  "About Me",
		Visits: 0, 
	}

	tmplPath := "templates/about.html"
	tmpl := template.Must(template.ParseFiles(tmplPath))

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Funkcja obsługująca projekt 1
func serveProject1(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:  "Project 1",
		Visits: 0,
	}

	tmplPath := "templates/project1.html"
	tmpl := template.Must(template.ParseFiles(tmplPath))

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Funkcja obsługująca projekt 2
func serveProject2(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:  "Project 2",
		Visits: 0,
	}

	tmplPath := "templates/project2.html"
	tmpl := template.Must(template.ParseFiles(tmplPath))

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Funkcja obsługująca stronę kontaktową
func serveContact(w http.ResponseWriter, r *http.Request) {
	setCSP(w) // Dodanie nagłówka CSP

	data := PageData{
		Title:  "Contact",
		Visits: 0,
	}

	tmplPath := "templates/contact.html"
	tmpl := template.Must(template.ParseFiles(tmplPath))

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Główna funkcja
func main() {
	// Obsługa tras
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/about", serveAbout)
	http.HandleFunc("/project1", serveProject1)
	http.HandleFunc("/project2", serveProject2)
	http.HandleFunc("/contact", serveContact)

	// Serwowanie plików statycznych (np. CSS, JS)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Uruchomienie serwera
	log.Println("Serwer działa na http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Błąd serwera:", err)
	}
}
