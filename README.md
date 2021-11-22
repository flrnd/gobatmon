# gobatmon

gobatmon is a simple command line tool to monitor battery discharges during a period of time.

The original idea was to create a systemd service that logged the charge at shutdown and then calculate the discharge percentage during the time the laptop was off. 
Right now you can create timestamps (Storing date and current charge), list saved timestamps and calculate discharge since last saved timestamp. It also prints basic information stats (like number of cycles, full design capacity, etc).

examples:

```shell
$ gobatmon stats

Manufacturer: HP
Status: Discharging
Energy full design : 70.07 Wh
Energy full        : 70.07 Wh
Energy now         : 41.40 Wh
Energy rate        : 7.90 W
Charge             : 59%
Cycle count: 25
```

```shell
$ gobatmon create
created timestamp (7) 59% at 2021-11-22 14:39:50
```

```shell
$ gobatmon list  
id: 1 charge: 59% created: 2021-11-19 20:38:55
id: 2 charge: 61% created: 2021-11-20 11:05:27
id: 3 charge: 60% created: 2021-11-20 11:08:32
```

```shell
$ gobatmon last
Discharge      : 23%
Time elapsed   : 4h47m38.972305457s
Discharge ratio: 4.798Wh
```

## TODO

- [ ] Create project with proper issues and tasks.
- [ ] Write systemd services for shutdown and boot.
- [ ] Implement data schema for future graphs and analytics.
- [ ] Add tests.
