#include "matrix_hal/uv_sensor.h"
#include "matrix_hal/uv_data.h"
#include "matrix_hal/imu_sensor.h"
#include "matrix_hal/imu_data.h"
#include "matrix_hal/pressure_sensor.h"
#include "matrix_hal/pressure_data.h"
#include "matrix_hal/humidity_sensor.h"
#include "matrix_hal/humidity_data.h"
#include "bus.hpp"
extern "C"{
    #include "sensors.h"
}

/////////////////////
/////   IMU    /////
matrix_hal::IMUData imu_data;
matrix_hal::IMUSensor imu_sensor;

// setup matrix bus
void imu_init(){
    imu_sensor.Setup(&bus);
}

// get the latest IMU values
imu_values imu_read(){
    imu_values data;

    imu_sensor.Read(&imu_data);
    data.accel_x = imu_data.accel_x;
    data.accel_y = imu_data.accel_y;
    data.accel_z = imu_data.accel_z;
    data.gyro_x = imu_data.gyro_x;
    data.gyro_y = imu_data.gyro_y;
    data.gyro_z = imu_data.gyro_z;
    data.yaw = imu_data.yaw;
    data.pitch = imu_data.pitch;
    data.roll = imu_data.roll;
    data.mag_x = imu_data.mag_x;
    data.mag_y = imu_data.mag_y;
    data.mag_z = imu_data.mag_z;

    return data;
}

/////////////////////
/////   UV     /////
matrix_hal::UVData uv_data;
matrix_hal::UVSensor uv_sensor;

// setup matrix bus
void uv_init(){
    uv_sensor.Setup(&bus);
}

// get the latest UV values
uv_values uv_read(){
    uv_values data;

    uv_sensor.Read(&uv_data);
    data.uv = uv_data.uv;

    return data;
}

/////////////////////
///// HUMIDITY /////
matrix_hal::HumidityData humidity_data;
matrix_hal::HumiditySensor humidity_sensor;

// setup matrix bus
void humidity_init(){
    humidity_sensor.Setup(&bus);
}

// get the latest Humidity values
humidity_values humidity_read(){
    humidity_values data;

    humidity_sensor.Read(&humidity_data);
    data.humidity = humidity_data.humidity;
    data.temperature = humidity_data.temperature;

    return data;
}

/////////////////////
///// PRESSURE /////
matrix_hal::PressureData pressure_data;
matrix_hal::PressureSensor pressure_sensor;

// setup matrix bus
void pressure_init(){
    pressure_sensor.Setup(&bus);
}

// get the latest Pressure values
pressure_values pressure_read(){
    pressure_values data;
    
    pressure_sensor.Read(&pressure_data);
    data.altitude = pressure_data.altitude;
    data.pressure = pressure_data.pressure;
    data.temperature = pressure_data.temperature;

    return data;
}