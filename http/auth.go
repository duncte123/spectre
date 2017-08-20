package http

import (
	"net/http"

	"howett.net/spectre"
)

type LoginService interface {
	GetLoggedInUser(r *http.Request) spectre.User
	SetLoggedInUser(w http.ResponseWriter, r *http.Request, u spectre.User)
}

type PermitterProvider interface {
	GetPermitterForRequest(r *http.Request) spectre.Permitter
}
