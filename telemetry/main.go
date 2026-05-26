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
			http.Error(w, "File could not be read",1);
			return;
		}

		//converts JSON into payload struct
		var payload Payload
		if error := json.Unmarshal(content,&payload); error != nil{
			http.Error(w,"JSON could not be parsed",1);
			return;
		}

		//converts struct back into JSON
		jsonData, error := json.Marshal(payload);

		if error!= nil{
			http.Error(w, "Struct couldn't be turned into JSON",1);
		}

		fmt.Fprintf(w, "data: %s\n\n", jsonData)

		flusher.Flush();

		time.Sleep(1*time.Second);
	}


}



func main() {
	//make goroutine that runs in background
	//generates payload every second

	go func(){
		for{
			writeTelemetryToFile();
			time.Sleep(1*time.Second);
		}
	}()

	// /data for request/repsonse, /events for open connection
	http.HandleFunc("/events", eventsHandler)

	http.Handle("/", http.FileServer(http.Dir("static")))
	
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}

}
