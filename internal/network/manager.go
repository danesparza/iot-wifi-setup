package network

import (
	"bytes"
	"context"
	"fmt"
	"github.com/danesparza/iot-wifi-setup/internal/convert"
	"github.com/danesparza/iot-wifi-setup/internal/network/model"
	"github.com/rs/zerolog/log"
	"os/exec"
	"strings"
)

type NetworkManagerService interface {
	ActiveConnections(ctx context.Context) ([]model.Connection, error)
	StartAPMode(ctx context.Context, SSIDBaseName, passphrase string, passwordless bool) error
	ListAccessPoints(ctx context.Context) ([]model.AccessPoint, error)
	UpdateLocalWifiSettings(ctx context.Context, SSID, passphrase string) error
}

type networkManagerService struct{}

// ActiveConnections shows active network connections.  If there are
// no wifi connections, AP mode should be started
func (n networkManagerService) ActiveConnections(ctx context.Context) ([]model.Connection, error) {
	//	sudo nmcli con show --active
	//	Should probably also include information about general status:
	// 	sudo nmcli --terse --fields state,connectivity,wifi general status
	//TODO implement me
	panic("implement me")
}

// UpdateLocalWifiSettings updates the local device's network connection settings so
// the device can connect to the specified SSID with the given passphrase
func (n networkManagerService) UpdateLocalWifiSettings(ctx context.Context, SSID, passphrase string) error {
	//	(optional) disconnect from previous wifi network if currently connected
	//	sudo nmcli con down <AP name>

	//	Connect to the new AP
	//	nmcli device wifi connect <AP name> password <password>

	//	This will automatically create a file in /etc/NetworkManager/system-connections/
	//	with the AP name, which will contain the configuration.

	//TODO implement me
	panic("implement me")
}

// StartAPMode starts the device in AP mode so other nearby devices can discover it.  The access point that
// is created uses the SSIDBaseName as the first part of the access point name.
func (n networkManagerService) StartAPMode(ctx context.Context, SSIDBaseName, passphrase string, passwordless bool) error {
	//	With a password it's simpler:
	// 	sudo nmcli dev wifi hotspot ifname wlan0 ssid test password "test1234"

	//	Passwordless is apparently possible and takes a few more commands:
	//	sudo nmcli connection add type wifi ifname $WIFI_INTERFACE con-name $AP autoconnect yes ssid $AP
	//	sudo nmcli connection modify $AP 802-11-wireless.mode ap 802-11-wireless.band bg ipv4.method shared
	//	sudo nmcli connection modify $AP wifi-sec.key-mgmt none
	//	sudo nmcli connection up $AP
	//	sudo nmcli connection modify $AP connection.autoconnect yes

	//TODO implement me
	panic("implement me")
}

// ListAccessPoints lists the nearby access points
func (n networkManagerService) ListAccessPoints(ctx context.Context) ([]model.AccessPoint, error) {
	retval := make([]model.AccessPoint, 0)

	//	Create the command with context and request a wifi rescan
	cmdRescan := exec.CommandContext(ctx, "nmcli", "device", "wifi", "rescan")
	err := cmdRescan.Run()
	if err != nil {
		log.Err(err).Msg("Rescan failed")
		return retval, fmt.Errorf("problem rescanning: %w", err)
	}

	var stdout, stderr bytes.Buffer
	cmdList := exec.CommandContext(ctx, "nmcli", "--terse", "--fields", "in-use,bssid,ssid,chan,rate,signal,security", "dev", "wifi", "list")
	cmdList.Stdout = &stdout
	cmdList.Stderr = &stderr
	err = cmdList.Run()
	if err != nil {
		log.Err(err).Msg("List failed")
		return retval, fmt.Errorf("problem listing aps: %w", err)
	}

	//	Parse each line of the output
	outputLines := ParseCliOutput(stdout.String())
	for _, line := range outputLines {
		if len(strings.TrimSpace(line)) > 0 {
			ap := convert.ConvertFieldsToAccessPoint(ParseCliOutputLine(line))

			//	Add the AP to the list if the SSID isn't blank
			if strings.TrimSpace(ap.SSID) != "" {
				retval = append(retval, ap)
			}
		}
	}

	return retval, nil
}

// NewNetworkManagerService creates a new network manager service
func NewNetworkManagerService() NetworkManagerService {
	svc := networkManagerService{}
	return svc
}
