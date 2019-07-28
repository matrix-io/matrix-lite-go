# MATRIX-Lite-Go
[![GoDoc](https://godoc.org/github.com/Hermitter/fileman?status.svg)](https://godoc.org/github.com/matrix-io/matrix-lite-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/matrix-io/matrix-lite-go)](https://goreportcard.com/report/github.com/matrix-io/matrix-lite-go)

MATRIX Lite go is a Golang library that allows users of varying skill levels to easily program their MATRIX Device.

# Roadmap
This roadmap is for achieving a basic implementation of the checklist below. **As this package develops, the API will improve and may change.**
- [x] Leds
- [x] Sensors
  - [x] IMU
  - [x] Humidity
  - [x] Pressure
  - [x] UV
- [x] GPIO
- [ ] Microphones
  - [ ] Hal Mics
  - [ ] Alsa Mics

# Installation
Ensure you have a Raspberry Pi, attached with a MATRIX device, that's flashed with [Raspbian](https://www.raspberrypi.org/downloads/raspbian/).

## 1. Install MATRIX HAL
https://matrix-io.github.io/matrix-documentation/matrix-hal/getting-started/installation-package/

## 2. Install Golang & Create A Project
**Download the latest version of Golang**
```bash
sudo apt install golang
```

## 3. Create A Project
```
mkdir myapp
cd myapp
go mod init myapp
```

## 4. Install matrix-lite-go
```
go get -u github.com/matrix-io/matrix-lite-go
```

# Usage (may change in the future)

## Everloop
```go
package main

import (
	"fmt"
	"time"

	"github.com/matrix-io/matrix-lite-go"
)

func main() {
	m := matrix.Init()

	fmt.Println("This device has", m.Led.Length, "LEDs")

	// A single string or object sets all LEDs
	// Below are different ways of expressing a color (number values are from 0-255)
	m.Led.Set("blue")
	m.Led.Set(matrix.RGBW{0, 0, 10, 0})

	// LEDs off
	m.Led.Set("black")
	m.Led.Set(matrix.RGBW{})

	// Slices & Arrays can set individual LEDs
	m.Led.Set([]interface{}{"red", "", matrix.RGBW{}, "black", matrix.RGBW{G: 255}})// Slice with different data types
	m.Led.Set([]string{"red", "gold", "black", "purple"}) // Slice of strings
	m.Led.Set([5]string{"red", "gold", "black", "purple"}) // Array of strings

	// Slices & Arrays can simulate motion
	// It's recommended to use Slices so that m.Led.Length can be used
	everloop := make([]matrix.RGBW, m.Led.Length)
	everloop[0] = matrix.RGBW{B: 100}

	for {
		lastLed := everloop[0]
		everloop = everloop[1:]
		everloop = append(everloop, lastLed)

		m.Led.Set(everloop)
		time.Sleep(50 * time.Millisecond)
	}
}
```

## Sensors
```go
package main

import (
	"github.com/matrix-io/matrix-lite-go"
)

func main() {
	m := matrix.Init()

	m.Imu.Read()
	m.Uv.Read()
	m.Humidity.Read()
	m.Pressure.Read()

	fmt.Println("IMU: ", m.Imu)
	fmt.Println("UV: ", m.Uv)
	fmt.Println("Humidity: ", m.Humidity)
	fmt.Println("Pressure: ", m.Pressure)
}
```

## GPIO
```go

package main

import (
	"github.com/matrix-io/matrix-lite-go"
)

func main() {
	m := matrix.Init()

	// Read GPIO pin 0 (digital)
	m.Gpio.SetFunction(0, "DIGITAL")
	m.Gpio.SetMode(0, "input")
	fmt.Println(m.Gpio.GetDigital(0))

	// Set GPIO pin 1 (digital)
	m.Gpio.SetFunction(1, "DIGITAL")
	m.Gpio.SetMode(1, "output")
	m.Gpio.SetDigital(1, "ON")

	// Set GPIO pin 2 (PWM)
	m.Gpio.SetFunction(2, "PWM");
	m.Gpio.SetMode(2, "output");
	m.Gpio.SetPWM(2, 25, 50);

	// Set Servo Angle pin 3
	m.Gpio.SetFunction(3, "PWM");
	m.Gpio.SetMode(3, "output");
	m.Gpio.SetServoAngle(3, 90, 0.8);
}
```
