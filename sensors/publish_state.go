package sensors

import (
	"airhass/config"
	"airhass/global"
	"airhass/types"
	"airhass/utils"
	"encoding/json"
	"fmt"
)

func PublishState(cfg *types.SensorDiscoveryConfig) {

	// Finish config
	cfg.Platform = "mqtt"
	cfg.StateTopic = utils.GetStateTopic()
	cfg.ValueTemplate = fmt.Sprintf("{{ value_json.%s }}", cfg.ID)
	cfg.EntryID = fmt.Sprintf("%s_%s", config.Config.HASS.DeviceName, cfg.ID)
	cfg.UniqueID = fmt.Sprintf("sensor_airhass_%s_%s", config.Config.HASS.DeviceName, cfg.ID)
	if cfg.DeviceClass != "" {
		cfg.DeviceClassP = &cfg.DeviceClass
	}
	if cfg.Icon != "" {
		cfg.IconP = &cfg.Icon
	}
	cfg.Device.Connections = [][]string{{"tty", config.Config.Device.TTY}}
	cfg.Device.Manufacturer = "AirHass (Nya Candy)"
	cfg.Device.Name = config.Config.HASS.DeviceName

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
