package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"
)

type Payload struct {
	TemperatureC  float64 `json:"temperatureC"`
	VoltageV      float64 `json:"voltageV"`
	PowerW        float64 `json:"powerW"`
	HfPingRateHz  float64 `json:"hfPingRateHz"`
	LfPingRateHz  float64 `json:"lfPingRateHz"`
	TriggerRateHz float64 `json:"triggerRateHz"`
	Timestamp     string  `json:"timestamp"`
}

//make a bunch of fake telemetry data
func generateTelemetryValues() Payload{
	return Payload{
		TemperatureC: rand.Float64()*20,
		VoltageV: rand.Float64()*20,
		PowerW: rand.Float64()*20,
		HfPingRateHz: rand.Float64()*20,
		LfPingRateHz: rand.Float64()*20,
		TriggerRateHz: rand.Float64()*20,
		Timestamp: time.Now().String(),

	}

}

func writeTelemetryToFile(){
	payload := generateTelemetryValues();
	jsonData, _ := json.MarshalIndent(payload,"", " ");
	os.WriteFile("payload.json",jsonData,0644);
}