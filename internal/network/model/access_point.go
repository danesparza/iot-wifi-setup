package model

type AccessPoint struct {
	InUse    bool
	BSSID    string
	SSID     string
	Channel  int
	Rate     string
	Signal   int
	Security string
}
