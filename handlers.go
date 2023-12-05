package main

import (
	"html/template"
	"log"
	"net/http"
)

// os templates sao usados para servir conteudos dinamicos
// para executar os templates criamos esta funcao
func (app *Application) ContactHandler(w http.ResponseWriter, r *http.Request) {
	app.RenderTemplate(w, r, "contact")

}
func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	app.RenderTemplate(w, r, "index")

}
func (app *Application) AboutHandler(w http.ResponseWriter, r *http.Request) {
	app.RenderTemplate(w, r, "about")

}

func (app *Application) PageLoginHandler(w http.ResponseWriter, r *http.Request) {
	app.RenderTemplate(w, r, "login")
}

// responsável pelo logout
func (app *Application) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// definimos a variavel responsavel por verificar a auth como false
	IsAuth = false

	// removemos o auth dos cookies
	cookie := http.Cookie{
		Name:   "IsAuthenticated",
		Value:  "false",
		MaxAge: -1, // Define o tempo de vida do cookie para que ele expire imediatamente.
		// Configure outras opções do cookie conforme necessário.
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
func (app *Application) AdminHandler(w http.ResponseWriter, r *http.Request) {

	// se o método de request for diferente de POST, o usuario é direcionado pro login
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// pegamos os valores dos forms
	email := r.FormValue("email")
	senha := r.FormValue("senha")

	// Lógica de autenticação aqui...
	if app.VerificaUsuario(email, senha) {
		// Redirecionar para a página de contato em caso de autenticação bem-sucedida
		// se o usuario existir, altero a variavel auth para true... lembrando que essa variável é responsável
		// por todo conteúdo exclusivo para os clientes logados

		IsAuth = true

		// adiciono nos cookies ao fazer o login
		cookie := http.Cookie{
			Name:  "IsAuthenticated",
			Value: "true",
			// Configure outras opções do cookie conforme necessário (tempo de expiração, domínio, caminho, etc.).
		}
		http.SetCookie(w, &cookie)

		var t *template.Template
		var err error

		t, err = template.ParseFiles(
			"templates/admin.gohtml",
			"templates/navbar.gohtml",
			"templates/base.gohtml",
		)

		// Se houver um erro no parse, imprime o erro no log e retorna.
		if err != nil {
			log.Println(err)
			return
		}

		// Executa o template, escrevendo o resultado em http.ResponseWriter.
		// podemos passar um objeto no como parametro
		user, _ := app.GetUserByName(email)

		// eu passo essa data por causa do nome do usuario
		data := struct {
			IsAuthenticated bool
			CurrentUser     User
		}{
			IsAuthenticated: IsAuth,
			CurrentUser:     user,
		}

		t.Execute(w, data)
		return
	} else {
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
	}

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
	app.RenderTemplate(w, r, "index")
}
