package weather_station

import (
	"log"

	"github.com/PatternDesign-Golang/observe_pattern/observe"
)

type WeatherData struct {
	Temperature float32
	Humanity    float32
	Pressure    float32
}

type WeatherStation struct {
	observable observe.Observable
	data       WeatherData
}

func (w *WeatherStation) RegisterDisplayPlane(observer observe.Observer) {
	err := w.observable.RegisterObservers(observer)
	if err != nil {
		log.Println(err)
		return
	}
}

func (w *WeatherStation) RemoveDisplayPlane(observer observe.Observer) {
	err := w.observable.RemoveObservers(observer)
	if err != nil {
		log.Println(err)
		return
	}
}

func (w *WeatherStation) NotifyDisplayPlane(observer observe.Observer) {
	w.observable.NotifyObservers(observer, w.data)
}
