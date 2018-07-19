package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/d2r2/go-dht"
	rpio "github.com/stianeikeland/go-rpio"
)

const maxTemp = 34
const maxHumidity = 75
const minTemp = 30

type Fan struct {
	status bool
	pin    rpio.Pin
}

func (f *Fan) On() {
	f.status = true
	f.pin.High()
}

func (f *Fan) Off() {
	f.status = false
	f.pin.Low()
}

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()

	fan := Fan{
		status: false,
		pin:    rpio.Pin(10),
	}
	fan.pin.Output() // Output mode
	for {
		// Read DHT11 sensor data from pin 4, retrying 10 times in case of failure.
		// You may enable "boost GPIO performance" parameter, if your device is old
		// as Raspberry PI 1 (this will require root privileges). You can switch off
		// "boost GPIO performance" parameter for old devices, but it may increase
		// retry attempts. Play with this parameter.
		temperature, humidity, retried, err := dht.ReadDHTxxWithRetry(dht.DHT11, 4, true, 10)
		if err != nil {
			log.Println(err)
			continue
		}
		// Print temperature and humidity
		log.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
			temperature, humidity, retried)

		if fan.status && temperature < minTemp {
			fan.On()
		} else if !fan.status && (temperature > maxTemp || humidity > maxHumidity) && temperature > minTemp {
			fan.Off()
		}
		time.Sleep(time.Second)
	}
}
