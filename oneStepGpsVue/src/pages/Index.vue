<template>
    <section>
        <div class="my-2 offset-1">
            <GoBack v-if="this.$route.params.id != undefined" />
        </div>
        <div class="mx-4 alert alert-danger " role="alert" v-if="!hidden"> {{ error }}
        </div>
        <div id="map"></div>
    </section>
</template>

<script>
import config from '../../config/dev.env.js'
import GoBack from '@/components/GoBack'
import axios from 'axios'

export default {
    components: {
        GoBack
    },

    props: {
        id: {
            type: String
        }
    },
    data() {
        return {
            map: '',
            error: '',
            hidden: true,
            devices: '',
            focus: { lat: 37.6000, lng: -95.6650 },
            zoom: 5
        }
    },

    async mounted() {
        await this.getData();
        this.initMap();
        this.setDefaultCoords();
        await this.showMarkers();
    },

    methods: {
        initMap() {
            let id = this.$route.params.id;
            console.log(id);
            if (id != undefined) {
                let device = this.devices.find(x => x.device_id === id);
                if (device != null) {
                    this.focus.lat = device.lat;
                    this.focus.lng = device.lng;
                    this.zoom = 18;
                }
            }
            this.map = new google.maps.Map(document.getElementById('map'), {
                zoom: this.zoom,
                center: new google.maps.LatLng(this.focus.lat, this.focus.lng)
            });
        },

        async getData() {
            let key = { apiKey: localStorage.getItem('key') };
            let resp = await axios.post(config.API_HOST + "devices", key);
            if (resp.data.length === 0) {
                this.showError("Could not get data");
                return;
            }
            this.devices = resp.data;
        },

        setDefaultCoords() {
            if (navigator.geolocation) {
                navigator.geolocation.getCurrentPosition(
                    (position) => {
                        new google.maps.Marker({
                            position: new google.maps.LatLng(position.coords.latitude, position.coords.longitude),
                            map: this.map
                        });
                    });
            }
        },

        async showMarkers() {
            for (let i = 0; i < this.devices.length; i++) {
                await this.createMarker(await this.getMarkerInfo(this.devices[i]));
            }
        },

        async createMarker(marker) {
            let color = marker.status == "active" ? "green" : "red";
            let m = new google.maps.Marker({
                position: new google.maps.LatLng(marker.lat, marker.lng),
                map: this.map,
                icon: this.getMarkerOptions(color),
                pixelOffset: new google.maps.Size(0, 10)
            });

            m.addListener('click', (googleMapsEvent) => {
                let infoWindow = this.createInfoWindow(m, marker);
                infoWindow.iwObject.open(infoWindow.iwOpenOptions);
            })
        },

        createInfoWindow(marker, data) {
            let infoWindow = new google.maps.InfoWindow({
                content: this.getContentString(data),
                anchor: marker,
                pixelOffset: new google.maps.Size(0, 10)
            });

            let infoWindowOpenOptions = {
                map: this.map
            }

            return { iwObject: infoWindow, iwOpenOptions: infoWindowOpenOptions };
        },

        async getMarkerInfo(deviceInfo) {
            let object = {
                deviceId: deviceInfo.device_id,
                name: deviceInfo.display_name,
                status: deviceInfo.active_state,
                odometer: deviceInfo.odometer_mi,
                speed: deviceInfo.speed_mph,
                coords: `${deviceInfo.lat}&#176;, ${deviceInfo.lng}&#176;`,
                lat: deviceInfo.lat,
                lng: deviceInfo.lng,
                lastUpdate: deviceInfo.dt_tracker,
                address: ""
            };

            object.address = await this.getAddressFrom(object.lat, object.lng);
            return object;
        },

        getMarkerOptions(color) {
            return {
                path: config.MARKER_SVG_PATH,
                fillColor: color,
                fillOpacity: 1,
                strokeWeight: 0,
                scale: 2
            };
        },

        getContentString(data) {
            return `<div>
                        <div>
                            <p><b>Object: </b>${data.name}</p>
                            <p><b>Address: </b>${data.address}</p>
                            <p><b>Position: </b>${data.coords}</p>
                            <p><b>Speed: </b>${data.speed} mph</p>
                            <p><b>Status: </b>${data.status}</p>
                            <p><b>Odometer: </b>${data.odometer} mi</p>
                            <p><b>LastUpdate: </b>${data.lastUpdate}</p>
                        </div>
                    </div>`;

        },

        async getAddressFrom(lat, long) {
            let url = `https://maps.googleapis.com/maps/api/geocode/json?latlng=${lat},${long}&key=${config.GOOGLE_API_KEY}`
            let response = await axios.get(url);
            if (response.data.error_message) {
                this.showEror(response.data.error_message);
            } else {
                return response.data.results[0].formatted_address;
            }
        },

        showError(data) {
            this.error = data;
            this.hidden = false;
        }
    }
}
</script>
<style>
#map {
    height: calc(100vh - 10px);
    overflow: hidden;
}
</style>