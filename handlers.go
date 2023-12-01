package main

import "net/http"

// os templates sao usados para servir conteudos dinamicos
// para executar os templates criamos esta funcao
func (app *Application) ContactHandler(w http.ResponseWriter, _ *http.Request) {
	app.RenderTemplate(w, "contact")

}
func (app *Application) HomeHandler(w http.ResponseWriter, _ *http.Request) {
	app.RenderTemplate(w, "index")

}
func (app *Application)AboutHandler(w http.ResponseWriter, _ *http.Request) {
	app.RenderTemplate(w, "about")

}