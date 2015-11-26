package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type GetPair struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}

func GetKeysClient(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var counter int
	counter = 0
	psgo := make([]GetPair, len(MClient))
	for key, value := range MClient {
		fmt.Println("Key:", key, "Value:", value)
		psgo[counter].Key = key
		psgo[counter].Value = value
		counter++
		fmt.Println("counter is", counter)
	}

	uj, _ := json.Marshal(psgo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

func GetKeyValue(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	keyid := ps.ByName("key_id")
	Keyidint, _ := strconv.Atoi(keyid)

	value, ok := MClient[Keyidint]
	if ok {
		fmt.Println("value: ", value)
	} else {
		fmt.Println("key not found")
	}

	pskey := GetPair{}
	pskey.Key = Keyidint
	pskey.Value = value
	uj, _ := json.Marshal(pskey)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}
