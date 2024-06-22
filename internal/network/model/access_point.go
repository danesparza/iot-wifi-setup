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

type APModeRequest struct {
	SSID       string `json:"ssid"`
	Passphrase string `json:"passphrase"`
}
