package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// func StaticHandler(w http.ResponseWriter, r *http.Request) {

// responsável por ler um arquivo estático, retorna o arquivo em si
// 	f, err := os.Open("static" + r.URL.Path)

// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// fazemos o tratamento do arquivo css para o navegador entender o css

// 	if strings.HasSuffix(r.URL.Path, ".css") {
// 		w.Header().Add("Content-Type", "text/css")
// 	}

//  responsável por copiar o conteudo entre um reader e um writer
//  primeiro passamos o destino, e depois de onde vem
// 	io.Copy(w, f)
// }

// os templates sao usados para servir conteudos dinamicos
// para executar os templates criamos esta funcao
func ContactHandler(w http.ResponseWriter, _ *http.Request) {
	RenderTemplate(w, "contact")

}
func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	RenderTemplate(w, "index")

}
func AboutHandler(w http.ResponseWriter, _ *http.Request) {
	RenderTemplate(w, "about")

}

// variavel responsavel por ignorar o cache durante o desenvolvimento
var env = "dev"

// aqui começamos a tratar o cache
var cache map[string]*template.Template

// RenderTemplate renderiza uma página usando um template e escreve o resultado em http.ResponseWriter.
func RenderTemplate(w http.ResponseWriter, page string) {
	// Declaração de variáveis para o template e erro.
	var t *template.Template
	var err error

	// Verifica se a página está em cache.
	_, exists := cache[page]

	// Se a página não está em cache ou estamos no ambiente de desenvolvimento ("dev").
	if !exists || env == "dev" {
		// Faz o parse dos arquivos de template.
		t, err = template.ParseFiles(
			"templates/"+page+".html",
			"templates/base.html",
		)
		// Se houver um erro no parse, imprime o erro no log e retorna.
		if err != nil {
			log.Println(err)
			return
		}
		// Adiciona o template ao cache.
		cache[page] = t
	} else {
		// Se a página está em cache, imprime "Cache hit".
		fmt.Println("Cache hit")
		// Obtém o template do cache.
		t = cache[page]
	}

	// Executa o template, escrevendo o resultado em http.ResponseWriter.
	t.Execute(w, nil)
}


func main() {

	cache = make(map[string]*template.Template)

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/contact", ContactHandler)
	http.HandleFunc("/about", AboutHandler)

	// vamos utilizar a função http.FileServer para servir conteúdo estático de forma mais fácil
	// neste caso, é usado para servir os ficheiros js e css
	http.Handle("/static/",
	http.StripPrefix("/static/",
	http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":3000", nil)

}