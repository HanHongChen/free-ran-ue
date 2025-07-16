package model

type GnbConfig struct {
	Gnb    GnbIE    `yaml:"gnb"`
	Logger LoggerIE `yaml:"logger"`
}

type GnbIE struct {
	AmfN2Ip string `yaml:"amfN2Ip"`
	RanN2Ip string `yaml:"ranN2Ip"`
	UpfN3Ip string `yaml:"upfN3Ip"`
	RanN3Ip string `yaml:"ranN3Ip"`

	RanControlPlaneIp string `yaml:"ranControlPlaneIp"`
	RanDataPlaneIp    string `yaml:"ranDataPlaneIp"`

	AmfN2Port int `yaml:"amfN2Port"`
	RanN2Port int `yaml:"ranN2Port"`
	UpfN3Port int `yaml:"upfN3Port"`
	RanN3Port int `yaml:"ranN3Port"`

	RanControlPlanePort int `yaml:"ranControlPlanePort"`
	RanDataPlanePort    int `yaml:"ranDataPlanePort"`

	NgapPpid uint32 `yaml:"ngapPpid"`

	GnbId   string `yaml:"gnbId"`
	GnbName string `yaml:"gnbName"`

	PlmnId PlmnIdIE `yaml:"plmnId"`

	Tai    TaiIE    `yaml:"tai"`
	Snssai SnssaiIE `yaml:"snssai"`
}
