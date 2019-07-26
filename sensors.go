package matrix

//#include "sensors.h"
import (
	"C"
)

///////////////////////
// SENSOR BUS INITS //
/////////////////////
func imuInit() Imu {
	C.imu_init()
	return Imu{}
}

func uvInit() Uv {
	C.uv_init()
	return Uv{}
}

func humidityInit() Humidity {
	C.humidity_init()
	return Humidity{}
}

func pressureInit() Pressure {
	C.pressure_init()
	return Pressure{}
}

///////////////////////
//  SENSOR STRUCTS  //
/////////////////////

// Imu sensor values
type Imu struct {
	AccelX, AccelY, AccelZ float32
	GyroX, GyroY, GyroZ    float32
	Yaw, Pitch, Roll       float32
	MagX, MagY, MagZ       float32
}

// Uv sensor values
type Uv struct {
	Uv float32
}

// Humidity sensor values
type Humidity struct {
	Humidity    float32
	Temperature float32
}

// Pressure sensor values
type Pressure struct {
	Altitude    float32
	Pressure    float32
	Temperature float32
}

///////////////////////
//  SENSOR UPDATES  //
/////////////////////

// Update Imu{} values
func (s *Imu) Read() {
	var values C.imu_values = C.imu_read()
		s.AccelX = float32(values.accel_x)
		s.AccelY = float32(values.accel_y)
		s.AccelZ = float32(values.accel_z)
		s.GyroX  = float32(values.gyro_x)
		s.GyroY  = float32(values.gyro_y)
		s.GyroZ  = float32(values.gyro_z)
		s.Yaw    = float32(values.yaw)
		s.Pitch  = float32(values.pitch)
		s.Roll   = float32(values.roll)
		s.MagX   = float32(values.mag_x)
		s.MagY   = float32(values.mag_y)
		s.MagZ   = float32(values.mag_z)
}

// Update Uv{} values
func (s *Uv) Read() {
	var values C.uv_values = C.uv_read()
	s.Uv = float32(values.uv)
}

// Update Humidity{} values
func (s *Humidity) Read() {
	var values C.humidity_values = C.humidity_read()
		s.Humidity    = float32(values.humidity)
		s.Temperature = float32(values.temperature)
}

// Update Pressure{} values
func (s *Pressure) Read() {
	var values C.pressure_values = C.pressure_read()
		s.Altitude    = float32(values.altitude)
		s.Pressure    = float32(values.pressure)
		s.Temperature = float32(values.temperature)
	
}