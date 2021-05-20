package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/SuperMatt/OopsMyBAD/config"
	"github.com/SuperMatt/OopsMyBAD/pd"
)

func main() {
	configFile := flag.String("config", "/usr/local/etc/OopsMyBAD/config.yaml", "config file")
	flag.Parse()

	c := config.NewConfig(configFile)

	command := flag.Args()

	cmd := exec.Command(command[0], command[1:]...)
	out, err := cmd.CombinedOutput()
	fmt.Print(string(out))
	message := ""
	if err != nil {
		if err == exec.ErrNotFound {
			message = err.Error()
		} else {
			message = string(out)
		}

		_, err := pd.Event(c.PagerDuty.ApiKey, "trigger", strings.Join(command, "_"), message, c.Name)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		_, err := pd.Event(c.PagerDuty.ApiKey, "resolve", strings.Join(command, "_"), string(out), c.Name)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	os.Exit(cmd.ProcessState.ExitCode())
}
