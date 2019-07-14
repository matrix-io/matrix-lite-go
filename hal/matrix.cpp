extern "C" {
    #include "matrix.h"
}

#include "matrix_hal/matrixio_bus.h"

// #include "drivers/everloop.h"
// #include "drivers/gpio.h"
// #include "drivers/info.h"
// #include "drivers/sensors/sensors.h"

// Initialize bus for MATRIX hardware communication
matrix_hal::MatrixIOBus bus;


    bool busInit(){
        return bus.Init();  
        // return true;
    }
