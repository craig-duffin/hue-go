package main

import "log"
import "fmt"
import "time"
import "github.com/craig-duffin/hue-go/huego"
import "flag"

func main() {

	userPtr := flag.String("user", "","An authorized username returned by the bridge")
	bridgePtr := flag.String("bridge", "","The bridge IP address")

	flag.Parse()

	if *userPtr == "" {
		log.Fatal("user is a required flag")
	}

	if *bridgePtr == "" {
		log.Fatal("bridge is a required flag")
	}
	
	client := huego.NewClient(*userPtr,"http://"+*bridgePtr+"/api/")

	lightStateMap := make(map[string]huego.LightState)

	for range time.Tick(time.Second){
		lightList, err := client.GetAllLights()

		if err != nil {
			log.Fatal(err)
		}

		 for _, light := range *lightList {
		 	if isDefaultState(light.Attributes.State) == false {
		 		
				if lightStateMap[light.Attributes.Name] != light.Attributes.State{
					lightStateMap[light.Attributes.Name] = light.Attributes.State
					fmt.Printf("Stored new state for light %s \n", light.Attributes.Name)
				}
				 
				 
		 	} else {
		 		fmt.Printf("The light %s is the default state, sending previous state \n", light.Attributes.Name)
		 		err := client.SetLightState(light.ID, lightStateMap[light.Attributes.Name])
	
		 		if err != nil {
		 			log.Fatal(err)
		 		}
		 	}
		 }

	}

}


func isDefaultState(ls huego.LightState) bool {

	xy := [2]float32{0.4573, 0.41}

	defaultLightState := huego.LightState{
		On:                true,
		Brightness:        254,
		Hue:               8418,
		Saturation:        140,
		XY:                xy,
		ColourTemperature: 366,
		Alert:             "select",
		Effect:            "none",
		ColourMode:        "ct",
		Mode:              "homeautomation",
		Reachable:         true,
	}

	if ls == defaultLightState {
		return true
	}
	return false
}
