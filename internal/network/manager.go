package network

import "github.com/danesparza/iot-wifi-setup/internal/network/model"

// https://www.networkmanager.dev/docs/api/latest/nmcli-examples.html
type NetworkManagerService interface {
	ListAccessPoints() ([]model.AccessPoint, error)
}

type networkManagerService struct {
}

func (n networkManagerService) ListAccessPoints() ([]model.AccessPoint, error) {
	//TODO implement me
	panic("implement me")
}

func NewNetworkManagerService() NetworkManagerService {
	svc := networkManagerService{}

	//	Initialize the service?

	return svc
}
