package config

type cfg struct {
	System struct {
		Debug bool `yaml:"debug"`
	} `yaml:"system"`
	Device struct {
		TTY        string `yaml:"tty"`
		Baud       int    `yaml:"baud"`
		BufferSize int    `yaml:"bufSize"`
	} `yaml:"device"`
	MQTT struct {
		DiscoveryPrefix string `yaml:"discoveryPrefix"`
		TopicPrefix     string `yaml:"topicPrefix"`
		ClientID        string `yaml:"clientId"`
		Broker          string `yaml:"broker"`
		Port            int    `yaml:"port"`
		Username        string `yaml:"username"`
		Password        string `yaml:"password"`
	} `yaml:"mqtt"`
}
