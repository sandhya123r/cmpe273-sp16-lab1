package main

import (
	//"github.com/drone/routes"
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
	"fmt"
	"log"
)

type FoodType struct {
	Type string `json:"type"`
	Drink_alcohol string `json:"drink_alcohol"`
}
type Profile struct {
	Email string `json:"email"`
	Zip string `json:"zip"`
	Country string `json:"country"`
	
	Food FoodType `json:"food"`
	Profession string `json:"profession"`
	Favorite_color string `json:"favorite_color"`
	Is_smoking string `json:"is_smoking"`
	Favorite_sport string `json:"favorite_sport"`
	Music struct {
		Spotify_user_id string `json:"spotify_user_id"`
	} `json:"music"`
	Movie struct {
		Tv_shows []string `json:"tv_shows"`
		Movies []string `json:"movies"`
	} `json:"movie"`
	Travel struct {
		Flight struct {
			Seat string `json:"seat"`
		} `json:"flight"`
	} `json:"travel"`
	
}

var all_profiles = make(map[string]*Profile)

func main() {
	
	router :=mux.NewRouter()
	
	router.HandleFunc("/profile/{email}",GetProfile).Methods("GET")
	router.HandleFunc("/profile",CreateProfile).Methods("POST")
	router.HandleFunc("/profile/{email}",UpdateProfile).Methods("PUT")
	router.HandleFunc("/profile/{email}",DeleteProfile).Methods("DELETE")
	log.Println("Listening..")
	http.ListenAndServe(":3000", router)
}


func GetProfile(w http.ResponseWriter, r *http.Request) {

	vars:=mux.Vars(r)
	email :=vars["email"]
	
	profile:=  all_profiles[email]
	x := make([]*Profile, 1)
	x[0] = profile
	js, error := json.Marshal(x)
	
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(js))
}	


func CreateProfile(w http.ResponseWriter, r *http.Request) {	
	vars := mux.Vars(r)
	email :=vars["email"]
	profile := new(Profile)
	profile.Email = email
	
	decoder := json.NewDecoder(r.Body)
	error := decoder.Decode(&profile)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
	all_profiles[profile.Email] = profile
	js, err := json.Marshal(profile)
	if err != nil {
		
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(js))
	
}

func DeleteProfile(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	delete(all_profiles,vars["email"])
	w.WriteHeader(http.StatusNoContent)
	fmt.Fprint(w,"Success")
}

func UpdateProfile(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	email :=vars["email"]
	profile,ok := all_profiles[email]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	decoder := json.NewDecoder(r.Body)
	error := decoder.Decode(&profile)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	profile.Email=email
	all_profiles[email]=profile
	js, err := json.Marshal(profile)

	if err != nil {		
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	fmt.Fprint(w, string(js))
	
	  
}
