package fsm

import (
	"fmt"
	"context"
	"time"
	"math/rand"
	"regexp"
	"errors"
)

var (
	randomer = rand.New(rand.NewSource(99))
	regexer, _ = regexp.Compile("deploy:([0-9]+)")
)

type State struct {
	Name string
	ETF  int64 // estimate to finish in mili-seconds
}

type FiniteStateMachine struct {
	initialState State
	terminalState State
	currentState State
	transitions map[string]func(int) string
	states map[string]State
	cmd string
}


func (f *FiniteStateMachine) addStates(states... State) {
	for _, state := range states {
		f.states[string(state.Name)] = state
	}
}

func (f *FiniteStateMachine) addTransition(source State, handler func(int) string) {
	f.transitions[string(source.Name)] = handler
}

func (f *FiniteStateMachine) turnON() {
	fmt.Println("Starting machine...")
	f.applyState(ON)
}

func (f *FiniteStateMachine) PushCommand(command string) (int, error) {
	if !regexer.MatchString(command) {
		return 400, errors.New("Wrong format!")
	}
	if f.currentState.Name != STAND_BY.Name {
		return 403, errors.New("Device is not ready to receive command!")
	}
	f.cmd = command
	f.applyState(COMMAND_RECEIVE)
	return 204, nil
}

func (f *FiniteStateMachine) applyState(state State) {
	// apply
	f.currentState = state

	if f.currentState == COMMAND_BROADCASTING {
		fmt.Println("Broadcasting... " + f.cmd)
	}

	// delay to job finished
	delayCtx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(state.ETF) * time.Millisecond)
	defer cancelFunc()
	<-delayCtx.Done()

	// keep transiting states
	f.transitState()
}

func (f *FiniteStateMachine) transitState() {
	if f.currentState.Name == STAND_BY.Name || f.currentState.Name == STOP.Name {

	}
	nextState := f.states[f.transitions[string(f.currentState.Name)](randomer.Int())]
	switch f.currentState.Name {
		case STAND_BY.Name:
		case OFF.Name:
			break;
		default:
			go func() {
				f.applyState(nextState)
			}()
			break;
	}
}

func NewFSM(init, terminal State) *FiniteStateMachine {
	return &FiniteStateMachine{
		initialState: init,
		terminalState: terminal,
		transitions: map[string]func(int)string{},
		states: map[string]State{},
	}
}

// assume running a state machine for a blue-tooth device to broadcast command (1-way command)
// transition #1: ON -> STARTING -> STAND-BY
//
// transition #2: STAND-BY -> COMMAND-RECEIVE -> COMMAND-ENCRYPTING -> COMMAND-BROADCASTING -> STAND-BY
//                                                                  -> COMMAND-ERROR -> STAND-BY
//
// transition #3: STOP -> SHUTTING-DOWN -> OFF
var (
	ON = State{"ON", 500 }
	OFF = State {"OFF", 0 }
	STARTING = State {"STARTING", 1500 }
	STAND_BY = State{ "STAND-BY", 0 }
	COMMAND_RECEIVE = State{"COMMAND-RECEIVE", 0 }
	COMMAND_ENCRYPTING= State{"COMMAND-ENCRYPTING", 500 }
	COMMAND_BROADCASTING = State{"COMMAND-BROADCAST", 1000 }
	COMMAND_ERROR = State { "COMMAND-ERROR", 200 }
	STOP = State {"STOP", 200 }
	SHUTTING_DOWN = State {"SHUTTING-DOWN", 2000 }

	// FSM instance
	FSM *FiniteStateMachine
)

func init() {

	// setup for the machine
	FSM = NewFSM(ON, OFF)
	FSM.addStates(STARTING, STAND_BY, COMMAND_RECEIVE, COMMAND_ENCRYPTING, COMMAND_BROADCASTING, COMMAND_ERROR, STOP, SHUTTING_DOWN)
	FSM.addTransition(ON, func(i int) string {return STARTING.Name})
	FSM.addTransition(STARTING, func(i int) string {
		fmt.Println("Machine is ready to use!")
		return STAND_BY.Name
	})
	FSM.addTransition(STAND_BY, func(i int) string {return STAND_BY.Name})
	FSM.addTransition(COMMAND_RECEIVE, func(i int) string {return COMMAND_ENCRYPTING.Name})
	FSM.addTransition(COMMAND_ENCRYPTING, func(i int) string {
		if i > 10 { return COMMAND_BROADCASTING.Name }
		return COMMAND_ERROR.Name
	}) // simulate 10% fail to encrypt command
	FSM.addTransition(COMMAND_BROADCASTING, func(i int) string {return STAND_BY.Name})
	FSM.addTransition(COMMAND_ERROR, func(i int) string {return STAND_BY.Name})
	FSM.addTransition(STOP, func(i int) string {return SHUTTING_DOWN.Name})
	FSM.addTransition(SHUTTING_DOWN, func(i int) string {return OFF.Name})

	FSM.turnON()
}