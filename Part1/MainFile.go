package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var M1 = make(map[int]string)
var M2 = make(map[int]string)
var M3 = make(map[int]string)

func main() { 

	router1 := httprouter.New()
	router1.GET("/keys", GetKeyServer3000)
	router1.PUT("/keys/:key_id/:value", PutServer1)
	server := http.Server{
		Addr:    "0.0.0.0:3000",
		Handler: router1,
	}
	router2 := httprouter.New()
	router2.GET("/keys", GetKeyServer3001)
	router2.PUT("/keys/:key_id/:value", PutServer2)

	go func() {
		http.ListenAndServe("localhost:3001", router2)
	}()

	router3 := httprouter.New()
	router3.GET("/keys", GetKeyServer3002)
	router3.PUT("/keys/:key_id/:value", PutServer3)

	go func() {
		http.ListenAndServe("localhost:3002", router3)
	}()

	server.ListenAndServe()

}
