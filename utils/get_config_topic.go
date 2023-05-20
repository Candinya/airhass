package utils

import (
	"airhass/config"
	"fmt"
)

func GetConfigTopic(id string) string {
	return fmt.Sprintf("%s/sensor/%s_%s/config", config.Config.MQTT.DiscoveryPrefix, config.Config.MQTT.ClientID, id)
}
