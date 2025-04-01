package driver

import (
	"sync"

	"github.com/kubeedge/mapper-framework/pkg/common"
)

// CustomizedDev is the customized device configuration and client information.
type CustomizedDev struct {
	Instance         common.DeviceInstance
	CustomizedClient *CustomizedClient
}

type CustomizedClient struct {
	// TODO add some variables to help you better implement device drivers
	intMaxValue int
	deviceMutex sync.Mutex
	ProtocolConfig
}

type ProtocolConfig struct {
	ProtocolName string `json:"protocolName"`
	ConfigData   `json:"configData"`
}

type ConfigData struct {
	// TODO: add your protocol config data
	DeviceID   int    `json:"DeviceID,omitempty"`
	SerialPort string `json:"SerialPort"`
	DataBits   int    `json:"DataBits"`
	BaudRate   string `json:"BaudRate"`
	Parity     string `json:"Parity"`
	StopBits   int    `json:"StopBits"`
	ProtocolID int    `json:"ProtocolID"`
}

type VisitorConfig struct {
	ProtocolName      string `json:"protocolName"`
	VisitorConfigData `json:"configData"`
}

type VisitorConfigData struct {
	// TODO: add your visitor config data
	DataType       string `json:"DataType"`
	Register       string `json:"Register"`
	Offset         int    `json:"Offset"`
	Limit          int    `json:"Limit"`
	Scale          int    `json:"Scale"`
	IsSwap         bool   `json:"IsSwap"`
	IsRegisterSwap bool   `json:"IsRegisterSwap"`
}
