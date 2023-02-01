package main

import "net/http"

func main() {
	//routes
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/contact", ContactHandler)

	//start the server
	http.ListenAndServe(":3000", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hola mundo desde Go</h1>"))
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Pagina de contacto desde Go</h1>"))
}
