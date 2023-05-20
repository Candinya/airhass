package jobs

import (
	"airhass/global"
	"airhass/types"
	"airhass/utils"
	"encoding/json"
	"strings"
)

func Process(csvStr string) {
	// Separate by values
	values := strings.Split(csvStr, ",")

	// global.Logger.Debugf("> 0.3 μm 颗粒数:\t%s\t个/0.1L", values[0])
	// global.Logger.Debugf("PM 2.5 浓度值:\t\t%s\tμg/m³", values[1])
	// global.Logger.Debugf("甲醛浓度值:\t\t%s\tμg/m³", values[2])
	// global.Logger.Debugf("CO₂ 浓度值:\t\t%s\tppm", values[3])
	// global.Logger.Debugf("VOC:\t\t\t%s\tppb", values[6])

	// Format state
	state := types.State{
		Particulates03: values[0],
		PM25:           values[1],
		Formaldehyde:   values[2],
		CarbonDioxide:  values[3],
		VOC:            values[6],
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
