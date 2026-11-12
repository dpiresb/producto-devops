package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type Pagina struct {
	Titulo  string
	Mensaje string
	Imagen  string
}

var plantilla = template.Must(template.ParseFiles("templates/index.html"))

func inicio(w http.ResponseWriter, r *http.Request) {
	nombrePod, _ := os.Hostname()
	w.Header().Set("X-Pod-Name", nombrePod)
	datos := Pagina{
		Titulo:  "Producto 1 - Entorno de desarrollo",
		Mensaje: "Soy alumno de la UOC",
		Imagen:  "/static/img/uoc.jpg",
	}
	if err := plantilla.Execute(w, datos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	puerto := os.Getenv("PORT")
	if puerto == "" {
		puerto = "8080"
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", inicio)

	log.Printf("Servidor escuchando en http://localhost:%s", puerto)
	log.Fatal(http.ListenAndServe(":"+puerto, nil))
}