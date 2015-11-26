package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() { 

	router := httprouter.New()
	router.GET("/keys", GetKeysClient)
	router.GET("/keys/:key_id", GetKeyValue)
	router.PUT("/keys/:key_id/:value", PutUser)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}

	server.ListenAndServe()

}
