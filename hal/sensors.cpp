#include "matrix_hal/uv_sensor.h"
#include "matrix_hal/uv_data.h"

#include "bus.hpp"
extern "C"{
    #include "sensors.h"
}



matrix_hal::UVData uv_data;
matrix_hal::UVSensor uv_sensor;

void uv_init(){
    uv_sensor.Setup(&bus);
}

uv_values uvRead(){
    struct uv_values data;

    // Update output values
    uv_sensor.Read(&uv_data);
    data.uv = uv_data.uv;

    return data;
}