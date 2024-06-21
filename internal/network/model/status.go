package model

type NetworkStatus struct {
	State             string
	Connectivity      string
	Wifi              string
	ActiveConnections []Connection
}
