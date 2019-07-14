#include <stdexcept>
#include "matrix_hal/matrixio_bus.h"
extern "C" {
    #include "matrix.h"
}
#include "bus.hpp"


// Initialize bus for MATRIX hardware communication
matrix_hal::MatrixIOBus bus;

bool busInit(){
    if (bus.Init())
        return true;
    
    return false;
}
