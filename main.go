package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/d2r2/go-dht"
	rpio "github.com/stianeikeland/go-rpio"
)

const fan1Pin = 10
const fan2Pin = 9
const sensor1Pin = 4
const sensor2Pin = 5
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

	fan1 := Fan{
		status: false,
		pin:    rpio.Pin(fan1Pin),
	}
	fan1.pin.Output() // Output mode

	fan2 := Fan{
		status: false,
		pin:    rpio.Pin(fan2Pin),
	}
	fan2.pin.Output() // Output mode
	for {
		// Read DHT11 sensor data from pin 4, retrying 10 times in case of failure.
		// You may enable "boost GPIO performance" parameter, if your device is old
		// as Raspberry PI 1 (this will require root privileges). You can switch off
		// "boost GPIO performance" parameter for old devices, but it may increase
		// retry attempts. Play with this parameter.
		temperature1, humidity1, retried, err := dht.ReadDHTxxWithRetry(dht.DHT11, sensor1Pin, true, 10)
		if err != nil {
			log.Println(err)
			continue
		}

		// Print temperature and humidity
		log.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
			temperature1, humidity1, retried)

		if fan1.status && temperature1 < minTemp {
			fan1.On()
		} else if !fan1.status && (temperature1 > maxTemp || humidity1 > maxHumidity) && temperature1 > minTemp {
			fan1.Off()
		}

		temperature2, humidity2, retried, err := dht.ReadDHTxxWithRetry(dht.DHT11, sensor2Pin, true, 10)
		if err != nil {
			log.Println(err)
			continue
		}

		// Print temperature and humidity
		log.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
			temperature2, humidity2, retried)

		if fan2.status && temperature2 < minTemp {
			fan2.On()
		} else if !fan2.status && (temperature2 > maxTemp || humidity2 > maxHumidity) && temperature2 > minTemp {
			fan2.Off()
		}

		time.Sleep(time.Second * 10)
	}
}
