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
}

type LoggerIE struct {
	Level string `yaml:"level"`
}
