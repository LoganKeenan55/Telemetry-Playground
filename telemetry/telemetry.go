package main

import (
	"encoding/json"
	"math"
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
	var currentNoise float64

	switch currentConfig.NoiseLevel{
	case low:
		currentNoise = 5;
	case medium:
		currentNoise = 20;
	case high:
		currentNoise = 50;
	}

	t := float64(time.Now().UnixMilli())/1000.0;

	//default values
	var temperature float64
	var voltage float64
	var power float64
	var hfPing float64
	var lfPing float64
	var triggerRate float64

		switch currentConfig.Scenario {

	//smooth sine wave telemetry
	case normal:

		base := 25 + math.Cos(t)*25

		temperature = 25 + math.Sin(t)*25+ (rand.Float64()*2-1)*currentNoise
		voltage = base + (rand.Float64()*2-1)*currentNoise
		power = base + (rand.Float64()*2-1)*currentNoise
		hfPing = base + (rand.Float64()*2-1)*currentNoise
		lfPing = base + (rand.Float64()*2-1)*currentNoise
		triggerRate = base + (rand.Float64()*2-1)*currentNoise

	//flat constant values
	case warmup:

		temperature = 5 + math.Sin(t)*5
		voltage = 5 + math.Cos(t)*5
		power = 5 + math.Cos(t)*5
		hfPing = 1
		lfPing = 1
		triggerRate = 1

	//normal values with occasional huge spike
	case powerSpike:

		temperature = rand.Float64() * currentNoise
		voltage = rand.Float64() * currentNoise
		power = rand.Float64() * currentNoise
		hfPing = rand.Float64() * currentNoise
		lfPing = rand.Float64() * currentNoise
		triggerRate = rand.Float64() * currentNoise

		//5% chance of giant power spike
		if rand.Float64() < 0.05 {
			power += 500
		}

	//everything drops to zero
	case triggerDropout:

		temperature = 0
		voltage = 0
		power = 0
		hfPing = 0
		lfPing = 0
		triggerRate = 0
	}

	return Payload{
		TemperatureC: temperature,
		VoltageV: voltage,
		PowerW: power,
		HfPingRateHz: hfPing,
		LfPingRateHz: lfPing,
		TriggerRateHz: triggerRate,
		Timestamp: time.Now().String(),
	}
}

func writeTelemetryToFile(){
	payload := generateTelemetryValues();
	jsonData, _ := json.MarshalIndent(payload,"", " ");
	os.WriteFile("payload.json",jsonData,0644);
}