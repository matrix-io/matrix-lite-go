#ifndef EVERLOOP_H
#define EVERLOOP_H

void everloop_init();

// Create LED configurations
typedef struct Led{
    int r, g, b, w;
} Led;

// Number of LEDs on device
int everloopLength;

// void newLed(led);
void everloop_set(Led led[35]);

#endif