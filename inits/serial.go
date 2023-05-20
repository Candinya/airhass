package inits

import (
	"airhass/config"
	"airhass/global"

	"github.com/tarm/serial"
)

func Serial() error {
	// Initialize device configuration
	serialCfg := &serial.Config{
		Name: config.Config.Device.Name,
		Baud: config.Config.Device.Baud,
	}

	// Open port with config
	var err error
	global.SerialPort, err = serial.OpenPort(serialCfg)
	if err != nil {
		return err
	}

	// Successfully opened serial port
	return nil

}
