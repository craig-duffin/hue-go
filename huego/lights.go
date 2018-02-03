package huego

import (
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"
)

type LightAttributes struct {
	State             LightState `json:"state"`
	LightType         string     `json:"type"`
	Name              string     `json:"name"`
	ModelID           string     `json:"modelid"`
	UniqueID          string     `json:"uniqueid"`
	ManufacturerName  string     `json:"manufacturername"`
	LuminaireUniqueID string     `json:"luminaireuniqueid"`
	SoftwareVersion   string     `json:"swversion"`
}

type LightIdAndAttributes struct {
	ID string
	Attributes LightAttributes
}

type LightState struct {
	On                bool       `json:"on"`
	Brightness        uint8      `json:"bri"`
	Hue               uint16     `json:"hue"`
	Saturation        uint8      `json:"sat"`
	XY                [2]float32 `json:"xy"`
	ColourTemperature uint16     `json:"ct"`
	Alert             string     `json:"alert"`
	Effect            string     `json:"effect"`
	ColourMode        string     `json:"colormode"`
	Mode              string     `json:"mode"`
	Reachable         bool       `json:"reachable,omitempty"`
}

func (c *Client) GetAllLights() (*[]LightIdAndAttributes, error) {
	var lights []LightIdAndAttributes
	url := fmt.Sprintf(c.bridge+"%s/lights", c.username)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return &lights, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return &lights, err
	}
	lightMap := make(map[string]LightAttributes)

	err = json.Unmarshal(bytes, &lightMap)

	for i, light := range lightMap{
		lights = append(lights, LightIdAndAttributes{ID:i,Attributes:light})
	}

	if err != nil {
		return &lights, err
	}
	return &lights, err
}



func (c *Client) SetLightState(lightID string, newState LightState) error {
	url := fmt.Sprintf(c.bridge+"%s/lights/%s/state", c.username, lightID)
	fmt.Println(url)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(newState)

	req, err := http.NewRequest(http.MethodPut, url, b)

	bytes, err := c.doRequest(req)
	if err != nil {
		return err
	}
	fmt.Println(string(bytes))
	return nil
}
