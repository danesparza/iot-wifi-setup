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
	NetworkStatus(ctx context.Context) (model.NetworkStatus, error)
	StartAPMode(ctx context.Context, SSIDBaseName, passphrase string) error
	StopAPMode(ctx context.Context, SSID string) error
	ListAccessPoints(ctx context.Context) ([]model.AccessPoint, error)
	SetClientWifiConnection(ctx context.Context, SSID, passphrase string) error
}

type networkManagerService struct {
	// ssid is the full name of the hotspot to use for AP mode
	ssid string

	// apModeEnabled indicates whether AP mode is active or not
	apModeEnabled bool
}

// NetworkStatus shows network status and lists active network connections (if any)
func (n networkManagerService) NetworkStatus(ctx context.Context) (model.NetworkStatus, error) {
	retval := model.NetworkStatus{
		ActiveConnections: make([]model.Connection, 0),
	}

	//	Get general status 🫡
	var stdout, stderr bytes.Buffer
	cmdStatus := exec.CommandContext(ctx, "nmcli", "--terse", "--fields", "state,connectivity,wifi", "general", "status")
	cmdStatus.Stdout = &stdout
	cmdStatus.Stderr = &stderr
	err := cmdStatus.Run()
	if err != nil {
		log.Err(err).Msg("General status failed")
		return retval, fmt.Errorf("problem getting general status: %w", err)
	}

	//	Parse each line of the output
	outputLines := ParseCliOutput(stdout.String())

	//	We should have at least one line
	if len(outputLines) >= 1 {
		if len(strings.TrimSpace(outputLines[0])) > 0 {
			retval = convert.ConvertFieldsToNetworkStatus(ParseCliOutputLine(outputLines[0]))
		}
	} else {
		log.Error().Strs("output", outputLines).Msg("Unexpected output while getting general status")
	}

	//	Collect information about active connections
	var constdout, constderr bytes.Buffer
	cmdList := exec.CommandContext(ctx, "nmcli", "--terse", "--fields", "name,uuid,type,device", "con", "show", "--active")
	cmdList.Stdout = &constdout
	cmdList.Stderr = &constderr
	err = cmdList.Run()
	if err != nil {
		log.Err(err).Msg("Active connections list failed")
		return retval, fmt.Errorf("problem listing active connections: %w", err)
	}

	//	Parse each line of the output
	outputLines = ParseCliOutput(constdout.String())
	for _, line := range outputLines {
		if len(strings.TrimSpace(line)) > 0 {
			connection := convert.ConvertFieldsToActiveConnections(ParseCliOutputLine(line))

			//	Add the connection to the list if the UUID isn't blank
			if strings.TrimSpace(connection.UUID) != "" {
				retval.ActiveConnections = append(retval.ActiveConnections, connection)
			}
		}
	}

	return retval, nil
}

// SetClientWifiConnection updates the local device's network connection settings so
// the device can connect to the specified SSID with the given passphrase
func (n networkManagerService) SetClientWifiConnection(ctx context.Context, SSID, passphrase string) error {
	//	(optional) disconnect from previous wifi network if currently connected
	//	sudo nmcli con down <AP name>

	//	Add a connection to the AP.  This should be set to autoconnect by default
	//	sudo nmcli dev wifi connect "MyNetwork" password "MyPassword123"
	err := exec.CommandContext(ctx, "nmcli", "dev", "wifi", "connect", SSID, "password", passphrase).Run()
	if err != nil {
		log.Err(err).Msg("Problem connecting to wifi")
		return fmt.Errorf("problem connecting to wifi: %w", err)
	}

	//	This will automatically create a file in /etc/NetworkManager/system-connections/
	//	with the AP name, which will contain the configuration.

	//	Reboot
	err = exec.CommandContext(ctx, "reboot").Run()
	if err != nil {
		log.Err(err).Msg("Problem rebooting after setting client wifi connection")
		return fmt.Errorf("problem rebooting after setting client wifi connection: %w", err)
	}

	return nil
}

// StartAPMode starts the device in AP mode so other nearby devices can discover it.  The access point that
// is created uses the ssid as the access point name.  If the passphrase is blank, an open AP is created
func (n networkManagerService) StartAPMode(ctx context.Context, ssid, passphrase string) error {
	if strings.TrimSpace(passphrase) == "" {
		//	Start an AP without a password:
		//	sudo nmcli connection add type wifi con-name open-hotspot autoconnect no wifi.mode ap wifi.ssid "SuperOpenPidy" ipv4.method shared ipv6.method shared
		if err := exec.CommandContext(ctx, "nmcli", "connection", "add", "type", "wifi", "con-name", ssid, "autoconnect", "no", "wifi.mode", "ap", "wifi.ssid", ssid, "ipv4.method", "shared", "ipv6.method", "shared").Run(); err != nil {
			log.Err(err).Str("type", "open").Msg("problem adding connection")
			return fmt.Errorf("problem starting open AP - adding connection: %w", err)
		}

		//	sudo nmcli connection up open-hotspot
		if err := exec.CommandContext(ctx, "nmcli", "connection", "up", ssid).Run(); err != nil {
			log.Err(err).Str("type", "open").Msg("problem setting connection up")
			return fmt.Errorf("problem starting open AP - setting connection up: %w", err)
		}
	} else {
		//	Start an AP with a password:
		// 	sudo nmcli dev wifi hotspot ifname wlan0 ssid test password "test1234"
		err := exec.CommandContext(ctx, "nmcli", "dev", "wifi", "hotspot", "ifname", "wlan0", "ssid", ssid, "password", passphrase).Run()
		if err != nil {
			log.Err(err).Str("type", "secure").Msg("Problem starting AP mode")
			return fmt.Errorf("problem starting AP mode with password: %w", err)
		}
	}

	return nil
}

// StopAPMode stops AP mode (running as the given SSID)
func (n networkManagerService) StopAPMode(ctx context.Context, SSID string) error {
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
