package main

import (
	"net/http"
)

// os templates sao usados para servir conteudos dinamicos
// para executar os templates criamos esta funcao
func (app *Application) ContactHandler(w http.ResponseWriter, _ *http.Request) {
	app.RenderTemplate(w, "contact")

}
func (app *Application) HomeHandler(w http.ResponseWriter, _ *http.Request) {
	app.RenderTemplate(w, "index")

}
func (app *Application) AboutHandler(w http.ResponseWriter, _ *http.Request) {
	app.RenderTemplate(w, "about")

}

func (app *Application) PageLoginHandler(w http.ResponseWriter, _ *http.Request) {
	app.RenderTemplate(w, "login")

}

func (app *Application) AdminHandler(w http.ResponseWriter, _ *http.Request) {
	app.RenderTemplate(w, "admin")

}

func (app *Application) AccountHandler(w http.ResponseWriter, _ *http.Request) {
	app.RenderTemplate(w, "createaccount")

}

func (app *Application) CreateAccHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	nome := r.FormValue("nome")
	contacto := r.FormValue("contacto")
	senha := r.FormValue("senha")
	//senha2 := r.FormValue("senha2")

	app.InsertUsers(email, nome, contacto, senha)
	app.RenderTemplate(w, "index")

}
func (app *Application) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	senha := r.FormValue("senha")

	// Lógica de autenticação aqui...
	if app.VerificaUsuario(email, senha) {
		// Redirecionar para a página de contato em caso de autenticação bem-sucedida
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	// Em caso de falha na autenticação, você pode redirecionar para uma página de erro ou exibir uma mensagem.
	http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
}
