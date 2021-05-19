# MyBAD - My Basic Alerting Daemon
MyBAD is a VERY opinionated alerting daemon. You shouldn't need to configure very much.

## Usage
Use as a long running service in whatever way you fancy.

MyBAD expects a config.yaml file in the directory from which it is run, or for --config to be passed to it.

```
name: your hostname
mybad --config /etc/oops_mybad.yaml
```

## Config
MyBAD can use the same config file as Oops.

Below is an example config:

```
name: your.hostname
pagerduty:
    apikey: yourkey
interval: 60
disks:
  - path: /
    threshold: 80
    for: 3600
load:
  threshold: 2
  for: 300
memory:
  threshold: 80
  for: 300
services:
  services:
    - lighttpd
    - pihole-FTL
  for: 300
```

Below is a description of all the config options:

### Root

| key | required | description |
| --- | -------- | ----- |
| name | n | The hostname of the server running MyBAD. Defaults to the system hostname |
| pagerduty | y | All pagerduty related config |
| interval | n | How often to check in seconds. Defaults to 60|
| disks | n | A list of disks to monitor. Defaults to `/` |
| load | n | Load average check |
| memory | n | Memory check |
| services | n | Services to check |

### Disks
| key | required | description |
| --- | -------- | ----------- |
| path | y | The mountpoint to check |
| threshold | n | The alert threshold in %. Defaults to 80 |
| for | n | How long in second to be above the threshold before alerting. Defaults to 3600 |
 
### Load
| key | required | description |
| --- | -------- | ----------- |
| threshold | n | The alert threshold. Defaults to the number of CPUs on your host |
| for | n | How long in seconds to be above the threshold before alerting. Defaults to 300 |

### Memory
| key | required | description |
| --- | -------- | ----------- |
| threshold | n | The alert threshold in %. Defaults to 80 |
| for | n | How long in seconds to be above the threshold before alerting. Defaults to 300 |

### Services
| key | required | description |
| --- | -------- | ----------- |
| services | n | A list of services to monitor. Simply checks if they are active or not |
| for | n | How long in seconds to be above the threshold before alerting. Defaults to 300 |