package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

//Notes: handleFunc takes a handler function that is run when connecting to server
//JSON fields need to be capital letter
//float -> float64

var filePath string = "payload.json"

type NoiseLevel string
const(
	low NoiseLevel = "low" 
	medium NoiseLevel = "medium" 
	high NoiseLevel = "high"
)

type Scenario string
const(
	normal Scenario = "normal"
	warmup Scenario = "warmup"
	powerSpike Scenario = "powerSpike"
	triggerDropout Scenario = "triggerDropout"

)

type Config struct {
	TelemetryRate int `json:"telemetryRate"`
	NoiseLevel NoiseLevel `json:"noiseLevel"`
	Scenario Scenario `json:"scenario"`
}

var currentConfig = Config{
	TelemetryRate: 5,
	NoiseLevel: low,
	Scenario: normal,
}

/*
//browser askes, server responds, stop
func handler(w http.ResponseWriter, r *http.Request){
	//content = byte[], error = error
	content, error := os.ReadFile(filePath)

	if error != nil{
		http.Error(w, "File could not be read",1);
		return;
	}

	var payload Payload
	if error := json.Unmarshal(content,&payload); error != nil{
		http.Error(w,"JSON could not be parsed",1);
		return;
	}

	//send JSON to browser

	w.Header().Set("Content-type","application/json");
	json.NewEncoder(w).Encode(payload);
	
}
*/

//browser connects, connection stays open
//w lets you send info back to request
//r contains info about request
//reads current telemtry data, sends it to browser
func eventsHandler(w http.ResponseWriter, r *http.Request){

	//needed SSE headers to keep connection alive
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")


	//flusher is made so updates can be sent immediately
	flusher := w.(http.Flusher)
	for{
		//content = byte[], error = error
		content, error := os.ReadFile(filePath)

		if error != nil{
			http.Error(w, "File could not be read",500);
			return;
		}

		//converts JSON into payload struct
		var payload Payload
		if error := json.Unmarshal(content,&payload); error != nil{
			http.Error(w,"JSON could not be parsed",500);
			return;
		}

		//converts struct back into JSON
		jsonData, error := json.Marshal(payload);

		if error!= nil{
			http.Error(w, "Struct couldn't be turned into JSON",500);
		}

		//sends to browser
		fmt.Fprintf(w, "data: %s\n\n", jsonData)

		//forces data out asap
		flusher.Flush();

		//telemetry rate needs to be duration to divide by duration
		time.Sleep(time.Second/time.Duration(currentConfig.TelemetryRate));
	}


}

//runs when request arrives at config
func configHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var updates map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		http.Error(w, "bad JSON", http.StatusBadRequest)
		return
	}

	if val, ok := updates["telemetryRate"]; ok {
		currentConfig.TelemetryRate = int(val.(float64))
	}

	if val, ok := updates["noiseLevel"]; ok {
		currentConfig.NoiseLevel = NoiseLevel(val.(string))
	}

	if val, ok := updates["scenario"]; ok {
		currentConfig.Scenario = Scenario(val.(string))
	}

	fmt.Println("updated config:", currentConfig)

	w.WriteHeader(http.StatusOK)
}

func main() {
	//make goroutine that runs in background
	//generates payload every second
	//lightweight thread
	go func(){
		for{
			writeTelemetryToFile();
			time.Sleep(time.Second/time.Duration(currentConfig.TelemetryRate));
		}
	}()

	// "/data" for request/repsonse
	// "/events" for open connection
	http.HandleFunc("/events", eventsHandler)
	http.HandleFunc("/config", configHandler)
	//gets server frontend files
	http.Handle("/", http.FileServer(http.Dir("static")))
	
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}

}
