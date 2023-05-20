package jobs

import (
	"airhass/config"
	"airhass/global"
	"airhass/utils"
	"encoding/json"
	"strings"
)

func Process(csvStr string) {
	// Separate by values
	values := strings.Split(csvStr, ",")

	// Format state
	state := make(map[string]string)
	for _, sensor := range config.Config.Sensors {
		if sensor.DataIndex < len(values) {
			state[sensor.ID] = values[sensor.DataIndex]
		}
	}

	// Parse state into bytes
	stateJsonBytes, err := json.Marshal(&state)
	if err != nil {
		global.Logger.Errorf("Failed to marshal state %v into bytes with error: %v", state, err)
		return
	}

	// Publish state into message
	token := global.MQTT.Publish(utils.GetStateTopic(), 0, false, stateJsonBytes)
	token.Wait()
	if token.Error() != nil {
		global.Logger.Errorf("Failed to publish state %v to MQTT with error: %v", state, err)
		return
	}

	global.Logger.Debug("New state processed successfully.")

}
