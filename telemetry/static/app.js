function updateDashboard(data) {
  const dashboard = document.getElementById("dashboard");

  dashboard.innerHTML = "";

  for (const key in data) {
    const p = document.createElement("p");
    p.textContent = `${key}: ${data[key]}`;
    dashboard.appendChild(p);
  }
}

const events = new EventSource("/events");

events.onmessage = function (event) {
  const data = JSON.parse(event.data);
  updateDashboard(data);
};