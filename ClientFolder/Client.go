package main

import (
	"bytes"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"stathat.com/c/consistent"
	"strconv"
)

type ResponseStruct struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}

var url string
var MClient = make(map[int]string)

func PutUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	keyid := p.ByName("key_id")
	Val := p.ByName("value")

	Keyidint, _ := strconv.Atoi(keyid)
	MClient[Keyidint] = Val

	cons := consistent.New()

	cons.Add("127.0.0.1:3000")
	cons.Add("127.0.0.1:3001")
	cons.Add("127.0.0.1:3002")

	fmt.Println("Servers added")

	//Find a cache server for a user:

	server, err := cons.Get(Val)
	if err != nil {
	}
	fmt.Println("server:", server)

	s := []string{}
	s = append(s, "http://")
	s = append(s, server)
	s = append(s, "/")
	s = append(s, "keys")
	s = append(s, "/")
	s = append(s, keyid)
	s = append(s, "/")
	s = append(s, Val)

	url = s[0] + s[1] + s[2] + s[3] + s[4] + s[5] + s[6] + s[7]

	var urlString = []byte(url)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(urlString))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", string(body))

}
