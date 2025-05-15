package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

// Struktura do przechowywania danych strony
type PageData struct {
	Title  string
	Visits int
}

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

// Obsługuje stronę główną
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

// Obsługuje stronę 'projects'
func serveProject1(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:  "Projects",
		Visits: 0,
	}

	tmplPath := "templates/projects.html"
	tmpl := template.Must(template.ParseFiles(tmplPath))

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Obsługuje stronę 'Contact'
func serveContact(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Contact",
		Visits: 0,
	}

	tmplPath := "templates/contact.html"
	tmpl := template.Must(template.ParseFiles(tmplPath))

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func main() {
	// Ustawienie routingu
	http.HandleFunc("/", serveHome)       // Strona główna
	http.HandleFunc("/about", serveAbout) // Strona "About"
	http.HandleFunc("/projects", serveProjects) // Strona "Projecs"
	http.HandleFunc("/contact", serveContact)

	// Obsługa plików statycznych
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Uruchomienie serwera
	log.Println("Serwer działa na http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Błąd serwera:", err)
	}
}
