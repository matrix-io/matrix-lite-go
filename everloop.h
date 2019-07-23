#ifndef EVERLOOP_H
#define EVERLOOP_H

void everloop_init();

typedef struct{
    int r, g, b, w;
} led;

// number of LEDs on device
int everloopLength;

// render LEDS colors on MATRIX device
void everloop_set(led leds[]);

#endif