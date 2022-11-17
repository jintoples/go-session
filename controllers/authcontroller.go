package controllers

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/jintoples/go-session/config"
	"github.com/jintoples/go-session/entities"
	"github.com/jintoples/go-session/models"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Username string
	Password string
}

var userModel = models.NewUserModel()

func Index(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			data := map[string]any{
				"nama_lengkap": session.Values["nama_lengkap"],
			}

			temp, err := template.ParseFiles("views/index.html")
			if err != nil {
				panic(err)
			}

			temp.Execute(w, data)
		}
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/login.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		userInput := &UserInput{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}

		var user entities.User
		userModel.Where(&user, "username", userInput.Username)

		var message error
		if user.Username == "" {
			message = errors.New("Username atau password salah!")
		} else {
			errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))
			if errPassword != nil {
				message = errors.New("Username atau password salah!")
			}
		}

		if message != nil {
			data := map[string]any{
				"error": message,
			}

			temp, err := template.ParseFiles("views/login.html")
			if err != nil {
				panic(err)
			}

			temp.Execute(w, data)
		} else {
			session, _ := config.Store.Get(r, config.SESSION_ID)

			session.Values["loggedIn"] = true
			session.Values["emial"] = user.Email
			session.Values["username"] = user.Username
			session.Values["nama_lengkap"] = user.NamaLengkap

			session.Save(r, w)

			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)

	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
