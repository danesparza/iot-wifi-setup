# iot-wifi-setup [![Build and release](https://github.com/danesparza/iot-wifi-setup/actions/workflows/release.yaml/badge.svg)](https://github.com/danesparza/iot-wifi-setup/actions/workflows/release.yaml)
Present a wifi AP to get connected to local wifi network and then hand-off to your app.  This uses [Network Manager](https://www.networkmanager.dev/) (available in newer versions of [Raspberry Pi OS](https://www.raspberrypi.com/software/) and Linux)

## Installation
Install the package repo (you only need to do this once per machine)
```
wget https://packages.cagedtornado.com/prereq.sh -O - | sh
```

Install the package
```
sudo apt install iot-wifi-setup
```
