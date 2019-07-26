#ifndef GPIO_H
#define GPIO_H

#include <stdbool.h>

void gpio_init();

bool setMode(int pin, int mode);
bool setFunction(int pin, int pinFunction);
int getValue(int pin);
int getValues();
bool setDigital(int pin, int value);
bool setPWM(int pin, float percentage, float frequency);
bool setServoAngle(int pin, float angle, float min_pulse_ms);

#endif