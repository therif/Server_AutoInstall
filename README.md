## Server Auto Installer
Automation Install Tools for configured server.  
Provide by us as is. 

### Download binary application from release
[Download Latest Release](../../releases)


#### Internal software/scripts feature :
- Ubuntu (Tested on Ubuntu 22.04 LTS)
- Almalinux (Not Tested, Scheduled for Almalinux 9)

    ##### Internal Scripts Automation Avalable
    - Auto Update OS
    - Auto Upgrade System
    - OpenSSH Server
    - NTP
    - Samba (smbd)
    - Apache 2
    - Nginx
    - PHP 8.1 (php-fpm)
    - MySQL 8.0
    - phpMyAdmin 5.2.0 english



#### Run binary with :

```./ai-linux-x64```

> `./ai-linux-x64 help`


_Then follow the instructions._



## Build From Source
> golang installed

> `go env GOOS=linux GOARCH=amd64 go build -o bin/ai-linux-x64`

> `set GOOS=linux GOARCH=amd64 go build -o bin/ai-linux-x64`

**change GOOS=target-os**  
_**list target-os**_ : `linux` `windows` `darwin` `freebsd` `android` `ios` `js` `aix` `dragonfly` `hurd` `illumos` `nacl` `netbsd` `openbsd` `plan9` `solaris` `zos`

**change GOARCH=target-arch**  
_**target-arch**_ : `386` `amd64` `amd64p32` `arm` `arm64` `arm64be` `armbe` `loong64` `mips` `mips64` `mips64le` `mips64p32` `mips64p32le` `mipsle` `ppc` `ppc64` `ppc64le` `riscv` `riscv64` `s390` `s390x` `sparc` `sparc64` `wasm`

## Customize Script :
Rules :
- all requirement file must inside folder "bin" (as root external config).  
- the extension for main script is ".the". All ".the" will load as menu.  
- Filename of .the is must same with name of script (inside).  
- Filename will auto convert space to underscore ( _ ).
- put the binari/executable in same folder with all script and requirement files.
