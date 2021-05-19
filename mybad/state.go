package main

import (
	"fmt"
	"time"

	"github.com/SuperMatt/OopsMyBAD/pd"
)

type State struct {
	Name      string
	OK        bool
	Message   string
	LastOK    time.Time
	LastCheck time.Time
	LastFail  time.Time
	Alerted   bool
}

func NewState(name string) State {
	return State{Name: name, OK: true, LastCheck: time.Now(), LastOK: time.Now()}
}

func (a *App) Success(oldState State) State {
	s := NewState(oldState.Name)
	s.LastOK = time.Now()
	s.LastFail = oldState.LastFail
	s.OK = true
	if oldState.Alerted {
		r, err := pd.Event(a.Config.PagerDuty.ApiKey, "resolve", s.Name, oldState.Message, a.Config.Name)
		if err != nil {
			fmt.Println(err)
			s.Alerted = true
		} else {
			fmt.Println(r.Status)
			s.Alerted = false
		}
	}

	return s
}

func (a *App) Fail(oldState State, message string, for_time float64) State {
	s := NewState(oldState.Name)
	s.LastFail = time.Now()
	s.LastOK = oldState.LastOK
	s.Alerted = oldState.Alerted
	s.Message = message
	s.OK = false
	if s.LastFail.Sub(s.LastOK) > time.Duration(for_time)*time.Second && !oldState.Alerted {
		r, err := pd.Event(a.Config.PagerDuty.ApiKey, "trigger", s.Name, s.Message, a.Config.Name)

		if err != nil {
			fmt.Println(err)
			s.Alerted = false
		} else {
			fmt.Println(r.Status)
			s.Alerted = true
		}

	}

	return s

}

func (s State) String() string {
	return fmt.Sprintf("Name: %s, OK: %t, Message: %s", s.Name, s.OK, s.Message)
}
