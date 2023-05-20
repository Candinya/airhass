package sensors

import (
	"airhass/global"
	"airhass/types"
	"airhass/utils"
	"encoding/json"
	"fmt"
)

func PublishState(cfg *types.AutoDiscoveryConfig) {

	// Finish config
	cfg.Platform = "mqtt"
	cfg.StateTopic = utils.GetStateTopic()
	cfg.ValueTemplate = fmt.Sprintf("{{ value_json.%s }}", cfg.ID)
	// cfg.EntryID = fmt.Sprintf("sensor.%s_%s", config.Config.MQTT.ClientID, cfg.ID)
	if cfg.DeviceClass != "" {
		cfg.DeviceClassP = &cfg.DeviceClass
	}

	// Construct config json
	configJsonBytes, err := json.Marshal(cfg)
	if err != nil {
		global.Logger.Errorf("Failed to marshal config %v into bytes with error: %v", cfg, err)
		return
	}

	// Publish config
	token := global.MQTT.Publish(utils.GetConfigTopic(cfg.ID), 0, false, configJsonBytes)
	token.Wait()
	if token.Error() != nil {
		global.Logger.Errorf("Failed to publish config %v to MQTT with error: %v", cfg, token.Error())
	}
}

var SensorsList = []types.AutoDiscoveryConfig{{
	ID:                "particulates03",
	Name:              "> 0.3μm 颗粒物",
	UnitOfMeasurement: "PCS/0.1L",
}, {
	ID:                "pm25",
	Name:              "PM 2.5 浓度",
	UnitOfMeasurement: "μg/m³",
	DeviceClass:       "pm25",
}, {
	ID:                "formaldehyde",
	Name:              "甲醛浓度",
	UnitOfMeasurement: "µg/m³",
}, {
	ID:                "carbon_dioxide",
	Name:              "CO₂ 浓度",
	UnitOfMeasurement: "ppm",
	DeviceClass:       "carbon_dioxide",
}, {
	ID:                "voc",
	Name:              "VOC",
	UnitOfMeasurement: "ppb",
}}
