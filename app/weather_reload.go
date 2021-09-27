package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type weather struct {
	Status status `json:"status"`
}

type status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func ChangeWeatherStatus() {
	fmt.Println("updating weather status ...")
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100

	weatherStatus := new(weather)
	weatherStatus.Status.Water = rand.Intn(max-min) + min
	weatherStatus.Status.Wind = rand.Intn(max-min) + min

	weatherStatusJson, _ := json.MarshalIndent(weatherStatus, "", "    ")

	PATH, _ := os.Getwd()

	_ = ioutil.WriteFile(PATH+"/public/weather-status.json", weatherStatusJson, 0666)
	fmt.Println(string(weatherStatusJson))
}

func Reloader(function func(), delay time.Duration) chan bool {
	ticker := time.NewTicker(delay)
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				function()
			case <-stop:
				fmt.Println("Ticker stopped")
				ticker.Stop()
				return
			}
		}
	}()

	return stop
}
