package main

import (
	"html/template"
	"log"
	"net/http"
)

// func StaticHandler(w http.ResponseWriter, r *http.Request) {

// 	//responsável por ler um arquivo estático, retorna o arquivo em si
// 	f, err := os.Open("static" + r.URL.Path)

// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	// fazemos o tratamento do arquivo css para o navegador entender o css

// 	if strings.HasSuffix(r.URL.Path, ".css") {
// 		w.Header().Add("Content-Type", "text/css")
// 	}

// 	// responsável por copiar o conteudo entre um reader e um writer
// 	// primeiro passamos o destino, e depois de onde vem
// 	io.Copy(w, f)
// }

// os templates sao usados para servir conteudos dinamicos
// para executar os templates criamos esta funcao
func ContactHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "contact")

}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index")

}
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about")

}

func RenderTemplate(w http.ResponseWriter, page string) {
	// usamos a biblioteca template, a função parse files retorna
	// um template e um erro
	tp, err := template.ParseFiles("templates/"+ page + ".html")

	//tratamos o erro
	if err != nil {
		log.Println(err)
		return
	}

	// usamos a função execute para escrever o template para o usuario
	tp.Execute(w, nil)
}

func main() {

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