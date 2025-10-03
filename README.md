# Bin Day Go Web Server

This project provides a simple Go web server exposing an HTTP endpoint `/bin-day`.

## Features
- Returns the day of the week for bin collection (hardcoded as Thursday)
- Returns the type of waste collection for the current week (alternates between Recycling and General Waste)

## Usage

1. Build and run the server:
   ```powershell
   go run main.go
   ```
2. Access the endpoint:
   
   Open your browser or use curl:
     ```powershell
     curl http://localhost:8080/bin-day
     ```

## Output Example
```
{
  "day_of_week": "Thursday",
  "waste_type": "Recycling"
}
```
