package inits

import (
	"airhass/config"
	"airhass/global"
	"airhass/sensors"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	global.Logger.Debugf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	global.Logger.Debugf("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	global.Logger.Debugf("Connect lost: %v", err)
}

func MQTT() error {
	// Initialize with config options
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", config.Config.MQTT.Broker, config.Config.MQTT.Port))
	opts.SetClientID(config.Config.HASS.DeviceName)
	opts.SetUsername(config.Config.MQTT.Username)
	opts.SetPassword(config.Config.MQTT.Password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	// Create client
	global.MQTT = mqtt.NewClient(opts)

	// Wait for connection
	token := global.MQTT.Connect()
	if token.Wait(); token.Error() != nil {
		return token.Error()
	}

	// Publish configurations
	for _, sensor := range config.Config.Sensors {
		sensors.PublishState(&sensor)
	}

	return nil
}
