package huego

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"net/http"
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

type Light struct {
	On                        bool       `json:"on"`
	Brightness                uint8      `json:"bri"`
	Hue                       uint16     `json:"hue"`
	Saturation                uint8      `json:"sat"`
	XY                        [2]float32 `json:"xy"`
	ColourTemperature         uint16     `json:"ct"`
	Alert                     string     `json:"alert"`
	Effect                    string     `json:"effect"`
	TransitionTime            uint16     `json:"transitiontime"`
	BrightnessIncrement       int16      `json:"bri_inc"`
	SaturationIncrement       int16      `json:"sat_inc"`
	HueIncrement              int32      `json:"hue_inc"`
	ColourTransitionIncrement int32      `json:"ct_inc"`
	XYIncrement               [2]float32 `json:"xy_inc"`
}

type NewLights struct {
	DeviceIDs []string `json:"deviceid"`
}

func (c *Client) GetAllLights(new bool) (*[]LightAttributes, error) {
	var lights []LightAttributes
	url := fmt.Sprintf(bridgeURL+"%s/lights", c.username)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return &lights, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return &lights, err
	}
	spew.Dump(string(bytes))
	err = json.Unmarshal(bytes, &lights)
	if err != nil {
		return &lights, err
	}
	return &lights, err
}

//TODO
func (c *Client) GetNewLights() (*[]LightAttributes, error) {
	var lights []LightAttributes
	url := fmt.Sprintf(bridgeURL+"/%s/lights/new", c.username)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return &lights, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return &lights, err
	}

	err = json.Unmarshal(bytes, &lights)
	if err != nil {
		return &lights, err
	}
	return &lights, err
}

func (c *Client) SearchNewLights() (*NewLights, error) {
	var newLights NewLights
	url := fmt.Sprintf(bridgeURL+"/%s/lights/new", c.username)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return &newLights, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return &newLights, err
	}

	err = json.Unmarshal(bytes, &newLights)
	if err != nil {
		return &newLights, err
	}
	return &newLights, err
}

func (c *Client) SetLightState() {

}

//TODO
func (c *Client) DeleteLight(lightID string) error {
	url := fmt.Sprintf(bridgeURL+"/%s/lights/%s", c.username, lightID)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return err
	}
	var newLights NewLights
	err = json.Unmarshal(bytes, &newLights)
	if err != nil {
		return err
	}
	return nil
}
