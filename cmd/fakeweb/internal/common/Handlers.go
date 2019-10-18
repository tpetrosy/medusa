package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/securecookie"
	helpers "github.com/tpetrosy/medusa/cmd/fakeweb/internal/helpers"
	repos "github.com/tpetrosy/medusa/cmd/fakeweb/internal/repos"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

// Handlers

//LoginPageHandler for GET Main Page
func LoginPageHandler(response http.ResponseWriter, request *http.Request) {
	//Session validation on login page
	userName := GetUserName(request)
	if !helpers.IsEmpty(userName) {
		http.Redirect(response, request, "/index", 302)
	} else {
		var body, _ = helpers.LoadFile("internal/templates/login.html")
		fmt.Fprintf(response, body)
	}
}

//LoginHandler for POST Main
func LoginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	pass := request.FormValue("password")
	redirectTarget := "/"
	if !helpers.IsEmpty(name) && !helpers.IsEmpty(pass) {
		// Database check for user data!
		_userIsValid := repos.UserIsValid(name, pass)

		if _userIsValid {
			SetCookie(name, response)
			redirectTarget = "/index"
		} else {
			redirectTarget = "/register"
		}
	}
	http.Redirect(response, request, redirectTarget, 302)
}

//RegisterPageHandler for GET
func RegisterPageHandler(response http.ResponseWriter, request *http.Request) {
	var body, _ = helpers.LoadFile("internal/templates/register.html")
	fmt.Fprintf(response, body)
}

//RegisterHandler for POST
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	uName := r.FormValue("username")
	email := r.FormValue("email")
	pwd := r.FormValue("password")
	confirmPwd := r.FormValue("confirmPassword")

	_uName, _email, _pwd, _confirmPwd := false, false, false, false
	_uName = !helpers.IsEmpty(uName)
	_email = !helpers.IsEmpty(email)
	_pwd = !helpers.IsEmpty(pwd)
	_confirmPwd = !helpers.IsEmpty(confirmPwd)

	if _uName && _email && _pwd && _confirmPwd {
		fmt.Fprintln(w, "Username for Register : ", uName)
		fmt.Fprintln(w, "Email for Register : ", email)
		fmt.Fprintln(w, "Password for Register : ", pwd)
		fmt.Fprintln(w, "ConfirmPassword for Register : ", confirmPwd)
	} else {
		fmt.Fprintln(w, "This fields can not be blank!")
	}
}

//IndexPageHandler for GET
func IndexPageHandler(response http.ResponseWriter, request *http.Request) {
	userName := GetUserName(request)
	if !helpers.IsEmpty(userName) {
		var indexBody, _ = helpers.LoadFile("internal/templates/index.html")
		fmt.Fprintf(response, indexBody, userName)
	} else {
		http.Redirect(response, request, "/", 302)
	}
}

//LogoutHandler for POST
func LogoutHandler(response http.ResponseWriter, request *http.Request) {
	ClearCookie(response)
	http.Redirect(response, request, "/", 302)
}

//SetCookie Cookie
func SetCookie(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("cookie", value); err == nil {
		cookie := &http.Cookie{
			Name:  "cookie",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

//ClearCookie remove the cookie
func ClearCookie(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "cookie",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

//GetUserName Get User Name
func GetUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("cookie"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("cookie", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}
