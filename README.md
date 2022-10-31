## Server Auto Installer
Automation Install Tools for configured server 

Provide by [Fixdigital.NET](http://fixdigital.net) as is. 

#### Download binary application from release
[Download Latest Release](../../releases)

#### OS Supported :
- Ubuntu (Tested on Ubuntu 22.04 LTS)
- Almalinux (Not Tested, Scheduled for Almalinux 9)

##### Support Instalasi
- Auto Update OS
- Auto Upgrade System
- OpenSSH Server
- NTP
- Samba (smbd)
- Apache 2
- Nginx
- PHP 8.1 (php-fpm)
- MySQL 8.0
- phpMyAdmin 5.2.0 en (manual from zip)

#### Run binary with :

> ./autoinstaller-linux-x64
or
> ./autoinstaller-linux-x64 help
for available arguments

_THen follow the instructions._


##### CURRENT ERROR
none


## Build From Source
> golang installed

> go env GOOS=linux GOARCH=amd64 go build -o bin/autoinstaller-linux-x64

> set GOOS=linux GOARCH=amd64 go build -o bin/autoinstaller-linux-x64

_For windows, change GOOS=windows, untuk Mac GOOS=darwin. Pilihan lain android,freebsd,ios,js,openbsd,solaris ._



##### Customize Script :
Rules :
- all requirement file must inside folder "bin" (as root external config).
- the extension for main script is ".the". All ".the" will load as menu.
- Filename of .the is must same with name of script (inside).
- 

