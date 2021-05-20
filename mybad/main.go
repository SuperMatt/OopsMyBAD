package main

import (
	"flag"
	"sync"
	"time"

	"github.com/SuperMatt/OopsMyBAD/config"
)

type App struct {
	Config *config.Config
	States sync.Map
}

func NewApp(configFile *string) *App {
	a := App{Config: config.NewConfig(configFile)}
	return &a
}

func (a *App) doTheThing() {
	a.CheckMemory()
	a.CheckLoad()
	a.CheckAllDisks()
	a.CheckAllServices()
}

func main() {
	config := flag.String("config", "/usr/local/etc/OopsMyBAD/config.yaml", "config file")
	flag.Parse()

	a := NewApp(config)

	i := a.Config.Interval
	if i == 0 {
		i = 60
	}

	ticker := time.NewTicker(time.Duration(i) * time.Second)

	a.doTheThing()
	for range ticker.C {
		a.doTheThing()
	}

}
