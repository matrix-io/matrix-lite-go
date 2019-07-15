package lite

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
	return ImuValues{
		AccelX: float32(C.imu_read().accel_x),
		AccelY: float32(C.imu_read().accel_y),
		AccelZ: float32(C.imu_read().accel_z),
		GyroX:  float32(C.imu_read().gyro_x),
		GyroY:  float32(C.imu_read().gyro_y),
		GyroZ:  float32(C.imu_read().gyro_z),
		Yaw:    float32(C.imu_read().yaw),
		Pitch:  float32(C.imu_read().pitch),
		Roll:   float32(C.imu_read().roll),
		MagX:   float32(C.imu_read().mag_x),
		MagY:   float32(C.imu_read().mag_y),
		MagZ:   float32(C.imu_read().mag_z),
	}
}

// UvValues contains UV sensor values
type UvValues struct {
	UV float32
}

// ReadUV read MATRIX Creator UV data
func ReadUV() UvValues {
	return UvValues{
		UV: float32(C.uv_read().uv),
	}
}

// HumidityValues contains Humidity sensor values
type HumidityValues struct {
	Humidity    float32
	Temperature float32
}

// ReadHumidity read MATRIX Creator Humidity data
func ReadHumidity() HumidityValues {
	return HumidityValues{
		Humidity:    float32(C.humidity_read().humidity),
		Temperature: float32(C.humidity_read().temperature),
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
	return PressureValues{
		Altitude:    float32(C.pressure_read().altitude),
		Pressure:    float32(C.pressure_read().pressure),
		Temperature: float32(C.pressure_read().temperature),
	}
}
