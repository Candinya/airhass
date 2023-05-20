package utils

import (
	"airhass/config"
	"fmt"
)

func GetStateTopic() string {
	// return fmt.Sprintf("%s/state/%s", config.Config.MQTT.TopicPrefix, id)
	return fmt.Sprintf("%s/sensor/%s/state", config.Config.HASS.DiscoveryPrefix, config.Config.HASS.DeviceName)
}
