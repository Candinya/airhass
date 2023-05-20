package utils

import (
	"airhass/config"
	"fmt"
)

func GetConfigTopic(id string) string {
	return fmt.Sprintf("%s/sensor/%s_%s/config", config.Config.HASS.DiscoveryPrefix, config.Config.HASS.DeviceName, id)
}
