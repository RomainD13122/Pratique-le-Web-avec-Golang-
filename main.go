package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

type Etudiant struct {
	Prenom string
	Nom    string
	Age    int
	Sexe   string
}

type Classe struct {
	Nomdlc            string
	Filiere           string
	Niveau            string
	Compteurdetudiant int
	Etudiants         []Etudiant
}

func promoHandler(w http.ResponseWriter, r *http.Request) {

	classe := Classe{
		Nomdlc:            "B1 Informatique",
		Filiere:           "Informatique",
		Niveau:            "Bachelor 1",
		Compteurdetudiant: 3,
		Etudiants: []Etudiant{
			{Prenom: "Eddy", Nom: "Amir", Age: 17, Sexe: "Homme"},
			{Prenom: "Xerly", Nom: "Ji", Age: 38, Sexe: "Femme"},
			{Prenom: "Dimitri", Nom: "Mendeleiv", Age: 58, Sexe: "Homme"},
		},
	}

	tmpl, err := template.ParseFiles("templates/promo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, classe)
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/promo", promoHandler)

	http.ListenAndServe(":8080", nil)
}

var (
	viewCount int
	mu        sync.Mutex
)

func changeHandler(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	viewCount++
	currentViewCount := viewCount
	mu.Unlock()

	var message, couleur string
	if currentViewCount%2 == 0 {
		message = fmt.Sprintf("Le nombre de vues est pair : %d", currentViewCount)
		couleur = "#d1e7dd"
	} else {
		message = fmt.Sprintf("Le nombre de vues est impair : %d", currentViewCount)
		couleur = "#f8d7da"
	}

	data := struct {
		Message   string
		Couleurll string
	}{
		Message:   message,
		Couleurll: couleur,
	}

	tmpl, err := template.ParseFiles("templates/change.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}
