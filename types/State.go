package types

type State struct {
	Particulates03 string `json:"particulates03"`
	PM25           string `json:"pm25"`
	Formaldehyde   string `json:"formaldehyde"`
	CarbonDioxide  string `json:"carbon_dioxide"`
	VOC            string `json:"voc"`
}
