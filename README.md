# go-rfxcom-command

RFXCom lib to control RF devices

## Requirements

- `Go`
- [`reflex`](https://github.com/cespare/reflex)

```
$ make deps
```

## Development

### Build

You can build project

```
$ make build
```

#### For nas

You can build project for nas

```
$ make build-arm64
```

### Watch

You can watch project

```
$ make watch
```

## Docker

You can build container for all architecture using
[buildx](https://docs.docker.com/buildx/working-with-buildx/):

```
ocker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t itsalex/go-rfxcom --push .
```

## Deployement on nas

```
$ GOOS=linux GOARCH=arm64 go build .
```

Files on nas are in `/volume2/scripts/rfxcom/`. There is two files:

- `go-rfxcom-command` programs binary.
- `launch.sh` script used to setup nas and launch it, it contains:

```
#!/usr/bin/env bash

set -v

# Enable needed module to use TTY
sudo insmod /lib/modules/usbserial.ko || true
sudo insmod /lib/modules/ftdi_sio.ko || true

/volume2/scripts/rfxcom/go-rfxcom-command
```

A service is used to setup launch at startup at `/etc/init/go-rfxcom.conf`.

File service:

```
author "ItsAlex"
start on runlevel 2
stop on runlevel [06]

expect fork
respawn
respawn limit 5 10

pre-start script
    echo "Starting go-rfxcom"
end script

post-stop script
    echo "Stopped go-rfxcom"
end script

exec /volume2/scripts/rfxcom/launch.sh
```

You can start it:

```
$ start go-rfxcom
```

You can stop it:

```
$ stop go-rfxcom
```

You can see logs at `cat /var/log/upstart/go-rfxcom.log`
