Real time telemetry visualizer built with Go and a browser frontend. It generates fake sensor data and streams it to the client using SSE which is then plotted with chart.js.

## How to run
- go run .
- http://localhost:8080

## Features
- Real-time telemetry stream over Server-Sent Events (SSE)
- Configurable simulation scenarios:
  - Normal sinusoidal telemetry
  - Warmup (low, stable readings)
  - Power spike events
  - Trigger dropout simulation
- Adjustable noise levels (low / medium / high)
- Live charts for temperature, voltage, and power
- Control panel to modify system behavior at runtime
- Pure Go backend with no external libraries
- Vanilla JS frontend using Chart.js


## Data flow:

1. Go goroutine generates telemetry at a rate (rate can be changed in config)
2. Data is written to `payload.json`
3. SSE endpoint (`/events`) streams updates to the browser
4. Frontend updates DOM + charts in real time
5. Control panel sends config updates to `/config`
