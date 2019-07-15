#include "matrix_hal/everloop.h"
#include "matrix_hal/everloop_image.h"
#include "bus.hpp"
#include <stdio.h>
extern "C"{
    #include "everloop.h"
}

matrix_hal::Everloop hal_everloop;

// setup matrix bus
void everloop_init(){
    hal_everloop.Setup(&bus);

    // get number device LEDs
    everloopLength = bus.MatrixLeds();
}

// TODO REMOVE HARDCODED LED COUNT!!!!!!
// set LEDS on MATRIX device
void everloop_set(Led led[35]){

    // Create everloop image
    matrix_hal::EverloopImage everloop_image(35);
    for (int i = 0; i < 35; i++){
        everloop_image.leds[i].red = led[i].r;
        everloop_image.leds[i].green = led[i].g;
        everloop_image.leds[i].blue = led[i].b;
        everloop_image.leds[i].white = led[i].w;
    }

    // Render everloop image
    hal_everloop.Write(&everloop_image);
}