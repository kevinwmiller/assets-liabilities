package auth

import (
	"assets-liabilities/server/routes"
	"net/http"
)

// Router manages routes related to authentication of user objects
type Router struct{}

// List returns a list of authentication route handlers
func (r Router) List() routes.Routes {
	return routes.Routes{
		"/auth/login": routes.Methods{
			routes.Post: login,
		},
		"/auth/logout": routes.Methods{
			routes.Post: logout,
		}}
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login"))
}

func logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout"))
}
