package driver

import (
	"sync"

	"github.com/kubeedge/usb/pkg/common"
)

// CustomizedDev is the customized device configuration and client information.
type CustomizedDev struct {
	Instance         common.DeviceInstance
	CustomizedClient *CustomizedClient
}

type CustomizedClient struct {
	// TODO add some variables to help you better implement device drivers
	deviceMutex sync.Mutex
	ProtocolConfig
}

type ProtocolConfig struct {
	ProtocolName string `json:"protocolName"`
	ConfigData   `json:"configData"`
}

type ConfigData struct {
	// TODO: add your protocol config data
	SerialPort string `json:"serialPort"`

	DeviceID int `json:"deviceID,omitempty"`
	Width    int `json:"width,omitempty"`
	Height   int `json:"height,omitempty"`
	Format   int `json:"format,omitempty"`

	ProtocolID int `json:"protocolID"`
}

type VisitorConfig struct {
	ProtocolName      string `json:"protocolName"`
	VisitorConfigData `json:"configData"`
}

type VisitorConfigData struct {
	// TODO: add your visitor config data
	DataType    string `json:"dataType"`
	FeatureName string `json:"featureName"`
}
