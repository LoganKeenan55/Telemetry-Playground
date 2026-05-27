function updateDashboard(data) {
  //finds html element
  const dashboard = document.getElementById("dashboard");

  //clear
  dashboard.innerHTML = "";

  for (const key in data) {
    const p = document.createElement("p");
    p.textContent = `${key}: ${data[key]}`;
    dashboard.appendChild(p);
  }
}



const MAX_POINTS = 60;

let temperatureData = [];
let voltageData = [];
let powerData = [];

const temperatureChart = new Chart(document.getElementById("temperatureChart"),
{
type: "line",
data:{
    labels:[],
    datasets:[{
        label: "Temperature (c)",
        data: []
    }]
},
options: {
  scales: {
    y: {
      max: 100 ,
      min: 0
    }
  }
}
});
const voltageChart = new Chart(document.getElementById("voltageChart"),
{
type: "line",
data:{
    labels:[],
    datasets:[{
        label: "Voltage (V)",
        data: []
    }]
},
options: {
  scales: {
    y: {
      max: 100 ,
      min: 0
    }
  }
}
});
const powerChart = new Chart(document.getElementById("powerChart"),
{
type: "line",
data:{
    labels:[],
    datasets:[{
        label: "Power (W)",
        data: []
    }]
},
options: {
  scales: {
    y: {
      max: 100 ,
      min: 0
    }
  }
}
});

function addPointToChart(array, value){
    array.push(value);
    if(array.length > MAX_POINTS){
        array.shift();
        
    }
}


function updateCharts(){
    //x axis will just be numbers 1 -> max
    const labels = temperatureData.map((val,i) => i+1);

    temperatureChart.data.labels = labels; //x
    temperatureChart.data.datasets[0].data = temperatureData; //y
    temperatureChart.update("none");

    voltageChart.data.labels = labels;
    voltageChart.data.datasets[0].data = voltageData;
    voltageChart.update("none");

    powerChart.data.labels = labels;
    powerChart.data.datasets[0].data = powerData;
    powerChart.update("none");
}


//SSE connection
const events = new EventSource("/events");

//runs everytime server sennds {data: []} from go
events.onmessage = function (event) {
    const data = JSON.parse(event.data);
    updateDashboard(data);

    addPointToChart(temperatureData,data.temperatureC);
    addPointToChart(voltageData,data.voltageV);
    addPointToChart(powerData,data.powerW);


    updateCharts(data);
};
