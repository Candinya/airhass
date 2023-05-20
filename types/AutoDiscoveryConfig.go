package types

type SensorDiscoveryConfig struct {
	// Custom filled
	ID string `json:"-" yaml:"id"`

	DataIndex int `json:"-" yaml:"dataIndex"`

	Name              string `json:"name" yaml:"name"`
	UnitOfMeasurement string `json:"unit_of_measurement" yaml:"unitOfMeasurement"`
	DeviceClass       string `json:"-" yaml:"deviceClass"`
	Icon              string `json:"-" yaml:"icon"`

	// Automatically generated
	Platform      string  `json:"platform"` // mqtt
	StateTopic    string  `json:"state_topic"`
	DeviceClassP  *string `json:"device_class,omitempty"`
	IconP         *string `json:"icon,omitempty"`
	ValueTemplate string  `json:"value_template"`
	UniqueID      string  `json:"unique_id"`
	EntryID       string  `json:"object_id"`

	Device struct {
		Name         string     `json:"name"`
		Connections  [][]string `json:"connections"`
		Manufacturer string     `json:"manufacturer"`
	} `json:"device"`
}
