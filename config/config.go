package config

import (
	"io/ioutil"
	"os"

	sigar "github.com/cloudfoundry/gosigar"
	"github.com/go-yaml/yaml"
)

type PagerDuty struct {
	ApiKey string `yaml:"apikey"`
}

type BasicThreshold struct {
	Threshold float64 `yaml:"threshold"`
	For       float64 `yaml:"for"`
}

type Disk struct {
	Threshold float64 `yaml:"threshold"`
	For       float64 `yaml:"for"`
	Path      string  `yaml:"path"`
}

type Config struct {
	Name      string         `yaml:"name"`
	Interval  float64        `yaml:"interval"`
	PagerDuty PagerDuty      `yaml:"pagerduty"`
	Load      BasicThreshold `yaml:"load"`
	Memory    BasicThreshold `yaml:"memory"`
	Disks     []Disk         `yaml:"disks"`
	Services  Services       `yaml:"services"`
}

type Services struct {
	Services []string `yaml:"services"`
	For      float64  `yaml:"for"`
}

func NewConfig(config *string) *Config {
	configYaml, err := ioutil.ReadFile(*config)
	if err != nil {
		panic(err)
	}

	var c Config
	err = yaml.Unmarshal(configYaml, &c)
	if err != nil {
		panic(err)
	}

	if c.Name == "" {
		hostname, err := os.Hostname()
		if err != nil {
			panic(err)
		}
		c.Name = hostname
	}

	if c.Load.For == 0 {
		c.Load.For = 300
	}

	if c.Load.Threshold == 0 {
		cpuList := sigar.CpuList{}
		cpuList.Get()
		c.Load.Threshold = float64(len(cpuList.List))
	}

	if c.Memory.For == 0 {
		c.Memory.For = 300
	}
	if c.Memory.Threshold == 0 {
		c.Memory.Threshold = 80
	}

	if c.Services.For == 0 {
		c.Services.For = 300
	}

	for i, d := range c.Disks {
		if d.For == 0 {
			c.Disks[i].For = 3600
		}
		if d.Threshold == 0 {
			c.Disks[i].Threshold = 80
		}
	}

	return &c

}
