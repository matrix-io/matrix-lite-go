#ifndef SENSORS_H
#define SENSORS_H

/////////////////////
/////   IMU    /////
void imu_init();

struct imu_values {
    float accel_x, accel_y, accel_z;
    float gyro_x,  gyro_y,  gyro_z;
    float yaw,     pitch,   roll;
    float mag_x,   mag_y,   mag_z;
};
struct imu_values imu_read();

/////////////////////
/////   UV     /////
void uv_init();

struct uv_values {
    float uv;
};
struct uv_values uv_read();

/////////////////////
///// HUMIDITY /////
void humidity_init();

struct humidity_values {
    float humidity;
    float temperature;
};
struct humidity_values humidity_read();

/////////////////////
///// PRESSURE /////
void pressure_init();

struct pressure_values {
    float altitude;
    float pressure;
    float temperature;
};
struct pressure_values pressure_read();

#endif