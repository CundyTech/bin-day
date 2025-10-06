package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type BinInfo struct {
	DayOfWeek string `json:"day_of_week"`
	WasteType string `json:"waste_type"`
}

func getBinInfo() BinInfo {
	// Example hardcoded schedule: alternate weeks for "Recycling" and "General Waste"
	_, week := time.Now().ISOWeek()
	var wasteType string
	if week%2 == 0 {
		wasteType = "Recycling"
	} else {
		wasteType = "General Waste"
	}
	return BinInfo{
		DayOfWeek: "Thursday", // Change as needed
		WasteType: wasteType,
	}
}

func binHandler(w http.ResponseWriter, r *http.Request) {
	println("[INFO] received get bin info request")
	info := getBinInfo()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	println("[INFO] received health check request")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

func main() {
	http.HandleFunc("/bin-day", binHandler)
	http.HandleFunc("/health", healthHandler)
	println("[INFO] bin-day service started on :8080")
	_ = http.ListenAndServe(":8080", nil)
}
