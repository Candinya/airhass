package inits

import (
	"airhass/config"
	"airhass/global"
	"airhass/jobs"
)

func Watch() error {

	readBuf := make([]byte, config.Config.Device.BufferSize)
	var sendBuf []byte

	for {
		msgLen, err := global.SerialPort.Read(readBuf)
		if err != nil {
			global.Logger.Errorf("Failed to get data with error: %v", err)
			return err
		}

		global.Logger.Debug("---------------------------------------")
		global.Logger.Debugf("Serial data arrive (%d bytes):", msgLen)
		msg := readBuf[:msgLen]

		global.Logger.Debug(string(msg))

		sendBuf = append(sendBuf, msg...)

		// Count if is complete
		commaCount := 0
		for _, c := range sendBuf {
			if c == ',' {
				commaCount++
			}
		}

		if commaCount == config.Config.Device.CommaPerBatch {
			global.Logger.Debugf("Message batch match (%d comma), process", commaCount)
			global.Logger.Debug(string(sendBuf))
			jobs.Process(string(sendBuf))
			// Clear send buffer after batch finish
			sendBuf = nil
		} else if commaCount > config.Config.Device.CommaPerBatch {
			global.Logger.Debugf("Message corrupted (%d comma), discard", commaCount)
			global.Logger.Debug(string(sendBuf))
			// Discard
			sendBuf = nil
		} // else message is incomplete, continue accumulating

	}

}
