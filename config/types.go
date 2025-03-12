package config

import "airhass/types"

type cfg struct {
	System struct {
		Debug bool `yaml:"debug"`
	} `yaml:"system"`
	Device struct {
		TTY           string `yaml:"tty"`
		Baud          int    `yaml:"baud"`
		BufferSize    int    `yaml:"bufSize"`
		CommaPerBatch int    `yaml:"commaPerBatch"`
	} `yaml:"device"`
	HASS struct {
		DiscoveryPrefix string `yaml:"discoveryPrefix"`
		TopicPrefix     string `yaml:"topicPrefix"`
		DeviceName      string `yaml:"deviceName"`
	} `yaml:"hass"`
	MQTT struct {
		Broker   string `yaml:"broker"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"mqtt"`
	Sensors []types.SensorDiscoveryConfig `yaml:"sensors"`
}
