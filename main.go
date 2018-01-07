package main

import "log"
import "github.com/craig-duffin/hue-go/huego"
import "github.com/davecgh/go-spew/spew"

func main() {

	client := huego.NewClient("xYP5iv2-B4dAmTKDW9uOXnbG7nYvGpzcDBDe7Ap3")

	lightList, err := client.GetAllLights(false)

	if err != nil {
		spew.Dump(err)
		log.Fatal(err)
	}

	log.Print(lightList)
}
