package services

import "net/http"

type FSMServices interface {
	// Push a command to hardware device
	Push(w http.ResponseWriter, r *http.Request)
}
