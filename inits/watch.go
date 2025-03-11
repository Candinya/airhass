package inits

import (
	"airhass/config"
	"airhass/global"
	"airhass/jobs"
	"time"
)

func Watch() error {

	msgBuf := make([]byte, config.Config.Device.BufferSize)

	var bufferedMsg string

	// Delay process
	go func() {
		t := time.NewTicker(1 * time.Second)
		for range t.C {
			if bufferedMsg != "" {
				global.Logger.Debug("Send to process:", bufferedMsg)
				jobs.Process(bufferedMsg)
				bufferedMsg = ""
			}
		}
	}()

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

		bufferedMsg += msg

		go jobs.Process(msg)
	}

}
