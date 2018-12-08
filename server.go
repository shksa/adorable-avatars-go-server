package main

import (
	"fmt"
	"net/http"
	"time"

	adorable "github.com/ipsn/go-adorable"
)

func getRandomAvatar(w http.ResponseWriter, r *http.Request) {
	avatar := adorable.Random()
	w.Write(avatar)
}

func getRandomAvatarBasedOnUniqueUserID(w http.ResponseWriter, r *http.Request) {
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
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/getRandomAvatar", getRandomAvatar)
	http.HandleFunc("/getRandomAvatarBasedOnUniqueUserID", getRandomAvatarBasedOnUniqueUserID)
	http.ListenAndServe(":8080", nil)
}
