package main

import (
	"encoding/json"
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




func main() {
	//make goroutine that runs in background
	//generates payload every second

	go func(){
		for{
			writeTelemetryToFile();
			time.Sleep(1*time.Second);
		}
	}()


	http.HandleFunc("/data", handler)

	http.Handle("/", http.FileServer(http.Dir(".")))
	
	http.ListenAndServe(":8080", nil)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}

}
