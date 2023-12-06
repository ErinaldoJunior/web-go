package main

import (
	"net/http"
)

var IsAuth bool = false

// os templates sao usados para servir conteudos dinamicos
// para executar os templates criamos esta funcao
func (app *Application) ContactHandler(w http.ResponseWriter, r *http.Request) {
	app.RenderTemplate(w, r, "contact", TemplateData{
		IsAuthenticated: IsAuth,
	})

}
func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	app.RenderTemplate(w, r, "index", TemplateData{
		IsAuthenticated: IsAuth,
	})

}
func (app *Application) AboutHandler(w http.ResponseWriter, r *http.Request) {
	app.RenderTemplate(w, r, "about", TemplateData{
		IsAuthenticated: IsAuth,
	})

}

func (app *Application) PageLoginHandler(w http.ResponseWriter, r *http.Request) {
	app.RenderTemplate(w, r, "login", TemplateData{
		IsAuthenticated: IsAuth,
	})
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

		// Executa o template, escrevendo o resultado em http.ResponseWriter.
		// podemos passar um objeto no como parametro
		user, _ := app.GetUserByName(email)

		app.RenderTemplate(w, r, "admin", TemplateData{
			IsAuthenticated: IsAuth,
			CurrentUser:     user,
		})
		return
	} else {
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
	}
}

func (app *Application) CreateAccHandler(w http.ResponseWriter, r *http.Request) {

	app.RenderTemplate(w, r, "createaccount", TemplateData{
		IsAuthenticated: IsAuth,
	})

}

func (app *Application) CreaterUserHandler(w http.ResponseWriter, r *http.Request) {
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
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
