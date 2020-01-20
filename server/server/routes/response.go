package routes

import (
	"assets-liabilities/errors"
	"net/http"
)

// ResponseMessage is used to give context to non-200 error codes
type ResponseMessage struct {
	message string
}

// Respond sets the response header and body in the given ReponseWriter
func Respond(w http.ResponseWriter, status int, responseJSON []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(responseJSON)
}

// RespondWithError sets the response header and adds the given message to the body under a 'message' attribute
func RespondWithError(w http.ResponseWriter, errorWithCode *errors.ErrorWithCode) {
	http.Error(w, errorWithCode.Message, errorWithCode.Code)
}
