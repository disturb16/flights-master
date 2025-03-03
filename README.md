# Overview

This app is a test

## Setup

Create a configuration file as `conf.json` inside root folder
with the following structure:

```json
{
  "db_path": "data.db",
  "app_addr": ":9090",
  "log_level": -4,
  "serpapi": {
    "api_key": "your-serapi-apikey"
  }
}
```

## How to use

You should run program `go run main.go` and then call the endpoint `/travel` endpoint.

```shell
curl --location 'localhost:9090/travel?origin_airport_code=CDG&destination_airport_code=AUS&date=2025-03-04&destination=Austin%20Texas'
```
