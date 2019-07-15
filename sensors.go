package lite

//#include "sensors.h"
import (
	"C"
)
import "fmt"

//TODO: prase read_sensor functions into structs

// ReadIMU read MATRIX Creator IMU data
type imu struct {
}

func ReadIMU() {
	fmt.Println(C.imu_read())
}

// ReadUV read MATRIX Creator UV data
func ReadUV() {
	fmt.Println(C.uv_read())
	// return float32(C.uv_read().uv)
}

// ReadHumidity read MATRIX Creator UV data
func ReadHumidity() {
	fmt.Println(C.humidity_read())
}

// ReadPressure read MATRIX Creator UV data
func ReadPressure() {
	fmt.Println(C.pressure_read())
}
