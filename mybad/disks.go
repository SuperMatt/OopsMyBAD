package main

import (
	"fmt"

	"github.com/SuperMatt/OopsMyBAD/config"
	sigar "github.com/cloudfoundry/gosigar"
)

func (a *App) CheckAllDisks() {
	d := a.Config.Disks
	if len(d) == 0 {
		d = append(d, config.Disk{Path: "/", Threshold: 80, For: 3600})
	}

	for _, disk := range d {
		fs := sigar.FileSystemUsage{}
		fs.Get(disk.Path)

		stateName := fmt.Sprintf("disk_%s", disk.Path)
		s := NewState(stateName)
		oldState, _ := a.States.LoadOrStore(stateName, s)

		if fs.UsePercent() > disk.Threshold {
			s = a.Fail(oldState.(State), fmt.Sprintf("%s at %.2f capacity, more than %.2f full", disk.Path, fs.UsePercent(), disk.Threshold), disk.For)
		} else {
			s = a.Success(oldState.(State))
		}

		a.States.Store(stateName, s)
	}

}
