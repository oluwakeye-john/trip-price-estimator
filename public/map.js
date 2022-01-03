let _map;
let directionsService;
let directionsRenderer;

function initMap() {
  const defaultCenter = { lat: 7.3775, lng: 3.947 };

  directionsService = new google.maps.DirectionsService();
  directionsRenderer = new google.maps.DirectionsRenderer();

  _map = new google.maps.Map(document.getElementById("map"), {
    zoom: 10,
    center: defaultCenter,
  });

  directionsRenderer.setMap(_map);
}

function plotMap(origin, destination) {
  directionsService
    .route({
      origin,
      destination,
      travelMode: google.maps.TravelMode.DRIVING,
    })
    .then((res) => {
      directionsRenderer.setDirections(res);
      console.log({ res });
    });
}
