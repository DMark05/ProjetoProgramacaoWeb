package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGOCONNSTRING = os.Getenv("MONGOCONNSTRING")
	MongoClient *mongo.Client
)

type MyCollection string

const (
	Users MyCollection = "users"
)

func getCollectionFromMongo(coll MyCollection) *mongo.Collection {
	switch coll {
	case Users:
		return MongoClient.Database("ProjetoInternet").Collection(string(Users))
	default:
		log.Fatalln("Collection not implemented")
		return nil
	}
}

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
	ctx := context.Background()
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/login", login)
	MongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGOCONNSTRING))
	if err != nil {
		log.Fatalln(err)
	}
	_ = MongoClient
	err = http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
