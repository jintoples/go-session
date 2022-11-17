package config

import (
	"os"

	"github.com/gorilla/sessions"
)

const SESSION_ID = "go_auth_sess"

var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
