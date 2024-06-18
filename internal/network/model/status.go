package model

type NetworkStatus struct {
	State             string
	Connectivity      string
	Wifi              string
	activeConnections []Connection
}
