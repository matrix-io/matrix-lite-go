#ifndef SENSORS_H
#define SENSORS_H

/////////////////////
/////   IMU    /////
void imu_init();

struct imu_values {
    // Accelerometer Output
    float accel_x;
    float accel_y;
    float accel_z;
    // Gyroscope Output
    float gyro_x;
    float gyro_y;
    float gyro_z;
    // Yaw, Pitch, Roll Output
    float yaw;
    float pitch;
    float roll;
    // Magnetometer Output
    float mag_x;
    float mag_y;
    float mag_z;
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