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

func main() {

	cache := make(map[string]*template.Template)

	config := Config{
		Port:    "3000",
		Env:     "dev",
		Version: "1.0.0",
	}

	app := Application{
		Config: config,
		Cache:  cache,
	}

	app.Routes()
	log.Printf("Servidor ' %s ' escutando na porta %s", config.Env, config.Port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", config.Port), nil)

	if err != nil {
		log.Println(err)
	}
}
