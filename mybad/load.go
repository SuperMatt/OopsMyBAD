package main

import (
	"fmt"

	sigar "github.com/cloudfoundry/gosigar"
)

func (a *App) CheckLoad() State {
	load := sigar.LoadAverage{}
	load.Get()

	stateName := "load"
	s := NewState(stateName)
	oldState, _ := a.States.LoadOrStore(stateName, s)

	if load.One > a.Config.Load.Threshold {
		s = a.Fail(oldState.(State), fmt.Sprintf("Load %.2f greater than %.2f", load.One, a.Config.Load.Threshold), a.Config.Load.For)
	} else {
		s = a.Success(oldState.(State))
	}

	return s

}
