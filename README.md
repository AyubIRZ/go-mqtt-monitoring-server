# Go MQTT monitoring server

This is a simple MQTT based golang back-end service to monitor, take action and log temperature sensor data. It gets the continuously streamed/published data from an MQTT topic and then takes action based on desired logic implemented on the back-end. After that it will log the acknowledge to an online log channel based on a Telegram Bot, with the help of the official Bot API. 

# Associated resources
- Introduction and review video: [youtu.be/zXzmXzBmWdY](https://youtu.be/zXzmXzBmWdY)
- ESP8266 code repo: [AyubIRZ/esp8266-mqtt-temperature-monitoring](https://github.com/AyubIRZ/esp8266-mqtt-temperature-monitoring)

## Notes
- This is a simple demonstration of the project. You may not use it in production without reviewing the code and changing it to the proper version of your use!
- Data transmission between clients is not secured with TLS or authentication! It's simple and dumb!
-  Any contribution, PR or submitting issues using github issue tracker will be appreciated!
