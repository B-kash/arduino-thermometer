package main

import (
	"fmt"
	"machine"
	"time"

	"tinygo.org/x/drivers/dht"
)

func main() {
	pin := machine.D6
	dhtSensor := dht.New(pin, dht.DHT22)
	var temp float32
	var hum float32
	for {
		err := dhtSensor.ReadMeasurements()
		if err != nil {
			fmt.Println("Mesurement failed %s\n", err.Error())
		} else {
			temp, err = dhtSensor.TemperatureFloat(0)
			hum, err = dhtSensor.HumidityFloat()
			if err != nil {
				fmt.Println("could not get humidity and temprature")
				break
			}
			now := time.Now()
			fmt.Println(now)
			fmt.Println("temprature:", temp)
			fmt.Println("humidity:", hum)
		}
		time.Sleep(time.Second * 5)
	}
}
