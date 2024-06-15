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

// https://www.networkmanager.dev/docs/api/latest/nmcli-examples.html
type NetworkManagerService interface {
	ListAccessPoints(ctx context.Context) ([]model.AccessPoint, error)
}

type networkManagerService struct{}

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
