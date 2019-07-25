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
import (
	"fmt"
	"time"

	"github.com/matrix-io/matrix-lite-go"
)

func main() {
	matrix.Init()

	fmt.Println("This device has", matrix.LedLength(), "LEDs")

	// A single string or object sets all LEDs
	// Below are different ways of expressing a color (number values are from 0-255)
	matrix.LedSet("blue")
	matrix.LedSet(matrix.Led{0, 0, 10, 0})

	// LEDs off
	matrix.LedSet("black")
	matrix.LedSet(matrix.Led{})

	// Slices & Arrays can set individual LEDs
	matrix.LedSet([]interface{}{"red", "gold", matrix.Led{}, "black", "purple", matrix.Led{G: 255}})

	// Slices & Arrays can simulate motion
	everloop := make([]matrix.Led, matrix.LedLength())
	everloop[0] = matrix.Led{B: 100}

	for {
		lastLed := everloop[0]
		everloop = everloop[1:]
		everloop = append(everloop, lastLed)

		matrix.LedSet(everloop)
		time.Sleep(50 * time.Millisecond)
	}
}
```

## Sensors
```go
    package main
    import (
        "github.com/matrix-io/matrix-lite-go"
        "fmt"
    )
    
    func main(){
        matrix.Init()
        fmt.Println( matrix.ReadImu() )
        fmt.Println( matrix.ReadUv() )
        fmt.Println( matrix.ReadHumidity() )
        fmt.Println( matrix.ReadPressure() )
    }
```
