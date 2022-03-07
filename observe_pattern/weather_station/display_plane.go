package weather_station

import (
	"fmt"
	"log"

	"github.com/PatternDesign-Golang/observe_pattern/observe"
)

type DisplayPlane interface {
	Display()
}

type TemperaturePlane struct {
	temperature float32
}

func (t *TemperaturePlane) Display() {
	fmt.Printf("temperature is %f!", t.temperature)
}

func (t *TemperaturePlane) Update(observable *observe.Observable, data interface{}) {
	weatherData, ok := data.(WeatherData)
	if !ok {
		log.Println(" data type invalid")
	}
	t.temperature = weatherData.Temperature
}

type PressurePlance struct {
}
