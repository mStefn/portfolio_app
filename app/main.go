package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

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
	return err
}

// ---------- Strony główne ----------

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

	data := PageData{Title: "Moje Portfolio", Visits: visits}
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, data)
}

func serveAbout(w http.ResponseWriter, r *http.Request) {
	data := PageData{Title: "About Me"}
	tmpl := template.Must(template.ParseFiles("templates/about.html"))
	tmpl.Execute(w, data)
}

func serveProjects(w http.ResponseWriter, r *http.Request) {
	data := PageData{Title: "Projects"}
	tmpl := template.Must(template.ParseFiles("templates/projects.html"))
	tmpl.Execute(w, data)
}

func serveContact(w http.ResponseWriter, r *http.Request) {
	data := PageData{Title: "Contact"}
	tmpl := template.Must(template.ParseFiles("templates/contact.html"))
	tmpl.Execute(w, data)
}

// ---------- Dynamiczna obsługa projektów ----------

func serveDynamicProject(w http.ResponseWriter, r *http.Request) {
	project := strings.TrimPrefix(r.URL.Path, "/projects/")
	templatePath := "templates/projects/" + project + ".html"

	data := PageData{Title: "Project - " + project}
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	tmpl.Execute(w, data)
}

// ---------- Main ----------

func main() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/about", serveAbout)
	http.HandleFunc("/projects", serveProjects)
	http.HandleFunc("/contact", serveContact)
	http.HandleFunc("/projects/", serveDynamicProject)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Serwer działa na http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Błąd serwera:", err)
	}
}
