# gpio-to-api
 Configure to monitor buttons, temperature sensors and control various other GPIOs using API endpoints

## Currently supported
- SHT3x Humidity sensor
- Turning on/off gpio pin (fan on/off etc)
- Buttons

## Configuration
All configuration is set by environment variables.  
`GPIO_TYPE` - sht3x, gpio, button  
`GPIO_PIN` - [1-40](https://pinout.xyz/pinout/pin1_3v3_power)

### Examples
**sht3x**
```
GPIO_TYPE=sht3x
UNIT=C
```
**gpio**  
Performs simple gpio pin on/off
```
GPIO_TYPE=gpio
GPIO_PIN=40
```
**button**
```
GPIO_TYPE=button
GPIO_PIN=11
ON_PRESS=callUrl
URL=http://my-service.com/button-pressed
```

## Drivers

Drivers for various sensors, buttons etc can be found here:-  
https://github.com/hybridgroup/gobot/tree/release/drivers

## TODO
- Add go tests
- Add metrics
