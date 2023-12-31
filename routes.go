package main

import "net/http"

// função responsável pelas rotas
func (app *Application) Routes() {

	http.HandleFunc("/", app.HomeHandler)
	http.HandleFunc("/contact", app.ContactHandler)
	http.HandleFunc("/about", app.AboutHandler)
	http.HandleFunc("/login", app.PageLoginHandler)
	http.HandleFunc("/logout", app.LogoutHandler)
	http.HandleFunc("/admin", app.AdminHandler)
	http.HandleFunc("/createaccount", app.CreateAccHandler)
	http.HandleFunc("/newAcc", app.CreaterUserHandler)

	// vamos utilizar a função http.FileServer para servir conteúdo estático de forma mais fácil
	// neste caso, é usado para servir os ficheiros js e css

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.Handle("/assets/",
		http.StripPrefix("/assets/",
			http.FileServer(http.Dir("assets"))))

}
