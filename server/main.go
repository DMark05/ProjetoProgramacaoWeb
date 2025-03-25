package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var teste struct {
	valor string
} = struct{ valor string }{valor: "um valor"}


func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	if strings.Compare(r.Header.Get("Content-Type"), "application/json") != 0 {
		http.Redirect(w, r, "http://localhost:3000/", http.StatusSeeOther)
	}
	switch r.Method {
	case "GET":
		io.WriteString(w, teste.valor)
	case "PUT":
		output, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}
		teste.valor = string(output)
	}
}

func main() {

	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
