#ifndef GPIO_H
#define GPIO_H

#include <stdbool.h>

void gpio_init();

bool setPinMode(int pin, int mode);
bool setPinFunction(int pin, int pinFunction);
int getPinValue(int pin);
int getPinValues();
bool setDigital(int pin, int value);
bool setPWM(int pin, float percentage, float frequency);
bool setServoAngle(int pin, float angle, float min_pulse_ms);

#endif