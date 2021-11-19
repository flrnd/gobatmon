# gobatmon

gobatmon is a simple command line tool that I created to monitor and log battery discharges during long period of  storage.

The idea is to create a systemd service that logs the charge at shutdown and then calculate the discharge percentage during the time the laptop was off.

It also prints basic information like number of cycles, full design capacity, etc.

example:

```shell
$ gobatmon stats

Manufacturer: HP
Status: Discharging
Full design capacity: 70070 mWh
Full charge capacity: 70070 mWh
Current charge at: 59% | Discharge rate of 3.71 W
Cycle count: 23
```

## TODO

- [] Create project with proper issues and tasks.
- [] Write systemd services for shutdown and boot.
- [] Implement data schema for future graphs and analytics.
- [] Add tests.
