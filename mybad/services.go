package main

import (
	"fmt"
	"os/exec"
)

func (a *App) CheckAllServices() {
	for _, service := range a.Config.Services.Services {
		stateName := fmt.Sprintf("service_%s", service)
		s := NewState(stateName)
		oldState, _ := a.States.LoadOrStore(stateName, s)
		cmd := exec.Command("systemctl", "is-active", service)
		err := cmd.Run()
		if err != nil {
			if cmd.ProcessState.ExitCode() == 3 {
				s = a.Fail(oldState.(State), fmt.Sprintf("%s is inactive", service), a.Config.Services.For)
			} else {
				s = a.Fail(oldState.(State), err.Error(), a.Config.Services.For)
			}
		} else {
			s = a.Success(oldState.(State))
		}

		a.States.Store(stateName, s)
	}

}
