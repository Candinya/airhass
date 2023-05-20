package inits

import (
	"airhass/config"
	"airhass/global"
	"airhass/jobs"
)

func Watch() error {

	msgBuf := make([]byte, config.Config.Device.BufferSize)

	for {
		msgLen, err := global.SerialPort.Read(msgBuf)
		if err != nil {
			global.Logger.Errorf("Failed to get data with error: %v", err)
			return err
		}

		global.Logger.Debug("---------------------------------------")
		global.Logger.Debugf("Serial data arrive (%d bytes):", msgLen)
		msg := string(msgBuf[:msgLen])

		global.Logger.Debug(msg)

		// Send to processor
		go jobs.Process(msg)
	}

}
