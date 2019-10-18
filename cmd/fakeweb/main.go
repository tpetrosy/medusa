//Run loging web page for a tetsts
package main

import (
	"net/http"

	"github.com/gorilla/mux"

	common "github.com/tpetrosy/medusa/cmd/fakeweb/internal/common"
)

var router = mux.NewRouter()

func main() {
	router.HandleFunc("/", common.LoginPageHandler) // GET Main Page

	router.HandleFunc("/index", common.IndexPageHandler)             // GET Main Page
	router.HandleFunc("/login", common.LoginHandler).Methods("POST") //POST login page

	router.HandleFunc("/register", common.RegisterPageHandler).Methods("GET") //GET User Registration page
	router.HandleFunc("/register", common.RegisterHandler).Methods("POST")    //POST User Registration page

	router.HandleFunc("/logout", common.LogoutHandler).Methods("POST") //POST Logout

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
