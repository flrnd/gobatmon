# gobatmon

[![Go Report Card](https://goreportcard.com/badge/github.com/flrnd/gobatmon)](https://goreportcard.com/report/github.com/flrnd/gobatmon)

gobatmon is a command line tool to monitor battery discharges during a period of time.

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

$ gobatmon list saved
id: 1 | lastId: 1
Discharge    : 1%
From         : 2021-12-20 16:32:42
To           : 2021-12-20 16:37:29
Time elapsed : 4m46.298353448s
Ratio        : 12.574Wh

id: 2 | lastId: 1
Discharge    : 4%
From         : 2021-12-20 16:32:42
To           : 2021-12-20 16:56:16
Time elapsed : 23m34.063440814s
Ratio        : 10.183Wh

```

```shell
$ gobatmon last
Discharge      : 23%
Time elapsed   : 4h47m38.972305457s
Discharge ratio: 4.798Wh

$ gobatmon last save
Saved (2) 4% elapsed: 23m34.063440814s dr: 10.183Wh

```

