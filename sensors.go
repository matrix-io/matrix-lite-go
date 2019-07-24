package matrix

//#include "sensors.h"
import (
	"C"
)

// Imu sensor values
type Imu struct {
	AccelX, AccelY, AccelZ float32
	GyroX, GyroY, GyroZ    float32
	Yaw, Pitch, Roll       float32
	MagX, MagY, MagZ       float32
}

// ReadImu returns a populated Imu struct
func ReadImu() Imu {
	var values C.imu_values = C.imu_read()
	return Imu{
		AccelX: float32(values.accel_x),
		AccelY: float32(values.accel_y),
		AccelZ: float32(values.accel_z),
		GyroX:  float32(values.gyro_x),
		GyroY:  float32(values.gyro_y),
		GyroZ:  float32(values.gyro_z),
		Yaw:    float32(values.yaw),
		Pitch:  float32(values.pitch),
		Roll:   float32(values.roll),
		MagX:   float32(values.mag_x),
		MagY:   float32(values.mag_y),
		MagZ:   float32(values.mag_z),
	}
}

// Uv sensor values
type Uv struct {
	UV float32
}

// ReadUv returns a populated Uv struct
func ReadUv() Uv {
	var values C.uv_values = C.uv_read()
	return Uv{
		UV: float32(values.uv),
	}
}

// Humidity sensor values
type Humidity struct {
	Humidity    float32
	Temperature float32
}

// ReadHumidity returns a populated Humidity struct
func ReadHumidity() Humidity {
	var values C.humidity_values = C.humidity_read()
	return Humidity{
		Humidity:    float32(values.humidity),
		Temperature: float32(values.temperature),
	}
}

// Pressure sensor values
type Pressure struct {
	Altitude    float32
	Pressure    float32
	Temperature float32
}

// ReadPressure returns a populated Pressure struct
func ReadPressure() Pressure {
	var values C.pressure_values = C.pressure_read()
	return Pressure{
		Altitude:    float32(values.altitude),
		Pressure:    float32(values.pressure),
		Temperature: float32(values.temperature),
	}
}
