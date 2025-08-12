package model

type GnbConsoleRegistrationResponse struct {
	Message string  `json:"message"`
	GnbInfo GnbInfo `json:"gnbInfo"`
}

type GnbInfo struct {
	GnbId   string `json:"gnbId"`
	GnbName string `json:"gnbName"`

	PlmnId string `json:"plmnId"`

	Snssai SnssaiIE `json:"snssai"`
}

type SnssaiIE struct {
	Sst string `json:"sst"`
	Sd  string `json:"sd"`
}
