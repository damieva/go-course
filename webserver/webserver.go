package webserver

import (
	"net/http"
)

func MiWebServer() {
	http.HandleFunc("/", home)        // definimos un manejador
	http.ListenAndServe(":3000", nil) // quedate escuchando en el puerto 3000
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./webserver/index.html")
}
