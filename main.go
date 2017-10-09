package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)

func main() {
	router := mux.NewRouter()
	//mock data
	people = append(people, Person{ID:"1",Firstname:"David",Lastname:"Barboza", Address: &Address{City:"Zapopan",State:"Jalisco"}})
	people = append(people, Person{ID:"2",Firstname:"Lili",Lastname:"Alcalde"})

	router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", GetPeopleEndPoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":1234", router))
}

func GetPersonEndPoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	for _, item:= range people{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request){
	json.NewEncoder(w).Encode(people)
}

func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)

	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndPoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	for index, item:= range people{
		if item.ID == params["id"]{
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

var people []Person

type Person struct {
	ID string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Address *Address `json:"address,omitempty"`
}

type Address struct {
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}