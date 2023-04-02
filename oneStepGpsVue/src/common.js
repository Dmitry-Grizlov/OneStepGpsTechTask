import config from '../config/dev.env.js'
import axios from 'axios'
export default {
    async getDvices() {
        let key = { apiKey: localStorage.getItem('key') };
        let result;
        await axios.post(config.API_HOST + "devices", key)
            .then((response) => {
                result = response.data.Data.filter(x => x.lat != 0 && x.long != 0)
            })

        return result;
    },

    async getAddressFrom(lat, long) {
        let url = `https://maps.googleapis.com/maps/api/geocode/json?latlng=${lat},${long}&key=${config.GOOGLE_API_KEY}`
        let address;
        await axios.get(url)
            .then((response) => {
                address = response.data.results[0].formatted_address;
            })
            .catch((error) => {
                this.showError(error.message);
            })

        return address;
    },

    isAuthorised() {
        return localStorage.getItem("key") != null && localStorage.getItem("userId") != null;
    }
}