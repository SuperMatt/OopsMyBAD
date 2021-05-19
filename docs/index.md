## Welcome to OopsMyBad

`OopsMyBAD` is a collection of tools used for very opinionated alerting. `Oops` stands for **Ok or panic system** and `MyBAD` means **My Basic Alerting Daemon**.

### Core Principal
The core principal of these tools is to provide a monitoring system which requires very little configuration. There are many wonderful tools out there for monitoring complex systems, but this will suffice for servers that have very simple functionality.

### Oops
`Oops` is a comand wrapper which fires a PagerDuty event, triggering a new incident, when a command exits with a non-zero status. If the command is ever run again and exits with a successful status, it will also resolve the incident.

All it needs is a config file written in Yaml with the PagerDuty Integration Key to run. See the README.md file for more information.

The core usecase is to provide alerts for cronjobs/systed timers without relying on emails.

### MyBAD
`MyBAD` is a long running daemon for monitoring basic parts of your system. The tool is extremely opinionated. It allows you to monitor load, running services, memory and disk usage.

`MyBAD` is configured with a simple Yaml file, which can be built upon the `Oops` config file.