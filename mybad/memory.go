package main

import (
	"fmt"

	sigar "github.com/cloudfoundry/gosigar"
)

func (a *App) CheckMemory() {
	mem := sigar.Mem{}
	mem.Get()

	stateName := "memory"
	s := NewState(stateName)
	oldState, _ := a.States.LoadOrStore(stateName, s)

	if float64(100)*float64(mem.ActualUsed)/float64(mem.Total) > a.Config.Memory.Threshold {
		s = a.Fail(oldState.(State), fmt.Sprintf("Used memory %d greater than %.2f%% of total %d", mem.ActualUsed, a.Config.Memory.Threshold, mem.Total), a.Config.Memory.For)
	} else {
		s = a.Success(oldState.(State))
	}

	a.States.Store(stateName, s)

}
