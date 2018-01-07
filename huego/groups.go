package huego

type Group struct {
	Name      string   `json:"name"`
	Lights    []string `json:"lights"`
	GroupType string   `json:"group_type"`
	Action    Action   `json:"action,omitempty"`
}

type Action struct {
	On                bool       `json:"on"`
	Brightness        uint8      `json:"bri"`
	Hue               uint16     `json:"hue"`
	Saturation        uint8      `json:"sat"`
	XY                [2]float32 `json:"xy"`
	ColourTemperature uint16     `json:"ct"`
	Alert             string     `json:"alert"`
	Effect            string     `json:"effect"`
	ColourMode        string     `json:"colormode"`
}

type GroupAttributes struct {
	Action     Action   `json:"action"`
	Lights     []string `json:"lights"`
	Name       string   `json:"name"`
	GroupType  string   `json:"type"`
	ModelID    string   `json:"modelid"`
	UniqueID   string   `json:"uniqueid"`
	GroupClass string   `json:"class"`
}
