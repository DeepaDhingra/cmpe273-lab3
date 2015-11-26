package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type GetPair struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}

func GetKeyServer3000(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	fmt.Println("M is", M1)
	var counter int
	counter = 0

	psgo := make([]GetPair, len(M1))

	for key, value := range M1 {

		fmt.Println("Key:", key, "Value:", value)
		psgo[counter].Key = key
		psgo[counter].Value = value

		counter++
		fmt.Println("counter is", counter)

	}

	uj, _ := json.Marshal(psgo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Println("uj is", uj)
	fmt.Fprintf(w, "%s", uj)
}

func GetKeyServer3001(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	fmt.Println("M is", M2)
	var counter int
	counter = 0

	psgo := make([]GetPair, len(M2))

	for key, value := range M2 {

		fmt.Println("Key:", key, "Value:", value)
		psgo[counter].Key = key
		psgo[counter].Value = value

		counter++
		fmt.Println("counter is", counter)

	}

	uj, _ := json.Marshal(psgo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Println("uj is", uj)
	fmt.Fprintf(w, "%s", uj)
}

func GetKeyServer3002(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	fmt.Println("M is", M3)
	var counter int
	counter = 0

	psgo := make([]GetPair, len(M3))

	for key, value := range M3 {

		fmt.Println("Key:", key, "Value:", value)
		psgo[counter].Key = key
		psgo[counter].Value = value

		counter++
		fmt.Println("counter is", counter)

	}

	uj, _ := json.Marshal(psgo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Println("uj is", uj)
	fmt.Fprintf(w, "%s", uj)
}
