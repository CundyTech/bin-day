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
	info := getBinInfo()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func main() {
	http.HandleFunc("/bin-day", binHandler)
	_ = http.ListenAndServe(":8080", nil)
}
