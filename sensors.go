package matrix

//#include "sensors.h"
import (
	"C"
)

// ImuValues contains IMU sensor values
type ImuValues struct {
	AccelX, AccelY, AccelZ float32
	GyroX, GyroY, GyroZ    float32
	Yaw, Pitch, Roll       float32
	MagX, MagY, MagZ       float32
}

// ReadIMU read MATRIX Creator IMU data
func ReadIMU() ImuValues {
	var values C.imu_values = C.imu_read()
	return ImuValues{
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

// UvValues contains UV sensor values
type UvValues struct {
	UV float32
}

// ReadUV read MATRIX Creator UV data
func ReadUV() UvValues {
	var values C.uv_values = C.uv_read()
	return UvValues{
		UV: float32(values.uv),
	}
}

// HumidityValues contains Humidity sensor values
type HumidityValues struct {
	Humidity    float32
	Temperature float32
}

// ReadHumidity read MATRIX Creator Humidity data
func ReadHumidity() HumidityValues {
	var values C.humidity_values = C.humidity_read()
	return HumidityValues{
		Humidity:    float32(values.humidity),
		Temperature: float32(values.temperature),
	}
}

// PressureValues contains Pressure sensor values
type PressureValues struct {
	Altitude    float32
	Pressure    float32
	Temperature float32
}

// ReadPressure read MATRIX Creator UV data
func ReadPressure() PressureValues {
	var values C.pressure_values = C.pressure_read()
	return PressureValues{
		Altitude:    float32(values.altitude),
		Pressure:    float32(values.pressure),
		Temperature: float32(values.temperature),
	}
}
