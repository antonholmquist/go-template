package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"code.google.com/p/go.crypto/bcrypt"
)

func HeaderMiddleware(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {

	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Origin, Accept, Content-Type")
	res.Header().Set("Access-Control-Allow-Method", "GET, POST, PUT, DELETE")
	
	// do some stuff before
	next(res, req)
	// do some stuff after
}

func main() {

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "3001"
	}

	// gorilla/mux is a powerful URL router and dispatcher.
	router := mux.NewRouter()

	router.Methods("GET").Path("/").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		res.Header().Set("Content-Type", "text/plain")
		res.Write([]byte("test application"))
	})

	fmt.Println("Starting server on port", port)
	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(HeaderMiddleware))
	n.UseHandler(router)
	
	n.Run(":" + port)

}
