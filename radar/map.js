const config = {
  minZoom: 4,
  maxZoom: 12,
};

const zoom = 4;

// 39.50N, 98.35W
const lat = 39.50;
const long = -98.35;

const map = L.map("map", config).setView([lat, long], zoom);

L.tileLayer("https://tile.openstreetmap.org/{z}/{x}/{y}.png", {
  attribution:
    '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
}).addTo(map);
