package main

import (
	"fmt"
	"log"
	"net/http"

	authcontroller "github.com/jintoples/go-session/controllers"
)

func main() {
	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/logout", authcontroller.Logout)

	fmt.Println("Server run on : http://localhost:3000")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
