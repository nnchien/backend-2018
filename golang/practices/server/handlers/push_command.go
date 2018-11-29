package handlers

import (
	"net/http"
	"fmt"
)

func (h *Handler) Push(w http.ResponseWriter, r *http.Request) {
	code, err := h.fsm.PushCommand(r.FormValue("command"))
	w.WriteHeader(code)
	switch code {
		case 400:
		case 403:
			fmt.Fprintf(w, "%v", err)
		break;
	}
}
