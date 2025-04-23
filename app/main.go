package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type PageData struct {
	Title string
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

func main() {
	tmplPath := "templates/index.html"
	tmpl := template.Must(template.ParseFiles(tmplPath))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
			Title: "Moje Portfolio",
			Visits: visits,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Serwer działa na http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Błąd serwera:", err)
	}
}

