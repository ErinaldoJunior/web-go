package main

import "net/http"

// função responsável pelas rotas
func (app *Application) Routes() {

	http.HandleFunc("/", app.HomeHandler)
	http.HandleFunc("/contact", app.ContactHandler)
	http.HandleFunc("/about", app.AboutHandler)

	// vamos utilizar a função http.FileServer para servir conteúdo estático de forma mais fácil
	// neste caso, é usado para servir os ficheiros js e css
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
}