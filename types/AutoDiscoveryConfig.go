package types

type AutoDiscoveryConfig struct {
	ID string `json:"-"`

	Name              string `json:"name"`
	UnitOfMeasurement string `json:"unit_of_measurement"`
	DeviceClass       string `json:"-"`

	Platform      string  `json:"platform"` // mqtt
	StateTopic    string  `json:"state_topic"`
	DeviceClassP  *string `json:"device_class,omitempty"`
	ValueTemplate string  `json:"value_template"`
	// UniqueID      string  `json:"unique_id"`
	// EntryID string `json:"entry_id"`
}
