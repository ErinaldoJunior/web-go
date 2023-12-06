package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type TemplateData struct {
	IsAuthenticated bool
	CurrentUser     User
}

//usado para embutir arquivos que não são do tipo .go no build

// comentario...go:embed templates
//var TemplateFS embed.FS

// RenderTemplate renderiza uma página usando um template e escreve o resultado em http.ResponseWriter.
func (a *Application) RenderTemplate(w http.ResponseWriter, r *http.Request, page string, data any) {

	// Declaração de variáveis para o template e erro.
	var t *template.Template
	var err error

	// Verifica se a página está em cache.
	_, exists := a.Cache[page]

	// Se a página não está em cache ou estamos no ambiente de desenvolvimento ("dev").
	if !exists || a.Config.Env == "dev" {
		// Faz o parse dos arquivos de template.
		// t, err = template.ParseFS(
		// 	TemplateFS,
		// 	"templates/"+page+".html",
		// 	"templates/navbar.html",
		// 	"templates/login.html",
		// 	"templates/base.html",
		// )

		t, err = template.ParseFiles(
			"templates/base.gohtml",
			"templates/"+page+".gohtml",
			"templates/navbar.gohtml",
		)

		// Se houver um erro no parse, imprime o erro no log e retorna.
		if err != nil {
			log.Println(err)
			return
		}
		// Adiciona o template ao cache.
		a.Cache[page] = t
	} else {
		// Se a página está em cache, imprime "Cache hit".
		fmt.Println("Cache hit")
		// Obtém o template do cache.
		t = a.Cache[page]
	}

	// Executa o template, escrevendo o resultado em http.ResponseWriter.
	// podemos passar um objeto no como parametro

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	t.ExecuteTemplate(w, "base", data)
}
