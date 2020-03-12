# go-rfxcom-command

RFXCom lib to control RF devices

## Deployement on nas

```
$ GOOS=linux GOARCH=arm64 go build .
```

Files on nas are in `/volume2/scripts/rfxcom/`. There is two files:

- `go-rfxcom-command` programs binary.
- `launch.sh` script used to setup nas and launch it, it contains:

```
#!/usr/bin/env bash

set -e

# Enable needed module to use TTY
sudo insmod /lib/modules/usbserial.ko || true
sudo insmod /lib/modules/ftdi_sio.ko || true

/volume2/scripts/rfxcom/go-rfxcom-command > /var/log/go-rfxcom-command.log
```

A service is used to setup launch at startup at `/etc/init/go-rfxcom.conf`.

You can start it:

```
$ start go-rfxcom
```

You can stop it:

```
$ stop go-rfxcom
```
