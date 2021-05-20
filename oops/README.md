# Oops - Ok or page system
Oops is a command wrapper which triggers a PagerDutyV2 integration event if it exits with a non-zero status, creating an incident. If the same command is run again and exists successfully, it resolves the original incident.

## Usage
Oops expects a config.yaml file in the directory from which it is run, or for --config to be passed to it.

```
oops --config /etc/oops_mybad.yaml ./command.sh
```

## Config
Oops can use the same config file as MyBAD, and therefore the location of the api key should be the same. It should look like this:

```
name: your.hostname
pagerduty:
    apikey: yourkey
```