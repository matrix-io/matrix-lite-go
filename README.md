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
- [ ] GPIO
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

	matrix "github.com/matrix-io/matrix-lite-go"
)

func main() {
	m := matrix.Init()

	fmt.Println("This device has", m.Led.Length, "LEDs")

	// A single string or object sets all LEDs
	// Below are different ways of expressing a color (number values are from 0-255)
	m.Led.Set("blue")
	m.Led.Set(matrix.Rgbw{0, 0, 10, 0})

	// LEDs off
	m.Led.Set("black")
	m.Led.Set(matrix.Rgbw{})

	// Slices & Arrays can set individual LEDs
	m.Led.Set([]interface{}{"red", "gold", matrix.Rgbw{}, "black", matrix.Rgbw{G: 255}})

	// Slices & Arrays can simulate motion
	everloop := make([]matrix.Rgbw, m.Led.Length)
	everloop[0] = matrix.Rgbw{B: 100}

	for {
		lastLed := everloop[0]
		everloop = everloop[1:]
		everloop = append(everloop, lastLed)

		m.Led.Set(everloop)
		time.Sleep(10 * time.Millisecond)
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
