package model

import "github.com/free5gc/openapi/models"

type UeConfig struct {
	Ue     UeIE     `yaml:"ue" valid:"required"`
	Logger LoggerIE `yaml:"logger" valid:"required"`
}

type UeIE struct {
	RanControlPlaneIp string `yaml:"ranControlPlaneIp" valid:"required"`
	RanDataPlaneIp    string `yaml:"ranDataPlaneIp" valid:"required"`
	LocalDataPlaneIp  string `yaml:"localDataPlaneIp"`

	Ran2DataPlaneIp   string `yaml:"ran2DataPlaneIp" valid:"required"`
	LocalDataPlane2Ip string `yaml:"localDataPlane2Ip"`


	RanControlPlanePort int `yaml:"ranControlPlanePort" valid:"required"`
	RanDataPlanePort    int `yaml:"ranDataPlanePort" valid:"required"`

	PlmnId PlmnIdIE `yaml:"plmnId" valid:"required"`
	Msin   string   `yaml:"msin" valid:"required"`

	AccessType                 models.AccessType            `yaml:"accessType" valid:"required"`
	AuthenticationSubscription AuthenticationSubscriptionIE `yaml:"authenticationSubscription" valid:"required"`

	CipheringAlgorithm CipheringAlgorithmIE `yaml:"cipheringAlgorithm" valid:"required"`
	IntegrityAlgorithm IntegrityAlgorithmIE `yaml:"integrityAlgorithm" valid:"required"`

	PduSessions []PduSessionIE `yaml:"pduSessions" valid:"required"`

	Nrdc NrdcIE `yaml:"nrdc"`

	UeTunnelDevice string `yaml:"ueTunnelDevice" valid:"required"`
}

type AuthenticationSubscriptionIE struct {
	EncPermanentKey               string `yaml:"encPermanentKey" valid:"required"`
	EncOpcKey                     string `yaml:"encOpcKey" valid:"required"`
	AuthenticationManagementField string `yaml:"authenticationManagementField" valid:"required"`
	SequenceNumber                string `yaml:"sequenceNumber" valid:"required"`
}

type IntegrityAlgorithmIE struct {
	Nia0 bool `yaml:"nia0" valid:"required"`
	Nia1 bool `yaml:"nia1" valid:"required"`
	Nia2 bool `yaml:"nia2" valid:"required"`
	Nia3 bool `yaml:"nia3" valid:"required"`
}

type CipheringAlgorithmIE struct {
	Nea0 bool `yaml:"nea0" valid:"required"`
	Nea1 bool `yaml:"nea1" valid:"required"`
	Nea2 bool `yaml:"nea2" valid:"required"`
	Nea3 bool `yaml:"nea3" valid:"required"`
}

type PduSessionIE struct {
	Dnn    string   `yaml:"dnn" valid:"required"`
	Snssai SnssaiIE `yaml:"snssai" valid:"required"`
}

type NrdcIE struct {
	Enable             bool          `yaml:"enable" valid:"required"`
	DcRanDataPlane     DcDataPlaneIE `yaml:"dcRanDataPlane" valid:"required"`
	DcLocalDataPlaneIp string        `yaml:"dcLocalDataPlaneIp"`
}

type DcDataPlaneIE struct {
	Ip   string `yaml:"ip" valid:"required"`
	Port int    `yaml:"port" valid:"required"`
}
