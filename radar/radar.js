L.tileLayer.wms(
  "https://opengeo.ncep.noaa.gov/geoserver/conus/conus_bref_qcd/ows?",
  {
    layers: "conus_bref_qcd",
    format: "image/png",
    transparent: true,
  },
).addTo(map);
