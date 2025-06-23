package model

type GnbConfig struct {
	Gnb    GnbIE    `yaml:"gnb"`
	Logger LoggerIE `yaml:"logger"`
}

type GnbIE struct {
	AmfN2Ip string `yaml:"amfN2Ip"`
	GnbN2Ip string `yaml:"gnbN2Ip"`

	AmfN2Port int `yaml:"amfN2Port"`
	GnbN2Port int `yaml:"gnbN2Port"`

	NgapPpid uint32 `yaml:"ngapPpid"`

	GnbId   string `yaml:"gnbId"`
	GnbName string `yaml:"gnbName"`

	PlmnId PlmnIdIE `yaml:"plmnId"`

	Tai    TaiIE    `yaml:"tai"`
	Snssai SnssaiIE `yaml:"snssai"`
}

type PlmnIdIE struct {
	Mcc string `yaml:"mcc"`
	Mnc string `yaml:"mnc"`
}

type TaiIE struct {
	Tac             string   `yaml:"tac"`
	BroadcastPlmnId PlmnIdIE `yaml:"broadcastPlmnId"`
}

type SnssaiIE struct {
	Sst string `yaml:"sst"`
	Sd  string `yaml:"sd"`
}

type LoggerIE struct {
	Level string `yaml:"level"`
}
