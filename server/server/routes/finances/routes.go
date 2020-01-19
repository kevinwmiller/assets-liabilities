package finances

import (
	"assets-liabilities/server/routes"
	"net/http"
)

// Router manages routes related to authentication of user objects
type Router struct{}

// List returns a list of authentication route handlers
func (r Router) List() routes.Routes {
	return routes.Routes{
		"/finances/records": routes.Methods{
			routes.Get:  listRecords,
			routes.Post: createRecord,
		},
		"/finances/records/{id}": routes.Methods{
			routes.Get:    getRecord,
			routes.Put:    updateRecord,
			routes.Delete: deleteRecord,
		},
	}
}

func listRecords(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List Records"))
}

func createRecord(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Record"))
}

func getRecord(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Record"))
}

func updateRecord(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Record"))
}

func deleteRecord(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Record"))
}
