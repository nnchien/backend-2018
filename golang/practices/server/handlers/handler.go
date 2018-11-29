package handlers

import (
	"github.com/nnchien/backend-2018/golang/practices/internal/fsm"
	"github.com/nnchien/backend-2018/golang/practices/services"
)

type Handler struct {
	fsm *fsm.FiniteStateMachine
}

func NewHandler() services.FSMServices {
	return &Handler{
		fsm: fsm.FSM,
	}
}
