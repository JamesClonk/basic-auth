package main

import (
	"crypto/subtle"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", basicAuth)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func basicAuth(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("WWW-Authenticate", `Basic realm="Basic-Auth"`)

	username := os.Getenv("AUTH_USERNAME")
	password := os.Getenv("AUTH_PASSWORD")
	if len(username) < 3 && len(password) < 8 {
		http.Error(rw, "Unauthorized", 401)
		return
	}

	user, pass, _ := req.BasicAuth()
	if subtle.ConstantTimeCompare([]byte(user), []byte(username)) == 1 &&
		subtle.ConstantTimeCompare([]byte(pass), []byte(password)) == 1 {
		rw.WriteHeader(200)
		_, _ = rw.Write([]byte("OK"))
		return
	}
	http.Error(rw, "Unauthorized", 401)
	return
}
