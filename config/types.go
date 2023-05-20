package config

type cfg struct {
	System struct {
		Debug bool `yaml:"debug"`
	} `yaml:"system"`
	Device struct {
		Name       string `yaml:"name"`
		Baud       int    `yaml:"baud"`
		BufferSize int    `yaml:"bufSize"`
	} `yaml:"device"`
	Hass struct {
		Address string `yaml:"address"`
	} `yaml:"hass"`
}
