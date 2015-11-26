package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type PutStruct struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}

var Keyidint int
var Val string

func PutServer1(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	ps := PutStruct{}

	keyid := p.ByName("key_id")
	Val = p.ByName("value")

	Keyidint, _ = strconv.Atoi(keyid)
	M1[Keyidint] = Val
	fmt.Println("key is", Keyidint)
	fmt.Println("value is", Val)

	ps.Key = Keyidint
	ps.Value = Val
	fmt.Println("ps is", ps)
	uj, _ := json.Marshal(ps)
	fmt.Println("uj is", uj)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)

}

func PutServer2(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	ps := PutStruct{}

	keyid := p.ByName("key_id")
	Val = p.ByName("value")

	Keyidint, _ = strconv.Atoi(keyid)
	M2[Keyidint] = Val
	ps.Key = Keyidint
	ps.Value = Val
	uj, _ := json.Marshal(ps)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)

}

func PutServer3(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	ps := &PutStruct{}

	keyid := p.ByName("key_id")
	Val = p.ByName("value")

	Keyidint, _ = strconv.Atoi(keyid)
	M3[Keyidint] = Val
	ps.Key = Keyidint
	ps.Value = Val
	uj, _ := json.Marshal(ps)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)

}
