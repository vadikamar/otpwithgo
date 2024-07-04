package main

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"log"
	"math/big"
	"net/http"
)

func otp() *big.Int {
	RandomCrypto, _ := rand.Prime(rand.Reader, 18)
	return RandomCrypto
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	otp := otp().String() // Generate OTP and convert to string

	fmt.Println(otp)
	tmpl := template.Must(template.ParseFiles("myproject\\templates\\index.html"))
	tmpl.Execute(w, map[string]string{
		"OTP": otp,
	})
}

func main() {
	//fs := http.FileServer(http.Dir("myproject\\templates"))
	http.HandleFunc("/", serveHTML)

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
