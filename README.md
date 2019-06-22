![jft](logo/jft.jpg)

# calendar cli tool, just for today
This tool is a calendar that helps you plan and keep running every day.  
If you plan daily, it will surely be for you.

## Installing(requires Go 1.11.4)
```shell
$ brew install mysql // if you didn't install MySQL 
$ export GO111MODULE=auto
$ go get -u github.com/NasSilverBullet/jft/cmd/jft
```

## Setting
```shell
$ mysql.server start
$ mysql -u root mysql -e 'create database jft;'
```

## Usage
```shell
$ jft [command]
```

## Available Commands
```shell
add         Add today's new plans
list        show plans list
update      update today's each plan, please give me id
delete      delete today's each plan, please give me id
month       show monthly calendar
year        show yearly calendar
help        Help about any command
``` 

## Flags
```shell
 -h, --help   help for autospirit
```

## You can do this for example with the following command
### add
```shell
$ jft add 10:00 12:00 'check emails'
added a new plan!!
ID : 1
Start : 2019/06/22 10:00
End : 2019/06/22 12:00
Title : check emails
Description : 

$ jft add 13:00 14:30 'meeting' -d 'on conference room 10' // You can add detailed description
added a new plan!!
ID : 2
Start : 2019/06/22 13:00
End : 2019/06/22 14:30
Title : meeting
Description : on conference room 10
```
### list
```shell
$ jft list // You can check your today's plans
ID : 1
Start : 2019/06/22 10:00
End : 2019/06/22 12:00
Title : check emails
Description : 

ID : 2
Start : 2019/06/22 13:00
End : 2019/06/22 14:30
Title : meeting
Description : on conference room 10

$ jft list -w 2019/06/01 // You can check your each day's plans
```

### update
```shell
$ jft update 2 -s 13:30 -e 15:00 // You can update your plans
updated the plan!!
ID : 2
Start : 2019/06/22 13:30
End : 2019/06/22 15:00
Title : meeting
Description : on conference room 10
```

### delete
```shell
$ jft list
ID : 1  // check ID
Start : 2019/06/22 01:00
End : 2019/06/22 03:00
Title : check emails
Description : 
~
$ jft delete 1
deleted the plan!!
ID : 1
Start : 2019/06/22 01:00
End : 2019/06/22 03:00
Title : check emails
Description : 
````

### month
```shell
jft month // You can check your efforts on this month
2019/06/01 (Sat)  >>>  X
2019/06/02 (Sun)  >>>  X
2019/06/03 (Mon)  >>>  O
2019/06/04 (Tue)  >>>  O
2019/06/05 (Wed)  >>>  O
2019/06/06 (Thu)  >>>  O
2019/06/07 (Fri)  >>>  O
2019/06/08 (Sat)  >>>  O
2019/06/09 (Sun)  >>>  O
2019/06/10 (Mon)  >>>  X
2019/06/11 (Tue)  >>>  O
2019/06/12 (Wed)  >>>  O
2019/06/13 (Thu)  >>>  X
2019/06/14 (Fri)  >>>  O
2019/06/15 (Sat)  >>>  X
2019/06/16 (Sun)  >>>  X
2019/06/17 (Mon)  >>>  O
2019/06/18 (Tue)  >>>  X
2019/06/19 (Wed)  >>>  O
2019/06/20 (Thu)  >>>  X
2019/06/21 (Fri)  >>>  X
2019/06/22 (Sat)  >>>  O
2019/06/23 (Sun)  >>>  -
2019/06/24 (Mon)  >>>  -
2019/06/25 (Tue)  >>>  -
2019/06/26 (Wed)  >>>  -
2019/06/27 (Thu)  >>>  -
2019/06/28 (Fri)  >>>  -
2019/06/29 (Sat)  >>>  -
2019/06/30 (Sun)  >>>  -

$ jft month -w 2019/05 // You can check your efforts on each month
```

### year
```shell
$ jft year // You can check your efforts on this year
2019/01  >>>  0(0.00%)
2019/02  >>>  0(0.00%)
2019/03  >>>  0(0.00%)
2019/04  >>>  0(0.00%)
2019/05  >>>  0(0.00%)
2019/06  >>>  1(0.03%)
2019/07  >>>  --------
2019/08  >>>  --------
2019/09  >>>  --------
2019/10  >>>  --------
2019/11  >>>  --------
2019/12  >>>  --------

$ jft year -w 2018 // You can check your efforts on each year
```
