package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	adorable "github.com/ipsn/go-adorable"
)

func getRandomAvatar(w http.ResponseWriter, r *http.Request) {
	log.Printf("request for %s", r.URL.Path)
	avatar := adorable.Random()
	w.Write(avatar)
}

func getRandomAvatarBasedOnUniqueUserID(w http.ResponseWriter, r *http.Request) {
	log.Printf("request for %s", r.URL.Path)
	qParam := r.URL.Query()
	userID := qParam.Get("userID")
	if len(userID) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "userID is not given as a query parameter.")
	} else {
		avatar := adorable.PseudoRandom([]byte(userID))
		w.Write(avatar)
	}
}

func greet(w http.ResponseWriter, r *http.Request) {
	log.Printf("request for /")
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/getRandomAvatar", getRandomAvatar)
	http.HandleFunc("/getRandomAvatarBasedOnUniqueUserID", getRandomAvatarBasedOnUniqueUserID)
	log.SetPrefix("localhost:8080: ")
	log.Printf("The server will listen on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
