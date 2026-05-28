Real time telemetry visualizer built with Go and a browser frontend. It generates fake sensor data and streams it to the client using SSE which is then plotted with chart.js.

<img width="2879" height="1799" alt="image" src="https://github.com/user-attachments/assets/f5be6f4c-777d-4cd2-a85f-5a36db5e8e0f" />


## How to run
- go run .
- http://localhost:8080

## Features
- Real-time telemetry stream over Server-Sent Events (SSE)
- Configurable simulation scenarios:
  - Normal telemetry
  - Warmup
  - Power spike events
  - Trigger dropout simulation
- Adjustable noise levels (low / medium / high)
- Live charts for temperature, voltage, and power
- Control panel to modify system behavior at runtime
- Go backend with no external libraries
- JS frontend using Chart.js
- *No internet connection required*


## Data flow:

1. Go goroutine generates telemetry at a rate (rate can be changed in config)
2. Data is written to `payload.json`
3. SSE endpoint (`/events`) streams updates to the browser
4. Frontend updates DOM + charts in real time
5. Control panel sends config updates to `/config`

"Warmup Mode:"
<img width="1899" height="915" alt="image" src="https://github.com/user-attachments/assets/2b46337f-b8f0-4015-8d47-a48ea06347a9" />

"Power Spike Mode":
<img width="2879" height="1799" alt="image" src="https://github.com/user-attachments/assets/2117960e-e168-47d2-a9d3-48b46623e1bd" />

"Trigger Dropout": (similar to normal mode, but the triggerRate randomly drops to 0)
<img width="1899" height="909" alt="image" src="https://github.com/user-attachments/assets/957e8200-ed0c-446e-aead-b9f58315e818" />
