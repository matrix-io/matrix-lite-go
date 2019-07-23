#ifndef SENSORS_H
#define SENSORS_H

/////////////////////
/////   IMU    /////
void imu_init();

typedef struct {
    float accel_x, accel_y, accel_z;
    float gyro_x,  gyro_y,  gyro_z;
    float yaw,     pitch,   roll;
    float mag_x,   mag_y,   mag_z;
} imu_values;

imu_values imu_read();

/////////////////////
/////   UV     /////
void uv_init();

typedef struct {
    float uv;
} uv_values;

uv_values uv_read();

/////////////////////
///// HUMIDITY /////
void humidity_init();

typedef struct {
    float humidity;
    float temperature;
} humidity_values;

humidity_values humidity_read();

/////////////////////
///// PRESSURE /////
void pressure_init();

typedef struct {
    float altitude;
    float pressure;
    float temperature;
} pressure_values;

pressure_values pressure_read();

#endif