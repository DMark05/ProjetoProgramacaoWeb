package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var teste struct {
	valor string
} = struct{ valor string }{valor: "um valor"}

// {
// 	"e-mail": "example@example.com",
//	"password": "pass123"
// }

type LoginForm struct {
	Email    string `json:"e-mail"`
	Password string `json:"password"`
}

func getLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		{
			output, err := io.ReadAll(r.Body)
			if err != nil {
				log.Fatalln(err)
			}
			var login_form LoginForm = LoginForm{}
			err = json.Unmarshal(output, &login_form)
			if err != nil {
				log.Fatalln(err)
			}
			domain := strings.Split(login_form.Email, "@")[1]
			if strings.Compare(domain, "iscte-iul.pt") != 0 {

			}
			switch domain {
			case "iscte-iul.pt":
				{

				}
			case "gmail.com":
				{

				}
			}
		}
	}
}

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
