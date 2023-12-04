package main

import (
	"log"
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

func (app *Application) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	senha := r.FormValue("senha")

	// Log das credenciais (não faça isso em produção, apenas para depuração)
	log.Printf("Tentativa de login - Email: %s, Senha: %s\n", email, senha)

	// Lógica de autenticação aqui...
	if app.ValidarCredenciais(email, senha) {
		// Redirecionar para a página de contato em caso de autenticação bem-sucedida
		http.Redirect(w, r, "/contact", http.StatusSeeOther)
		return
	}

	// Em caso de falha na autenticação, você pode redirecionar para uma página de erro ou exibir uma mensagem.
	http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
}

// Função de exemplo para validar credenciais
func (app *Application) ValidarCredenciais(email, senha string) bool {
	// Lógica de validação de credenciais (substitua isso com sua lógica real)
	return email == "user@gmail.com" && senha == "123"
}
